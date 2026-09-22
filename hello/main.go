package main

import (
	"fmt"
	"modules_in_go/greetings"
)

func main() {
	msg := greetings.Hello("Olamilekan")
	fmt.Println(msg)
}
