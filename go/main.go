package main

import (
	"fmt"

	"math/rand"
	"os"

	"time"
)

var opcao int
var nome string
var sobrenome string
var email string
var senha1 string
var senha2 string
var cfrCodigoDeAcesso int
var acsLogin = make(map[string]string)

func main() {

	rand.Seed(time.Now().UnixNano())
	for {
		exibeNome()
		exibePossibilidades()
		fmt.Scan(&opcao)
		fmt.Println()

		switch opcao {
		case 1:
			fmt.Println("Criando conta...")
			fmt.Println()
			criarContaACS()
		case 2:
			fmt.Println("Realizando login...")
			fmt.Print(email)
			fmt.Print(acsLogin)
			fmt.Println()
			fazLogin()

		case 3:
			fmt.Println("Saindo do programa...")
			os.Exit(0)

		default:
			fmt.Println("Opção inválida.")
		}
	}

}

func exibeNome() {

	fmt.Println(`
░██████╗░█████╗░██████╗░██████╗░███████╗██╗░░░██╗██╗██████╗░░█████╗░░██████╗  ░█████╗░░█████╗░░██████╗
██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔════╝██║░░░██║██║██╔══██╗██╔══██╗██╔════╝  ██╔══██╗██╔══██╗██╔════╝
╚█████╗░██║░░██║██████╦╝██████╔╝█████╗░░╚██╗░██╔╝██║██║░░██║███████║╚█████╗░  ███████║██║░░╚═╝╚█████╗░
░╚═══██╗██║░░██║██╔══██╗██╔══██╗██╔══╝░░░╚████╔╝░██║██║░░██║██╔══██║░╚═══██╗  ██╔══██║██║░░██╗░╚═══██╗
██████╔╝╚█████╔╝██████╦╝██║░░██║███████╗░░╚██╔╝░░██║██████╔╝██║░░██║██████╔╝  ██║░░██║╚█████╔╝██████╔╝
╚═════╝░░╚════╝░╚═════╝░╚═╝░░╚═╝╚══════╝░░░╚═╝░░░╚═╝╚═════╝░╚═╝░░╚═╝╚═════╝░  ╚═╝░░╚═╝░╚════╝░╚═════╝░`)
	fmt.Println()
	fmt.Println()

}

func exibePossibilidades() {
	fmt.Println("(MENU)")
	fmt.Println("1- Criar conta")
	fmt.Println("2- Login")
	fmt.Println("3- Sair")
	fmt.Print("Digite a opção desejada: ")

}

func criarContaACS() {
	fmt.Print("Nome: ")
	fmt.Scan(&nome)

	fmt.Print("Sobrenome: ")
	fmt.Scan(&sobrenome)

	fmt.Print("Email: ")
	fmt.Scan(&email)
	_, exists := acsLogin[email]
	if exists {
		fmt.Println("Email já cadastrado no sistema")
		for exists == true {
			fmt.Println("Email: ")
			fmt.Scan(&email)
		}
	}

	fmt.Print("Digite sua senha: ")
	fmt.Scan(&senha1)

	fmt.Print("Confirme sua senha: ")
	fmt.Scan(&senha2)

	if senha1 != senha2 {
		for senha1 != senha2 {
			fmt.Println(" AS SENHA INSERIDAS SÃO DIFERENTES!")
			fmt.Print("Confirme sua senha: ")
			fmt.Scan(&senha2)
			fmt.Println()
		}
	}
	codigoDeAcesso := rand.Intn(100000)

	fmt.Println("Código de Acesso:", codigoDeAcesso)
	fmt.Print("Confirme seu código de acesso: ")
	fmt.Scan(&cfrCodigoDeAcesso)
	if cfrCodigoDeAcesso != codigoDeAcesso {
		for cfrCodigoDeAcesso != codigoDeAcesso {
			fmt.Println("O código de acesso não confere.")
			fmt.Print("Confirme seu código de acesso: ")
			fmt.Scan(&cfrCodigoDeAcesso)
		}
		fmt.Println()
	}
	fmt.Println()
	acs := make([]string, 0)
	acs = append(acs, nome, sobrenome, email, senha2)
	fmt.Println()
	fmt.Printf("Parabéns, %s. Sua conta foi criada com sucesso", nome)
	fmt.Println()
	fmt.Println("(VOLTANDO PARA O MENU...)")
	fmt.Println()

	acsLogin[email] = senha1

}

var emailLogin string
var senhaLogin string

func fazLogin() {

	fmt.Print("Digite seu email: ")
	fmt.Scan(&emailLogin)

	_, ok := acsLogin[emailLogin]
	if !ok {

		for !ok {
			fmt.Println("Email inválido")
			fmt.Print("Digite seu email: ")
			fmt.Scan(&emailLogin)
			_, ok = acsLogin[emailLogin]
		}
	}

	fmt.Print("Digite a senha: ")
	fmt.Scan(&senhaLogin)
	fmt.Println()

	if acsLogin[emailLogin] == senhaLogin {
		fmt.Println()
		fmt.Println("Login realizado com sucesso!")
		fmt.Println("(VOLTANDO PARA O MENU...)")
		fmt.Println()

	} else {
		var alternativa int
		fmt.Println("Senha inválida")
		fmt.Println("1- tentar novamente")
		fmt.Println("2- redefinir senha")
		fmt.Scan(&alternativa)

		switch alternativa {
		case 1:
			for i := 0; i < 3; i++ {
				if acsLogin[emailLogin] != senhaLogin {
					fmt.Printf("Você tem %d tentativas para inserir a senha correta \n ", (3 - i))
					fmt.Printf("Digite sua senha: ")
					fmt.Scan(&senhaLogin)
				} else {
					fmt.Println("Login realizado com sucesso!")
					fmt.Println("(VOLTANDO PARA O MENU...)")
					fmt.Println()
					break
				}

			}
		case 2:
			redefinirSenha()

		}

	}
}

func redefinirSenha() {

}
