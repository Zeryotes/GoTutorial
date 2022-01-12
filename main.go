package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int = -1
	var area float64
	var perimetro float64
	var volume float64
	// Loop while para permanecer no código
	for opcao != 0 {
		// Menu de opções do usuário
		fmt.Println("")
		fmt.Println("-----------------------------------------")
		fmt.Println("Calculadora de Geometria Plana e Espacial")
		fmt.Println("-----------------------------------------")
		fmt.Println("(1) Triângulo equilátero")
		fmt.Println("(2) Retângulo")
		fmt.Println("(3) Quadrado")
		fmt.Println("(4) Círculo")
		fmt.Println("(5) Pirâmide com base quadrangular")
		fmt.Println("(6) Cubo")
		fmt.Println("(7) Paralelepipedo")
		fmt.Println("(8) Esfera")
		fmt.Println("-----------------------------------------")
		fmt.Println("(0) Sair")
		fmt.Println("-----------------------------------------")
		fmt.Print("\nDigite a sua opção: ")
		fmt.Scanln(&opcao)
		fmt.Println("")

		switch opcao {
		case 1:
			// Referência de calculos: https://www.preparaenem.com/matematica/triangulo-equilatero.htm

			// Variaveis
			var forma = "Triângulo equilátero"
			var lado float64
			var altura float64

			// Entrada de valores
			fmt.Printf("Digite o tamanho do lado do %v: ", forma)
			fmt.Scanln(&lado)

			// Calculos
			altura = (lado * math.Sqrt(3)) / 2
			area = (lado * altura) / 2
			perimetro = lado * 3

			// Print
			fmt.Printf("Área do %v: %v\n", forma, area)
			fmt.Printf("Perimetro do %v: %v\n", forma, perimetro)
		case 2:
			// Variaveis
			var forma = "Retângulo"
			var base float64
			var altura float64

			// Entrada de valores
			fmt.Printf("Digite o tamanho do base do %v: ", forma)
			fmt.Scanln(&base)
			fmt.Printf("Digite o tamanho da altura do %v: ", forma)
			fmt.Scanln(&altura)

			// Calculos
			area = base * altura
			perimetro = 2 * (base + altura)

			// Print
			fmt.Printf("Área do %v: %v\n", forma, area)
			fmt.Printf("Perimetro do %v: %v\n", forma, perimetro)
		case 3:
			// Variaveis
			var forma = "Quadrado"
			var lado float64

			// Entrada de valores
			fmt.Printf("Digite o tamanho do lado do %v: ", forma)
			fmt.Scanln(&lado)

			// Calculos
			area = math.Pow(lado, 2)
			perimetro = 4 * lado

			// Print
			fmt.Printf("Área do %v: %v\n", forma, area)
			fmt.Printf("Perimetro do %v: %v\n", forma, perimetro)
		case 4:
			// Variaveis
			var forma = "Círculo"
			var raio float64

			// Entrada de valores
			fmt.Printf("Digite o tamanho do raio do %v: ", forma)
			fmt.Scanln(&raio)

			// Calculos
			area = math.Pi * math.Pow(raio, 2)
			perimetro = 2 * math.Pi * raio

			// Print
			fmt.Printf("Área do %v: %v\n", forma, area)
			fmt.Printf("Perimetro do %v: %v\n", forma, perimetro)
		case 5:
			// Referência dos calculos: https://brasilescola.uol.com.br/matematica/area-piramide.htm

			// Variaveis
			var forma = "Pirâmide com base quadrangular"
			var arestaBase float64
			var alturaPiramide float64
			var areaBase float64
			var areaLateral float64

			// Entrada de valores
			fmt.Printf("Digite o tamanho da aresta da base da %v: ", forma)
			fmt.Scanln(&arestaBase)
			fmt.Printf("Digite a altura da %v: ", forma)
			fmt.Scanln(&alturaPiramide)

			// Calculos
			//            |                        calculo da apótema                          |   face lateral   | 4 faces
			areaLateral = ((math.Sqrt((math.Pow(arestaBase/2, 2) + math.Pow(alturaPiramide, 2))) * arestaBase) / 2) * 4
			areaBase = math.Pow(arestaBase, 2)
			area = areaBase + areaLateral
			volume = 1.0 / 3 * areaBase * areaLateral

			// Print
			fmt.Printf("Área da %v: %v\n", forma, area)
			fmt.Printf("Volume da %v: %v\n", forma, volume)
		case 6:
			// Variaveis
			var forma = "Cubo"
			var aresta float64

			// Entrada de valores
			fmt.Printf("Digite o tamanho da aresta do %v: ", forma)
			fmt.Scanln(&aresta)

			// Calculos
			area = 6 * math.Pow(aresta, 2)
			volume = math.Pow(aresta, 3)

			// Print
			fmt.Printf("Área do %v: %v\n", forma, area)
			fmt.Printf("Volume do %v: %v\n", forma, volume)
		case 7:
			// Variaveis
			var forma = "Paralelepipedo"
			var aresta1 float64
			var aresta2 float64
			var aresta3 float64

			// Entrada de valores
			fmt.Printf("Digite o tamanho da primeira aresta do %v: ", forma)
			fmt.Scanln(&aresta1)
			fmt.Printf("Digite o tamanho da segunda aresta do %v: ", forma)
			fmt.Scanln(&aresta2)
			fmt.Printf("Digite o tamanho da terceira aresta do %v: ", forma)
			fmt.Scanln(&aresta3)

			// Calculos
			area = (2 * aresta1 * aresta2) +
				(2 * aresta1 * aresta3) +
				(2 * aresta2 * aresta3)

			volume = aresta1 * aresta2 * aresta3

			// Print
			fmt.Printf("Área do %v: %v\n", forma, area)
			fmt.Printf("Volume do %v: %v\n", forma, volume)
		case 8:
			// Variaveis
			var forma = "Esfera"
			var raio float64

			// Entrada de valores
			fmt.Printf("Digite o tamanho do raio da %v: ", forma)
			fmt.Scanln(&raio)

			// Calculos
			area = 4 * math.Pi * math.Pow(raio, 2)
			volume = 4.0 / 3 * math.Pi * math.Pow(raio, 3)
			// Print
			fmt.Printf("Área do %v: %v\n", forma, area)
			fmt.Printf("Volume do %v: %v\n", forma, volume)
		case 0:
			fmt.Println("Encerrando o programa.")
		default:
			fmt.Println("Opção inválida! Por favor, digite uma opção válida.")
		}
	}
}
