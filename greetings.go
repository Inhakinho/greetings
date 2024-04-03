package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//hello devuelve un saludo para la persona especifica

func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("nombre vacío")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Saludar a varias personas
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// funcion devuelve saludos aleatorios
func randomFormat() string {
	formats := []string{
		"Hola %v!!! Bienvenido!",
		"Que bueno verte %v",
		"Que hay de nuevo... %v",
	}
	return formats[rand.Intn(len(formats))]
}
