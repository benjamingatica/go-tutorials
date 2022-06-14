package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := `clave123`
	cryptedPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(pass)
	fmt.Println(cryptedPass)
	fmt.Println(string(cryptedPass))

	clavePassword := `clave123`
	err = bcrypt.CompareHashAndPassword(cryptedPass, []byte(clavePassword))

	if err != nil {
		fmt.Println(err)
		return
	} 

	fmt.Println("correct pass")
}
