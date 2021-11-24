package main

import "fmt"

func main() {
	nome := "João"
	ver := "1.0"

	fmt.Println("Nome:", nome)
	fmt.Println("Versão", ver)

	fmt.Println("1- A")
	fmt.Println("2- B")
	fmt.Println("3- C")

	var escolha int
	fmt.Scan(&escolha)

	fmt.Println("Sua escolha foi", escolha)
}
