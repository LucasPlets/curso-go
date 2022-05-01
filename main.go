package main

import (
	"fmt"
	"github.com/eminetto/curso-go/festas"
)

func main() {
	casamento := festas.CalculaCasorio(festas.Parametros{
		Adultos: 200,
		Criancas: 30,
		NaoAlcoolicos: 0,
		Duracao: 8,
		Bar: true,
	})
	imprimeCasorio(casamento)
}

func imprimeCasorio(c festas.Casorio){
	un := "l"
	fmt.Printf("Total de pessoas: %d\n", c.TPessoas)
	fmt.Printf("Total de Agua: %d %s\n", c.TAgua, un)
	fmt.Printf("Total de Refrigerante: %d %s\n", c.TRefri, un)
	fmt.Printf("Total de Suco: %d %s\n", c.TSuco, un)
	fmt.Printf("Total de Choop: %d %s\n", c.TChoop,un)
}
