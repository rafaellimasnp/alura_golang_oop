package main

import (
	"clientes/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBOleto float64) {

	conta.Sacar(valorDoBOleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaDoDenis := contas.ContaPoupanca{}
	contaDaPati := contas.ContaCorrente{}

	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)

	contaDaPati.Depositar(500)
	PagarBoleto(&contaDaPati, 400)

	fmt.Println(contaDoDenis)
	fmt.Println(contaDaPati)

}
