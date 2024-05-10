package gmapi

import (
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	httputil "github.com/linkingthing/clxone-utils/http"
)

// ApiSm4EcbEncrypt ECB（Electronic Codebook）模式是一种基本的分组密码工作模式
const ApiSm4EcbEncrypt = "/sm4-ecb/encrypt"
const ApiSm4EcbDecrypt = "/sm4-ecb/decrypt"
const ApiHash = "/hash"

type GmResponse struct {
	Data    interface{} `json:"data"`
	Result  string      `json:"result"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

type GmRequest struct {
	Data interface{} `json:"data"`
}

type GmEncrypt struct {
	AgentUrl     string
	AuthKey      string
	err          error
	gmHttpClient *httputil.Client
}

var gmEncryptClient *GmEncrypt
var gmOnce = sync.Once{}

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

func InitGmEncrypt(url string, authKey string) {
	gmOnce.Do(func() {
		gmEncryptClient = &GmEncrypt{
			AgentUrl: url,
			AuthKey:  authKey,
		}
		if gmEncryptClient.AgentUrl == "" {
			gmEncryptClient.err = errors.New("agent url can not be empty")
			return
		}

		if gmEncryptClient.AuthKey == "" {
			log.Printf("auth key is empty")
			gmEncryptClient.err = errors.New("auth key can not be empty")
			return
		}

		gmEncryptClient.gmHttpClient = &httputil.Client{Client: &http.Client{
			Timeout:   30 * time.Second,
			Transport: &http.Transport{DisableKeepAlives: true, TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
		}

		gmEncryptClient.gmHttpClient.SetHeader("Authentication", authKey)
	})
}

func (g *GmEncrypt) ApiSm4EcbEncrypt(s interface{}) (string, error) {
	if g.err != nil {
		return "", g.err
	}

	var resp GmResponse
	req := GmRequest{
		Data: s,
	}

	api := g.genApiUrl(ApiSm4EcbEncrypt)
	if err := g.gmHttpClient.Post(api, &req, &resp); err != nil {
		return "", err
	}

	if resp.IsSuccess() {
		if res, ok := resp.Data.(string); ok {
			return res, nil
		} else {
			return "", errors.New("data is not a string")
		}
	}

	return "", fmt.Errorf("error:%s", resp.GetMessage())
}

func (g *GmEncrypt) ApiSm4EcbDecrypt(s string) (string, error) {
	if g.err != nil {
		return "", g.err
	}
	var resp GmResponse
	req := GmRequest{
		Data: s,
	}
	api := g.genApiUrl(ApiSm4EcbDecrypt)
	if err := g.gmHttpClient.Post(api, &req, &resp); err != nil {
		return "", err
	}

	if resp.IsSuccess() {
		if res, ok := resp.Data.(string); ok {
			return res, nil
		} else {
			return "", errors.New("data is not a string")
		}
	}

	return "", fmt.Errorf("error:%s", resp.GetMessage())
}

func (g *GmEncrypt) ApiHash(s string) (string, error) {
	if g.err != nil {
		return "", g.err
	}
	var resp GmResponse
	req := GmRequest{
		Data: s,
	}
	api := g.genApiUrl(ApiHash)
	if err := g.gmHttpClient.Post(api, &req, &resp); err != nil {
		return "", err
	}

	if resp.IsSuccess() {
		if res, ok := resp.Data.(string); ok {
			return res, nil
		} else {
			return "", errors.New("data is not a string")
		}
	}

	return "", fmt.Errorf("error:%s", resp.GetMessage())
}
