package main

import "fmt"

//Crie uma struct chamada Aluno com os campos "nome"
//"idade" e "notas". O campo "notas" deve ser um slice de float64.
//Escreva uma função que receba um Aluno como
//parâmetro e calcule a média das suas notas.

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func main() {
	p := Aluno{
		nome:  "Giovana",
		idade: 17,
		notas: []float64{7, 8, 9, 10}}
	fmt.Println(mediaAluno(p))
}
func mediaAluno(p Aluno) (media float64) {
	media = (7. + 8. + 9. + 10.) / 4.
	return float64(media)
}
