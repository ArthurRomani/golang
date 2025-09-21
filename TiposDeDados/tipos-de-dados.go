package main

import "fmt"

func main() {
	numero := -10000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	//int32 = rune
	//int8 = byte

	var numero3 rune = 12345
	fmt.Println(numero3)

	var numero4 byte = 43
	fmt.Println(numero4)
	//Fim numeros inteiros

	var numeroReal float32 = 123.43
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 12345.43
	fmt.Println(numeroReal2)

	numeroReal3 := 12341.43
	fmt.Println(numeroReal3)
	//Fim numeros reais

	var str string = "texto"
	fmt.Println(str)

	str2 := "str2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)
	//Fim strings 

	var texto int16
	fmt.Println(texto)
	//No go todo tipo tem seu valor inicial na str e vazio nao mostra nada e no int aparece 0

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro)
}

