package contascorrente

type ContaCorrente struct {
	Titular       Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorSaque
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
		c.Saldo += valorDeposito
		return "Deposito realizado com sucesso", c.Saldo
	} else {
		return "Valor informado é invlálido", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia < c.Saldo && valorTransferencia > 0 {
		c.Saldo -= valorTransferencia
		contaDestino.Depositar(valorTransferencia)
		return true
	} else {
		return false
	}

}
