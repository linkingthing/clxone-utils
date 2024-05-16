package gmapi

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func setup() {
}

func Test_ApiEncryptAndDecrypt(t *testing.T) {
	realData := "123456"
	err := InitGmEncrypt("http://192.168.31.211:39095/datahub/hsm-service/crypto", "b", "../testdata/SM2_CA.cer")
	if err != nil {
		t.Error(err)
		return
	}
	encryptData, err := GetGmClient().ApiSm4EcbEncrypt(realData)
	if err != nil {
		t.Error(err)
		return
	}

	encryptItem := encryptData[realData]
	if encryptItem == "" {
		t.Error("encrypt data is empty")
		return
	}

	decryptData, err := GetGmClient().ApiSm4EcbDecrypt(encryptItem)
	if err != nil {
		t.Error(err)
		return
	}

	if realData != decryptData[encryptItem] {
		t.Errorf("encrypt and decrypt are not equal")
		return
	}

	t.Logf("real data is %v,encrypt data is %v,decrypt data is %v", realData, encryptData, decryptData)
}

func Test_gmHttpsClient(t *testing.T) {
	err := InitGmEncrypt("https://localhost:443", "b", "../testdata/SM2_CA.cer")
	if err != nil {
		t.Error(err)
		return
	}

	response, err := GetGmClient().GetHttpClient().Get("https://localhost:443/test")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// 使用 response 做你需要的事情...
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("result is: %v", string(result))
}
