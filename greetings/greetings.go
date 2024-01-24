package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string, age string) (string, error) {
	if name == "" || age == "" {
		return "", errors.New("input name and age")
	}

	message := fmt.Sprintf(randomGreetings(), name, age)
	return message, nil
}

func Hellos(names []string, ages []string) (map[string]string, error) {
	if len(names) != len(ages) {
		return nil, errors.New("name must have corresponding age")
	}

	messages := make(map[string]string)

	for idx, name := range names {
		message, err := Hello(name, ages[idx])
		if err != nil {
			return nil, err
		}
		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
}

func randomGreetings() string {
	formats := []string{
		"Hi, %v. Welcome! You are %v years old",
		"Great to see you, %v! you look like you are %v years old",
		"Hail, %v! Well met the %v year old boss!",
	}

	return formats[rand.Intn(len(formats))]
}
