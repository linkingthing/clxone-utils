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
	InitGmEncrypt("http://192.168.31.211:43263/datahub/hsm-service/crypto", "")
	encryptData, err := GetGmClient().ApiSm4EcbEncrypt(realData)
	if err != nil {
		t.Error(err)
		return
	}

	decryptData, err := GetGmClient().ApiSm4EcbDecrypt(encryptData)
	if err != nil {
		t.Error(err)
		return
	}

	if realData != decryptData {
		t.Errorf("encrypt and decrypt are not equal")
		return
	}

	t.Logf("real data is %v,encrypt data is %v,decrypt data is %v", realData, encryptData, decryptData)
}
