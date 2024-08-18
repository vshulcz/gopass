package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"os"
)

func InitializeMasterKey(keyFilePath string) ([]byte, error) {
	if _, err := os.Stat(keyFilePath); os.IsNotExist(err) {
		log.Println("Master key not found. Generating a new one...")
		key := make([]byte, 32)
		if _, err := rand.Read(key); err != nil {
			return nil, err
		}

		encodedKey := hex.EncodeToString(key)
		if err := os.WriteFile(keyFilePath, []byte(encodedKey), 0600); err != nil {
			return nil, err
		}

		log.Println("Master key generated and saved successfully.")
		return key, nil
	}

	encodedKey, err := os.ReadFile(keyFilePath)
	if err != nil {
		return nil, err
	}

	key, err := hex.DecodeString(string(encodedKey))
	if err != nil {
		return nil, err
	}

	if len(key) != 32 {
		return nil, errors.New("invalid master key length")
	}

	return key, nil
}

func Encrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]

	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(key []byte, cryptoText string) (string, error) {
	ciphertext, _ := hex.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return string(ciphertext), nil
}
