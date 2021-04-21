package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var choice int

	fmt.Println("\n\tEnter 1 for athorization;\tEnter 2 for registration;")
	fmt.Scanln(&choice)

	switch choice {
		case 1:
			var user User
			jsonFile, err := os.Open("user.txt")

			if err != nil {
				fmt.Println(err)
			}
			defer jsonFile.Close()

			byteValue, _ := ioutil.ReadAll(jsonFile)

			json.Unmarshal(byteValue, &user)

			authorizate(user)
		case 2:
			login, password := Registrate()

			tMin, tMax := Study(password)

			user := &User{
				Login:    login,
				Password: password,
				TMin: tMin,
				TMax: tMax,
			}

			file, _ := json.MarshalIndent(user, "", " ")

			_ = ioutil.WriteFile("user.txt", file, 0644)

		default:

	}

}
