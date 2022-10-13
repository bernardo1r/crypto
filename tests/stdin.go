package main

import (
	"fmt"
	"bernardo1r/password"
	"bernardo1r/password/internal/check"
)

func main() {

	fmt.Printf("Password:")
	password, err := password.ReadStdin()
	check.Check(err)
	fmt.Printf("\nBytes: %v\nString: %s\n", password, string(password))
}

