package main

import "fmt"

func main()  {
	fmt.Print("digite o seu nome: ")
	var nome string //declara a variavel
	fmt.Scanf("%s", &nome)  //armazena a variavel

	serie := "Logica de programação em Go" // atribuindo algo a variavel e já declarando ela
	fmt.Printf("Seja Bem vindo á serie %s, %s", serie,nome)
}

//entrada de dados