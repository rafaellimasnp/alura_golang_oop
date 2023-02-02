package contas

import t "alura_golang_oop/clientes"

type ContaPoupanca struct {
	Titular                              t.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	} else {
		if valorSaque < 0 {
			return "Valor informado é invalido"
		}
		return "saldo insuficiente"
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor informado é invlálido", c.saldo
	}
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaPoupanca) bool {

	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		c.saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}

}
