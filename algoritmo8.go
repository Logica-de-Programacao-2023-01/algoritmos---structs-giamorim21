package main

//Crie uma struct chamada Viagem com os campos "origem",
// "destino", "data" e "preço". Escreva uma função que
//receba um slice de Viagens como parâmetro e retorne
//a viagem mais cara.

type Viagem struct {
	origem  string
	destino string
	data    int
	preco   float64
}

func viagemMaisCara(viagens []Viagem) Viagem {
	var viagemMax Viagem
	for _, viagem := range viagens {
		if viagem.preco > viagemMax.preco {
			viagemMax = viagem
		}
	}
	return viagemMax
}

func main() {

}
