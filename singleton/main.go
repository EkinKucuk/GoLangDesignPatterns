package main

import (
	"fmt"
)

func main() {
	go getInstance()
	fmt.Scanln()
	go getInstance()
	fmt.Scanln()
}
