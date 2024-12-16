package main

func main() {
	// Variáveis e Constantes

	/*
		Em Go existem duas formas de declarar variáveis, sendo elas de forma explicíta e implícita.
	*/

	var myName string = "Leonardo" // Explícita
	myNameTwo := "Leonardo"        // Implícita

	println("Aqui temos a forma explícita: " + myName)
	print("Aqui temos a forma implícita: " + myNameTwo)

	/*
		Os valores das váriaveis podem ser re-atribuídos para outros valores do mesmo tipo.
	*/

	myName = "José" // Correto
	//myName = 15 // Error, tipo diferente
	myNameTwo = "Carlos" // Correto
	//myNameTwo = 55 // Error, tipo diferente

	/*
		As constantes são declaradas para armazenar valores que não podem ser re-atribuidos.
	*/

	const pi float64 = 13.1415
	println(pi)

	//pi = 15 //Aqui vai acontecer um erro de re-atribuição, ou seja, que não é permitido executar essa ação

	/*
		obs: A variável atribuida de forma implícita não pode ser declarada fora do escopo de uma função
	*/
}
