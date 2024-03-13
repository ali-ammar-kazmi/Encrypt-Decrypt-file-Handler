package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func Encrypt(key []byte, file string) {

	srcFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()

	plainText, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	salt := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, salt, 4096, 24, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	cipherText := aesgcm.Seal(nil, salt, plainText, nil)
	cipherText = append(cipherText, salt...)

	dstFile, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()

	_, err = dstFile.Write(cipherText)
	if err != nil {
		panic(err.Error())
	}
}

func Decrypt(key []byte, file string) {

	srcFile, err := os.Open(file)
	if err != nil {
		panic(err.Error())
	}
	defer srcFile.Close()

	cipherText, err := io.ReadAll(srcFile)
	if err != nil {
		panic(err.Error())
	}

	salt := cipherText[len(cipherText)-12:]

	nonce := hex.EncodeToString(salt)

	salt, err = hex.DecodeString(nonce)
	if err != nil {
		panic(err.Error())
	}

	dk := pbkdf2.Key(key, salt, 4096, 24, sha1.New)

	block, err := aes.NewCipher(dk)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plainText, err := aesgcm.Open(nil, salt, cipherText[:len(cipherText)-12], nil)
	if err != nil {
		panic(err.Error())
	}

	dstFile, err := os.Create(file)
	if err != nil {
		panic(err.Error())
	}
	defer dstFile.Close()

	_, err = dstFile.Write(plainText)
	if err != nil {
		panic(err.Error())
	}

}
