package main

import (
	"fmt"
	"strings"
)

func main() {

	var username, password string

	validusername := "admin"
	validpassword := "password123,"

	fmt.Println("Please enter username : ")
	fmt.Scanln(&username)

	fmt.Println("Please enter password : ")
	fmt.Scanln(&password)

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if username == validusername && password == validpassword {
		fmt.Println("Login successful.")
	} else {
		fmt.Println("Login failed.")
	}

}
