package main

import (
	"fmt"
	"strings"
)

//Crie uma função que receba uma string e retorne a quantidade de vogais presentes.

func main() {
	fmt.Println(vogais("paralelepipedo"))

}
func vogais(palavra string) int {
	vogais := "aeiouAEIOU"
	count := 0

	for _, carac := range palavra {
		if strings.ContainsRune(vogais, carac) {
			count++
		}
	}
	return count
}
