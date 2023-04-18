package main

import "fmt"

//Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores.

func main() {
	fmt.Print(media([]int{7, 90, 5, 4}))

}
func media(lista []int) int {
	soma := 0
	for i := 0; i < len(lista); i++ {
		soma = soma + lista[i]
	}
	media := soma / len(lista)
	return media
}
