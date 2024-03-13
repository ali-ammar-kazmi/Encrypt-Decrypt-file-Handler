package handlers

import (
	"os"

	"github.com/ali-ammar-kazmi/Encrypt_Decrypt/service"
	"github.com/joho/godotenv"
)

func hash() []byte {
	godotenv.Load()

	key := os.Getenv("SECRET_KEY")
	return []byte(key)
}

func validate(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func EncryptHandler(filePath string) {

	if !validate(filePath) {
		panic("File not found")
	}
	key := hash()
	service.Encrypt(key, filePath)
}

func DecryptHandler(filePath string) {

	if !validate(filePath) {
		panic("File not found")
	}
	key := hash()
	service.Decrypt(key, filePath)
}
