package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	*http.Client
	token    string
	Username string
	Password string
}

func NewHttpClient() *Client {
	return &Client{Client: &http.Client{
		Timeout:   30 * time.Second,
		Transport: &http.Transport{DisableKeepAlives: true, TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
	}}
}

func NewHttpsClient(certPem, keyPem []byte) (*Client, error) {
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return nil, err
	}

	return &Client{Client: &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}}, nil
}

func NewHttpsClientSkipVerify(certPem, keyPem []byte) (*Client, error) {
	cert, err := tls.X509KeyPair(certPem, keyPem)
	if err != nil {
		return nil, err
	}

	return &Client{Client: &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
			},
		},
	}}, nil
}

func (cli *Client) SetBaseAuth(username, password string) *Client {
	cli.Username = username
	cli.Password = password
	return cli
}

func (cli Client) HasBaseAuth() bool {
	return cli.Username != "" && cli.Password != ""
}

func (cli *Client) SetToken(token string) *Client {
	cli.token = token
	return cli
}

func (cli *Client) SetTimeout(timeout time.Duration) *Client {
	cli.Timeout = timeout
	return cli
}

func (cli *Client) Post(url string, req, resp interface{}) error {
	return cli.request(http.MethodPost, url, req, resp)
}

func (cli *Client) Get(url string, resp interface{}) error {
	return cli.request(http.MethodGet, url, nil, resp)
}

func (cli *Client) Put(url string, req, resp interface{}) error {
	return cli.request(http.MethodPut, url, req, resp)
}

func (cli *Client) Delete(url string, req, resp interface{}) error {
	return cli.request(http.MethodDelete, url, req, resp)
}

func (cli *Client) request(httpMethod, url string, req, resp interface{}) error {
	var httpReqBody io.Reader
	if req != nil {
		reqBody, err := json.Marshal(req)
		if err != nil {
			return fmt.Errorf("marshal request failed: %s", err.Error())
		}

		httpReqBody = bytes.NewBuffer(reqBody)
	}

	httpReq, err := http.NewRequest(httpMethod, url, httpReqBody)
	if err != nil {
		return fmt.Errorf("new http request failed: %s", err.Error())
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if cli.token != "" {
		httpReq.Header.Set("Authorization", cli.token)
	}
	if cli.HasBaseAuth() {
		httpReq.SetBasicAuth(cli.Username, cli.Password)
	}

	httpResp, err := cli.Do(httpReq)
	if err != nil {
		return fmt.Errorf("send http request failed: %s", err.Error())
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	defer httpResp.Body.Close()
	if err != nil {
		return fmt.Errorf("read http response body failed: %s", err.Error())
	}

	if httpResp.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("http request failed:%s", body)
	}

	if len(body) > 0 && resp != nil {
		if err := json.Unmarshal(body, &resp); err != nil {
			return fmt.Errorf("unmarshal http response failed: %s", err.Error())
		}
	}

	return nil
}
