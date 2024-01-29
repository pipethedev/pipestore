package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
)

func CheckIfTableExists(tableName string) bool {
	_, err := os.Stat(tableName + ".json")
	return err == nil
}

func CreateTableFile(tableName string) error {
	file, err := os.Create(tableName + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	emptyData := []interface{}{}
	encodedData, err := json.Marshal(emptyData)
	if err != nil {
		return err
	}

	_, err = file.Write(encodedData)
	return err
}

func ReadTableData(tableName string) ([]interface{}, error) {
	file, err := os.Open(tableName + ".json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []interface{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteTableData(tableName string, data []interface{}) error {
	file, err := os.Create(tableName + ".json")
	if err != nil {
		return err
	}
	defer file.Close()

	encodedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	_, err = file.Write(encodedData)

	key := []byte("<SECRET_API_KEY>")

	encryptedData, err := EncryptFile(tableName+".json", key)

	if err != nil {
		log.Fatalln("Unable to encrypt file:", err)
	}
	os.WriteFile("encrypted_"+tableName+".json", encryptedData, 0644)

	return err
}

func EncryptFile(filename string, key []byte) ([]byte, error) {

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	encoded := base64.StdEncoding.EncodeToString(ciphertext)

	return []byte(encoded), nil
}

func DecryptFile(encryptedData []byte, key []byte) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(string(encryptedData))

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := decoded[:nonceSize], decoded[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
