package main

import "fmt"

//Crie uma struct chamada Retângulo com os campos "largura" e "altura".
//Escreva uma função que receba um Retângulo como parâmetro e
//calcule a área do retângulo (área = //largura * altura)

type Retangulo struct {
	altura  int
	largura int
	area    int
}

func main() {
	p := Retangulo{largura: 10, altura: 10}
	fmt.Println(areaRetangulo(p))

}

func areaRetangulo(p Retangulo) (area int) {
	area = p.largura * p.altura
	return area
}
