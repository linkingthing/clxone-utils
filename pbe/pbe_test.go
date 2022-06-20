package pbe

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
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

func TestDecrypt(t *testing.T) {
	password, err := Decrypt(&DecryptContext{
		KeyFactoryBase64: "ADnD/LjHGmdeUUdmoLHKcgXH9+roVeDozTywwkiQTu0=",
		EncryptWorkKey:   "cmaoqjXY9x1z7aWUHcTcS7bntPruOWYyzqS5XdciN9GihUrus1cJ_57UvYVYwDf2QZBzkMO79G4Y7LOgTw3MfQ==",
		EncryptPassword:  "A5tKnGwd7GeoAfR0EwvjpyChfRDizQ8_h2jLqEwac9mYhpu3EDC6hVT6oKCIVICdGOv9rkUN3Y4NIKM3NZIWnw==",
		Iterator:         10000,
	})
	if err != nil {
		t.Error(fmt.Errorf("parse es password failed:%s", err.Error()))
	}

	t.Log(password)
}

func TestPBKdf2(t *testing.T) {
	seed := "4F/Y86sOxsiTv4/hUKF4DXOaF+2hsCF5l7MObsYkkB0="
	//keyFactoryBase64, err := RandomBase64(32)
	keyFactory, err := base64.StdEncoding.DecodeString(seed)
	if err != nil {
		t.Error(err)
		return
	}

	rootVector := []byte("ipam@123456")
	rootKey := string(pbkdf2.Key(rootVector, keyFactory, 10000, sha512.BlockSize, sha512.New))
	t.Log(rootKey == string(pbkdf2.Key(rootVector, keyFactory, 10000, sha512.BlockSize, sha512.New)))
	t.Log(base64.StdEncoding.EncodeToString([]byte(rootKey)))
}

func TestRandomBase64(t *testing.T) {
	t.Log(RandomBase64(16))
	m := md5.New()
	t.Logf("%s", base64.StdEncoding.EncodeToString(m.Sum([]byte{1})))
}
