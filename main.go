package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int = -1

	// Loop while para permanecer no código
	for opcao != 0 {
		// Menu de opções do usuário
		fmt.Println("Calculadora de Geometria Plana e Espacial")
		fmt.Println("(1) Triângulo equilátero")
		fmt.Println("(2) Retângulo")
		fmt.Println("(3) Quadrado")
		fmt.Println("(4) Círculo")
		fmt.Println("(5) Pirâmide com base quadrangular")
		fmt.Println("(6) Cubo")
		fmt.Println("(7) Paralelepipedo")
		fmt.Println("(8) Esfera")
		fmt.Println("(0) Sair")

		fmt.Print("\nDigite sua opção: ")

		switch opcao {
		case 1:

		case 2:
		case 3:
		case 4:
		case 5:
		case 6:
		case 7:
		case 8:
		}
	}
	fmt.Println(math.Pi)
}
