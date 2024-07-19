package main

import (
	"Homework12/passwords"
	"fmt"
)

func main() {
	if err := passwords.LoadPasswords(); err != nil {
		fmt.Println("Error loading passwords:", err)
		return
	}

	for {
		fmt.Println("\nPassword Manager")
		fmt.Println("1. List all password descriptions")
		fmt.Println("2. Add a new password")
		fmt.Println("3. Get a password by description")
		fmt.Println("4. Exit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			passwords.PrintDescriptions()
		case 2:
			passwords.NewPsw()
		case 3:
			passwords.GetPassword()
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
