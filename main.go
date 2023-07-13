package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("")
	}

	fmt.Println("Port:", portString)
}
