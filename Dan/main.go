package main

import (
	"fmt"
	"math/rand"
	"time"
)

const lowerletters = "abcdefghijklmnopqrstuvwxyz"
const upperletters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers = "0123456789"
const specials = "!@#$%^&*()-=_+"

type character string
type source rand.Source

//PWPolicy defines the password policy.
type PWPolicy struct {
	Lowers   int
	Uppers   int
	Numbers  int
	Specials int
}

//Generate a password based on a policy.
func Generate(conf PWPolicy) {
	Source := rand.NewSource(time.Now().UnixNano())
	var allowed character
	allowed = lowerletters + upperletters + numbers + specials
	r := rand.New(Source)
	i := r.Intn(len(allowed))
	getcharacter(allowed, Source)
}

func getcharacter(c character, s source) {
	for character := 0; character < (len(allowed)); character++ {
		r := rand.New(s)
		i := r.Intn(len(allowed))
		fmt.Println(string(allowed[i]))
	}
}
