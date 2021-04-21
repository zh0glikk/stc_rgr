package main

import (
	"fmt"
)

type User struct {
	Login string `json:"login"`
	Password string `json:"password"`

	TMin []float64 `json:"t_min"`
	TMax []float64 `json:"t_max"`
}

func (u User) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n", u.Login, u.Password, u.TMin, u.TMax)
}
