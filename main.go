package main

import (
	"Homework12/passwords"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Program started")
	if err := passwords.LoadPasswords(); err != nil {
		fmt.Println("Error loading passwords:", err)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("expected 'list', 'put' or 'get' commands")
	}

	switch os.Args[1] {
	case "list":
		descriptions := passwords.ListDescriptions()
		fmt.Println("Saved passwords:")
		for _, desc := range descriptions {
			fmt.Println(desc)
		}
	case "put":
		if len(os.Args) < 4 {
			fmt.Println("Usage: put <description> <password>")
			return
		}
		description := os.Args[2]
		password := os.Args[3]
		err := passwords.AddPassword(description, password)
		if err != nil {
			fmt.Println("Error adding password:", err)
		} else {
			fmt.Println("New password added successfully")
		}
	case "get":
		if len(os.Args) < 3 {
			fmt.Println("Usage: get <description>")
			return
		}
		description := os.Args[2]
		password, err := passwords.GetPasswordByDescription(description)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Password for %v: %v\n", description, password)
		}
	default:
		fmt.Println("Expected 'list', 'put' or 'get'")
	}
}
