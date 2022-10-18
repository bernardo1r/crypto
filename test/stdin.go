package main

import (
	"fmt"
	"github.com/bernardo1r/password"
	"github.com/bernardo1r/password/internal/check"
)

func main() {

	fmt.Printf("Password:")
	password, err := password.ReadStdin()
	check.Check(err)
	fmt.Printf("\nBytes: %v\nString: %s\n", password, string(password))
}

