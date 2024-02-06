package main

import (
	"fmt"
	"log"

	"github.com/andyfusniak/base58"
)

func main() {
	s, err := base58.RandString(8)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("base58 string %s\n", s)
}
