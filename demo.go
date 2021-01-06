package main

import (
	"fmt"
	"log"
	"wu/greetings"
)

func main1() {
	names := []string{"adowo", "yunna", "doumi"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
