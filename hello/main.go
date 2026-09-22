package main

import (
	"fmt"
	"log"
	"modules_in_go/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	names := []string{
		"Olamilekan",
		"Olayinka",
		"Bello",
	}
	msg, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
