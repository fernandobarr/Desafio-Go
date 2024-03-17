package main

import "fmt"

func main() {
	fmt.Println("Calculo de Kelvin para Celsius")
	fmt.Println(" ")
	fmt.Println("Digite quantos kelvin")
	var numero int
	fmt.Scan(&numero)
	numero = numero - 273
	fmt.Printf("%v Celsius", numero)

}
