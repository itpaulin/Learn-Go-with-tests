package main

import "fmt"

const prefixHello = "Olá, "

func Ola(name string) string {
	return prefixHello + name
}
func main() {
	fmt.Println(Ola("mundo"))
}
