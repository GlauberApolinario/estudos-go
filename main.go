package main

import "fmt"

func main() {

	/* 	Parte 1 - Variaveis -----------------------------------------------------

	   	livro := "O Arqueiro"
	   	paginas := 300
	   	status := ""

	   	if paginas < 200 {
	   		status = "livro é de boa"
	   	} else if paginas < 500 {
	   		status = "livro é razoavel"
	   	} else {
	   		status = "livro é grande"
	   	}

	   	fmt.Println("Li o livro", livro, "que tem", paginas, "paginas. O", status, "de ler") */

	//Parte 2 - For ------------------------------------------------------------

	// for i := 1; i <= 7; i++ {
	// 	fmt.Println("Dia", i, "cheguei na pagina", i*30)
	// }

	//Desafio
	//Em quantos dias você termina um livro de 300 páginas lendo 35 por dia?

	/* 	paginas := 0
	   	dias := 0

	   	for paginas < 300 {
	   		dias++
	   		paginas += 35
	   	}

	   	fmt.Println("Para ler um livro de 300 páginas, eu levarei", dias, "dias.") */

	//Parte 3 - Funções
	//dias := diasParaTerminar(300, 25)
	fmt.Println("Pata ler o Arquiro que tem 300 paginas lendo 25 páginas por dia, irei demorar", diasParaTerminar(300, 25), "dias.")
	fmt.Println("Pata ler Duna que tem 680 paginas lendo 40 páginas por dia, irei demorar", diasParaTerminar(680, 40), "dias.")
	fmt.Println("Pata ler 1984 que tem 328 paginas lendo 20 páginas por dia, irei demorar", diasParaTerminar(328, 20), "dias.")
}

func diasParaTerminar(paginas int, porDia int) int {
	dias := 0
	lidas := 0
	for lidas < paginas {
		dias++
		lidas += porDia
	}
	return dias
}
