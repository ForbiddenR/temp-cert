package main

import (
	"fmt"

	"github.com/ForbiddenR/temp-cert/fcert"
)

func main() {
	sig, err := fcert.CalculateSignature("test.key", "README.md")
	// sig, err := fcert.CalculateSignature("test.key", "README.md")

	if err != nil {
		panic(err)
	}
	fmt.Println("sig", sig, len(sig))
}
