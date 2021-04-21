package main

import (
	"fmt"
	"time"
)

const attemptsAmount = 5

func authorizate(user User) {
	var login string
	fmt.Print("Login: ")
	fmt.Scanln(&login)

	var n1 float64

	if login == user.Login {

		time.Sleep(time.Second / 10)

		for j := 0; j < attemptsAmount; j++ {
			pass, dur := logKey(len(user.Password))

			if pass == user.Password {
				for i := 0; i < len(dur); i++ {
					durI := float64(dur[i])

					if durI > user.TMax[i] || durI < user.TMin[i] {
						fmt.Println(durI, user.TMax[i], user.TMin[i])
						n1 += 1
						break
					}
				}
			} else {
				fmt.Println("Wrong password")
			}
		}

		fmt.Printf("P1 = %v\nP2 = %v\n", n1 / attemptsAmount, (attemptsAmount - n1) / attemptsAmount)
	} else {
		fmt.Println("Wrong login")
	}
}
