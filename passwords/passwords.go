package passwords

import (
	"encoding/json"
	"fmt"
	"os"
)

type PswStorage struct {
	Description string
	Password    string
}

var PswList []PswStorage

const fileName = "passwords.json"

func LoadPasswords() error {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			PswList = []PswStorage{}
			return nil
		}
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(&PswList)
}
func SavePasswords() error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(PswList)
}
func ListDescriptions() []string {
	descriptions := []string{}
	for _, psw := range PswList {
		descriptions = append(descriptions, psw.Description)
	}
	return descriptions
}
func AddPassword(description, password string) error {
	PswList = append(PswList, PswStorage{Description: description, Password: password})
	return SavePasswords()
}

func GetPasswordByDescription(description string) (string, error) {
	for _, psw := range PswList {
		if psw.Description == description {
			return psw.Password, nil
		}
	}
	return "", fmt.Errorf("no password found for the given description")
}
