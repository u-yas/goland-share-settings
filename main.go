package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello worldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworld")

	_ = errors.New("hello world")

	a, e := os.ReadDir(".")

	fmt.Println(a, e)
}
