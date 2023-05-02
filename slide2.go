package main

import "fmt"

//Crie uma struct chamada Livro com os campos "título"
//"autor" e "ano de publicação".
//Escreva uma função que receba um Livro como parâmetro e
//imprima suas informações.

type Livro struct {
	titulo string
	autor  string
	ano    int
}

func main() {
	p := Livro{titulo: "Maus",
		autor: "Art Spiegelman",
		ano:   1980}
	imprimaLivro(p)

}

func imprimaLivro(p Livro) {
	fmt.Println("Titulo:", p.titulo)
	fmt.Println("Autor:", p.autor)
	fmt.Println("ano:", p.ano)
}
