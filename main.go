package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		if valorSaque < 0 {
			return "Valor informado é invalido"
		}
		return "Saldo insuficiente"
	}

}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor informado é invlálido", c.saldo
	}
}

func main() {
	contaRafael := ContaCorrente{"Rafael Lima", 1189, 1234, 1000}

	resultado := contaRafael.Sacar(100)

	fmt.Println(resultado, "Saldo atual:", contaRafael.saldo)

	status, valor := contaRafael.Depositar(900)
	fmt.Println(status, valor)

}
