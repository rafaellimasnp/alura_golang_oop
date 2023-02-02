package main

import (
	"alura_golang_oop/clientes"
	c "alura_golang_oop/contas"
	"fmt"
)

func main() {
	contaRafael := c.ContaCorrente{Titular: clientes.Titular{Nome: "Rafael",
		CPF:       "0258789",
		Profissao: "Programador"}, NumeroAgencia: 123132, NumeroConta: 1221}

	contaRafael.Depositar(1000)

	fmt.Println(contaRafael)
	fmt.Println(contaRafael.ObterSaldo())

}
