package main

import (
	"example/greetings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"994AK", "66a", "88g"}

	message, err := greetings.Hello("Gladys")
	messages, errs := greetings.Hellos(names)

	if errs != nil {
		log.Fatalln(errs)
	}

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(message)
	fmt.Println(messages)

}
