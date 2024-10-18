package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var longitud int
	fmt.Println("Introduce la longitud de la contraseña: ")
	fmt.Scanln(&longitud)

	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	rand.New(rand.NewSource(time.Now().UnixNano())) // Manejo automático de la semilla

	contraseña := make([]byte, longitud)
	for i := range contraseña {
		contraseña[i] = caracteres[rand.Intn(len(caracteres))]
	}
	fmt.Println("Tu contraseña generada es:", string(contraseña))
}
