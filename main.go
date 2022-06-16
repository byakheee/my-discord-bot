package main

import (
	"flag"
	"log"
)

func main() {
	var token string

	flag.StringVar(&token, "t", "", "API Token")
	flag.Parse()
	log.Println("Hello,", token)
}
