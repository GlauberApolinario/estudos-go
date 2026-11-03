package main

import "fmt"

func main() {

	/* 	Parte 1 - Variaveis

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

	//Parte 2 - For
	for i := 1; i <= 7; i++ {
		fmt.Println("Dia", i, "cheguei na pagina", i*30)
	}
}
