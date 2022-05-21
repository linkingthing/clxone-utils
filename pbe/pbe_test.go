package pbe

import (
	"testing"

	ut "github.com/linkingthing/cement/unittest"
)

func TestPbe(t *testing.T) {
	keyFactoryBase64, err := RandomBase64(32)
	ut.Assert(t, err == nil, "")
	workKey := "Linking123%^&*90Linking123%^&*90"

	password := "Linking@201907^%$#"
	decryptCtx, err := Encrypt(&EncryptContext{
		KeyFactoryBase64: keyFactoryBase64,
		WorkKey:          workKey,
		Password:         password,
		Iterator:         10000,
	})

	ut.Assert(t, err == nil, "")
	t.Logf("decryptCtx:%+v", decryptCtx)

	decryptPassword, err := Decrypt(decryptCtx)
	ut.Assert(t, err == nil, "")
	ut.Assert(t, decryptPassword == password, decryptPassword)
}
