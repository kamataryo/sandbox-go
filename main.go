package main

import (
	"fmt"

	cowsay "github.com/Code-Hex/Neo-cowsay/v2"
)

func main() {
	fmt.Println("Hello, World!")
	say, err := cowsay.Say("Hello, World!")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(say)
}
