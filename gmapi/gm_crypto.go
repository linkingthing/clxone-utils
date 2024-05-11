package gmapi

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	httputil "github.com/linkingthing/clxone-utils/http"
)

// ApiSm4EcbEncrypt ECB（Electronic Codebook）模式是一种基本的分组密码工作模式
const ApiSm4EcbEncrypt = "/sm4-ecb/encrypt"
const ApiSm4EcbDecrypt = "/sm4-ecb/decrypt"
const ApiHash = "/hash"

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
	AgentUrl     string
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

func (g *GmEncrypt) genApiUrl(api string) string {
	return strings.TrimRight(g.AgentUrl, "/") + "/" + strings.TrimLeft(api, "/")
}

func (g *GmResponse) IsSuccess() bool {
	return g.Success
}

func (g *GmResponse) GetMessage() string {
	return g.Message
}

func InitGmEncrypt(url string, authKey string) error {
	gmEncryptClient = &GmEncrypt{
		AgentUrl: url,
		AuthKey:  authKey,
	}
	if gmEncryptClient.AgentUrl == "" {
		return errors.New("agent url can not be empty")
	}

	if gmEncryptClient.AuthKey == "" {
		log.Printf("auth key is empty")
		return errors.New("auth key can not be empty")
	}

	gmEncryptClient.gmHttpClient = &httputil.Client{Client: &http.Client{
		Timeout:   30 * time.Second,
		Transport: &http.Transport{DisableKeepAlives: true, TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	},
	}

	gmEncryptClient.gmHttpClient.SetHeader("Authentication", authKey)
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
