package main

import (
	"fmt"

	"github.com/bernardo1r/crypto"
	"github.com/bernardo1r/crypto/internal/check"
)

func main() {

	fmt.Printf("Password:")
	password, err := crypto.ReadStdin()
	check.Check(err)
	fmt.Printf("\nBytes: %v\nString: %s\n", password, string(password))
}
