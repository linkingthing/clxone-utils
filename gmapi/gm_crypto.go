package gmapi

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/tjfoc/gmsm/gmtls"
	"github.com/tjfoc/gmsm/x509"

	httputil "github.com/linkingthing/clxone-utils/http"
)

// ApiSm4EcbEncrypt ECB（Electronic Codebook）模式是一种基本的分组密码工作模式
const (
	ApiSm4EcbEncrypt = "/datahub/hsm-service/crypto/sm4-ecb/encrypt"
	ApiSm4EcbDecrypt = "/datahub/hsm-service/crypto/sm4-ecb/decrypt"
	ApiHash          = "/datahub/hsm-service/crypto/hash"

	GmModeDev  = "dev"
	GmModeProd = "prod"
)

type GmResponse struct {
	Data    map[string]string `json:"data"`
	Result  string            `json:"result"`
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Success bool              `json:"success"`
}

type GmRequest struct {
	DataList []string `json:"dataList"`
}

type GmEncrypt struct {
	AgentAddr    string
	AuthKey      string
	err          error
	gmHttpClient *httputil.Client
}

var gmEncryptClient *GmEncrypt

func GetGmClient() *GmEncrypt {
	if gmEncryptClient == nil {
		gmEncryptClient = &GmEncrypt{
			err: errors.New("gm client is not initialized"),
		}
	}

	return gmEncryptClient
}

func (g *GmEncrypt) GetHttpClient() *http.Client {
	return gmEncryptClient.gmHttpClient.Client
}

func (g *GmEncrypt) genApiUrl(api string) string {
	return strings.TrimRight(g.AgentAddr, "/") + "/" + strings.TrimLeft(api, "/")
}

func (g *GmResponse) IsSuccess() bool {
	return g.Success
}

func (g *GmResponse) GetMessage() string {
	return g.Message
}

type GMEncryptConf struct {
	//MonitorPlatformAddr 监测平台地址
	MonitorPlatformAddr string `yaml:"monitor_platform_addr"`
	//MonitorPlatformToken 监测平台请求令牌
	MonitorPlatformToken string `yaml:"monitor_platform_token"`
	//GmCaCertPath 国密CA证书路径
	GmCaCertPath string `yaml:"gm_ca_cert_path"`
	//GmAuthCertPath 国密客户端认证证书路径
	GmAuthCertPath string `yaml:"gm_auth_cert_path"`
	//GmAuthKeyPath 国密客户端认证密钥对路径
	GmAuthKeyPath string `yaml:"gm_auth_key_path"`
	//GmMode 国密模式 dev-测试模式,不使用国密客户端 prod或空-正式模式,使用国密客户端
	GmMode string `yaml:"gm_mode"`
}

func InitGmEncrypt(conf GMEncryptConf) error {
	gmEncryptClient = &GmEncrypt{
		AgentAddr: conf.MonitorPlatformAddr,
		AuthKey:   conf.MonitorPlatformToken,
	}
	if gmEncryptClient.AgentAddr == "" {
		return errors.New("agent url can not be empty")
	}

	if gmEncryptClient.AuthKey == "" {
		log.Printf("auth key is empty")
		return errors.New("auth key can not be empty")
	}

	if conf.GmMode == GmModeDev {
		gmEncryptClient.gmHttpClient = &httputil.Client{Client: &http.Client{
			Timeout:   30 * time.Second,
			Transport: &http.Transport{DisableKeepAlives: true, TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		}}
	} else {
		certPool := x509.NewCertPool()
		caCert, err := ioutil.ReadFile(conf.GmCaCertPath)
		if err != nil {
			return err
		}
		certPool.AppendCertsFromPEM(caCert)

		var httpClient *http.Client
		if conf.GmAuthCertPath != "" && conf.GmAuthKeyPath != "" {
			clientAuthCert, err := gmtls.LoadX509KeyPair(conf.GmAuthCertPath, conf.GmAuthKeyPath)
			if err != nil {
				return err
			}
			httpClient = gmtls.NewAuthHTTPSClient(certPool, &clientAuthCert)
		} else {
			httpClient = gmtls.NewHTTPSClient(certPool)
		}
		httpClient.Timeout = 30 * time.Second
		gmEncryptClient.gmHttpClient = &httputil.Client{Client: httpClient}
	}

	gmEncryptClient.gmHttpClient.SetHeader("Authentication", conf.MonitorPlatformToken)
	return nil
}

func (g *GmEncrypt) ApiSm4EcbEncrypt(s ...string) (map[string]string, error) {
	resMap := make(map[string]string)
	if g.err != nil {
		return resMap, g.err
	}

	var resp GmResponse
	req := GmRequest{
		DataList: s,
	}

	api := g.genApiUrl(ApiSm4EcbEncrypt)
	if err := g.gmHttpClient.Post(api, &req, &resp); err != nil {
		return resMap, err
	}

	if resp.IsSuccess() {
		return resp.Data, nil
	}

	return resMap, fmt.Errorf("error:%s", resp.GetMessage())
}

func (g *GmEncrypt) ApiSm4EcbDecrypt(s ...string) (map[string]string, error) {
	resMap := make(map[string]string)
	if g.err != nil {
		return resMap, g.err
	}
	var resp GmResponse
	req := GmRequest{
		DataList: s,
	}
	api := g.genApiUrl(ApiSm4EcbDecrypt)
	if err := g.gmHttpClient.Post(api, &req, &resp); err != nil {
		return resMap, err
	}

	if resp.IsSuccess() {
		return resp.Data, nil
	}

	return resMap, fmt.Errorf("error:%s", resp.GetMessage())
}

func (g *GmEncrypt) ApiHash(s ...string) (map[string]string, error) {
	resMap := make(map[string]string)
	if g.err != nil {
		return resMap, g.err
	}
	var resp GmResponse
	req := GmRequest{
		DataList: s,
	}
	api := g.genApiUrl(ApiHash)
	if err := g.gmHttpClient.Post(api, &req, &resp); err != nil {
		return resMap, err
	}

	if resp.IsSuccess() {
		return resp.Data, nil
	}

	return resMap, fmt.Errorf("error:%s", resp.GetMessage())
}
