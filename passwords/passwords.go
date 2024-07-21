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
func PrintDescriptions() {
	fmt.Println("Saved passwords: ")
	for _, psw := range PswList {
		fmt.Printf("Description: %v\n", psw.Description)
	}
}
func NewPsw() {
	var newDescription, newPassword string
	fmt.Print("Insert description of new password: ")
	fmt.Scan(&newDescription)
	fmt.Print("Insert new password: ")
	fmt.Scan(&newPassword)
	PswList = append(PswList, PswStorage{Description: newDescription, Password: newPassword})
	if err := SavePasswords(); err != nil {
		fmt.Println("Error saving passwords:", err)
	} else {
		fmt.Println("New password added successfully")
	}
}
func GetPassword() {
	var description string
	fmt.Print("Insert description to retrieve the password: ")
	fmt.Scan(&description)
	for _, psw := range PswList {
		if psw.Description == description {
			fmt.Printf("Password for %v: %v\n", description, psw.Password)
			return
		}
	}
	fmt.Println("No password found for the given description")
}
