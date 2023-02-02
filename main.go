package main

import (
	"alura_golang_oop/clientes"
	c "alura_golang_oop/contas"
	"fmt"
)

func main() {
	contaRafael := c.ContaCorrente{Titular: clientes.Titular{Nome: "Rafael",
		CPF:       "0258789",
		Profissao: "Programador"}, NumeroAgencia: 123132, NumeroConta: 1221, Saldo: 1000}

	clienteMitali := clientes.Titular{"Mitali", "028875564", "Fisioterapeuta"}

	contaMitali := c.ContaCorrente{clienteMitali, 12313, 1222, 1000}

	fmt.Println(contaRafael)

	fmt.Println("\n")

	fmt.Println(contaMitali)

}
