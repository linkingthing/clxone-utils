package gmapi

import (
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
	err := InitGmEncrypt("http://192.168.31.211:39095/datahub/hsm-service/crypto", "b")
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
