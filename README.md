# Saludos en GO

Este paquete es una prueba de un curso de go donde obtiene distintos saludos personales

## Instalacion

Ejecuta el siguiente comando para instalar el paquete:

```bash
go get -u github.com/Inhakinho/greetings
```

## Uso
Aqui un ejemplo de codigo para usar el paquete en tu code:

```go
package main

import (
	"fmt"

	"github.com/Inhakinho/greetings"
)

func main() {
	message, err := greetings.Hello("Iñaki")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
```

Este ejemplo impora el paquete y llama a la funcion Hello con un saludo personalizado. Si ocurre un error, se imprime el error.