package main

import "fmt"

//Crie uma struct chamada Filme com os campos "título",
//"diretor", "ano" e "avaliações". O campo "avaliações"
//deve ser um slice de inteiros representando as notas
//que o filme recebeu dos espectadores. Escreva funções
//que permitam adicionar ou remover avaliações do filme,
//calcular a média das avaliações e imprimir as
//informações do filme e sua média de avaliações.

type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

func (p Filme) adicionarAvaliacao(avaliacao int) Filme {
	p.avaliacoes = append(p.avaliacoes, avaliacao)
	return p
}

func (p Filme) removerAvaliacao(indice int) Filme {
	if indice >= 0 && indice < len(p.avaliacoes) {
		p.avaliacoes = append(p.avaliacoes[:indice], p.avaliacoes[indice+1:]...)
	}
	return p
}

func (p Filme) calcularMedia() int {
	total := 0
	for _, avaliacao := range p.avaliacoes {
		total += avaliacao
	}
	if len(p.avaliacoes) > 0 {
		return total / int(len(p.avaliacoes))
	}
	return 0
}

func (p Filme) imprimirinformação() {
	fmt.Println("Titulo:", p.titulo)
	fmt.Println("Diretor:", p.diretor)
	fmt.Println("Media:", p.calcularMedia())
}

func main() {
	p := Filme{
		titulo:     "Diabo veste prada",
		diretor:    "David Frankel",
		avaliacoes: []int{10, 10, 10, 10},
	}
	p = p.adicionarAvaliacao(10)
	p.imprimirinformação()
}
