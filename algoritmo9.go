package main

import "fmt"

//Crie uma struct chamada Aluno com os campos "nome",
//"idade" e "notas". O campo "notas" deve ser um slice
//de float64, representando as notas que o aluno tirou
//em uma determinada disciplina. Escreva funções que
//permitam adicionar ou remover notas do aluno, calcular
//a média das notas e imprimir o nome, idade e média do
//aluno.

type Estudante struct {
	nome  string
	idade int
	notas []float64
}

func (p Estudante) adicionarNota(nota float64) Estudante {
	p.notas = append(p.notas, nota)
	return p
}

func (p Estudante) removerNota(indice int) Estudante {
	if indice >= 0 && indice < len(p.notas) {
		p.notas = append(p.notas[:indice], p.notas[indice+1:]...)
	}
	return p
}

func (p Estudante) calculaMedia() float64 {
	total := 0.0
	for _, nota := range p.notas {
		total += nota
	}
	if len(p.notas) > 0 {
		return total / float64(len(p.notas))
	}
	return 0.0
}

func (p Estudante) imprimirInformacao() {
	fmt.Println("Nome:", p.nome)
	fmt.Println("Idade:", p.idade)
	fmt.Println("Media:", p.calculaMedia())

}

func main() {
	p := Estudante{
		nome:  "giovana",
		idade: 18,
		notas: []float64{7.5, 8.0, 9.5},
	}
	p = p.adicionarNota(8.5)
	p.imprimirInformacao()

	p = p.removerNota(1)
	p.imprimirInformacao()
}
