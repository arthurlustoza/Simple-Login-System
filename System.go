package main

import (
	"fmt"
)

type Usuario struct { //usuario padrao
	Nome  string
	Idade int
	Peso  float64
	Ativo bool
	Senha int
}

type Cliente struct { //Cliente cadastrando

}

var senhaUsuario int
var nomeUsuario string
var opcCadastro string

func criandoCadastro() {

}

func main() {

	fmt.Println("Digite o seu nome:  ")
	fmt.Scanln(&nomeUsuario)

	fmt.Println("Digite a sua senha")
	fmt.Scanln(&senhaUsuario)

	fmt.Println("Deseja criar seu cadastro? (s/n)")
	fmt.Scanln(&opcCadastro)
	if opcCadastro == "s" {
		criandoCadastro()
	}

	user1 := Usuario{ //info usuario
		Nome:  "Arthur",
		Idade: 26,
		Peso:  56.7,
		Ativo: true,
		Senha: 123,
	}

	if senhaUsuario == user1.Senha && nomeUsuario == user1.Nome { //sistema de verificacao de acesso
		fmt.Println("Acesso Liberado")

		fmt.Println("Suas informacoes: ")
		fmt.Println("Idade: ", user1.Idade)
		fmt.Println("Peso: ", user1.Peso)
		fmt.Println("Assinatura ativa? ", user1.Ativo)

	} else {
		fmt.Println("Acesso Negado")
	}
}
