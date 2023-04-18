package main

import (
	"fmt"
	"strings"
)

//Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings.

func main() {
	fmt.Print(frase([]string{"eu", "te", "amo"}))

}
func frase(lista []string) string {
	return strings.Join(lista, " ")
}
