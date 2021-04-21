package main

import (
	"fmt"
)

func Registrate() (string, string){
	var login string
	var password string

	isCorrectPassword := false

	fmt.Print("Enter your login: ")
	fmt.Scanln(&login)

	for !isCorrectPassword {
		fmt.Print("Enter your password: ")
		fmt.Scanln(&password)

		if len(password) >= 6 && len(password) <= 10 {
			isCorrectPassword = true
			fmt.Println("Good.")
		} else {
			fmt.Println("Wrong password, try again.")
		}
	}

	return login, password
}
