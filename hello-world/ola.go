package main

import "fmt"

const spanish = "espanhol"
const franchise = "frances"
const english = "ingles"
const prefixHelloPortuguese = "Olá, "
const prefixHelloSpanish = "Hola, "
const prefixHelloFranchise = "Bonjour, "
const prefixHelloEnglish = "Hello, "

func Ola(name, idioma string) string {
	if name == "" {
		name = "mundo"
	}

	return prefixodeSaudacao(idioma) + name
}
func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case spanish:
		prefixo = prefixHelloSpanish
	case franchise:
		prefixo = prefixHelloFranchise
	case english:
		prefixo = prefixHelloEnglish
	default:
		prefixo = prefixHelloPortuguese
	}
	return
}
func main() {
	fmt.Println(Ola("mundo", ""))
}
