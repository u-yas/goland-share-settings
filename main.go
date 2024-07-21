package main

import "os"

func main() {
	a := "hogefuga"
	_, err := os.ReadDir("")
	neko(a)
	err.Error()
}

func neko(a string) {
}
