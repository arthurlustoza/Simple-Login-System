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

func criandoCadastro() (string, int, float64, bool, int) {
	var nomeCliente string
	var IdadeCliente int
	var PesoCliente float64
	var AtivoCliente bool
	var SenhaCliente int

	fmt.Println("Digite o seu nome: ")
	fmt.Scanln(&nomeCliente)

	fmt.Println("Digite a sua idade: ")
	fmt.Scanln(&IdadeCliente)

	fmt.Println("Digite o seu peso: ")
	fmt.Scanln(&PesoCliente)

	fmt.Println("Plano: ", AtivoCliente)

	fmt.Println("Digite a sua senha: ")
	fmt.Scanln(&SenhaCliente)

	return nomeCliente, IdadeCliente, PesoCliente, AtivoCliente, SenhaCliente
}

func main() {

	user1 := Usuario{ //info usuario
		Nome:  "Arthur",
		Idade: 26,
		Peso:  56.7,
		Ativo: true,
		Senha: 123,
	}

	var possuiCadatro string

	fmt.Println("Você possui um cadastro ?(s/n)")
	fmt.Scanln(&possuiCadatro)

	switch possuiCadatro {
	case "s", "S":

		var senhaUsuario int
		var nomeUsuario string

		fmt.Println("Digite o seu nome:  ")
		fmt.Scanln(&nomeUsuario)

		fmt.Println("Digite a sua senha")
		fmt.Scanln(&senhaUsuario)

		if senhaUsuario == user1.Senha && nomeUsuario == user1.Nome { //sistema de verificacao de acesso
			fmt.Println("Acesso Liberado")

			fmt.Println("Suas informacoes: ")
			fmt.Println("Idade: ", user1.Idade)
			fmt.Println("Peso: ", user1.Peso)
			fmt.Println("Assinatura ativa: ", user1.Ativo)

		} else {
			fmt.Println("Acesso Negado")
		}
	case "n", "N":

		var opcCadastro string

		fmt.Println("Deseja criar seu cadastro? (s/n)")
		fmt.Scanln(&opcCadastro)
		if opcCadastro == "s" || opcCadastro == "S" {

			nomeCliente, IdadeCliente, PesoCliente, AtivoCliente, SenhaCliente := criandoCadastro()

			user2 := Usuario{ //Cliente com os dados cadastrados
				Nome:  nomeCliente,
				Idade: IdadeCliente,
				Peso:  PesoCliente,
				Ativo: AtivoCliente,
				Senha: SenhaCliente,
			}

			fmt.Printf("\nCadastro criado com sucesso: %+v\n", user2)
		}
	}

}
