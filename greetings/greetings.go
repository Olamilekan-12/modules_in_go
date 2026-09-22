package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		greet, err := Hello(name)

		if err != nil {
			return nil, err
		}
		messages[name] = greet
	}
	return messages, nil
}

func randomFormat() string {
	format := []string{
		"Hi, %v. Welcome",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	randomGreeting := format[rand.Intn(len(format))]
	return randomGreeting
}
