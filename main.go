package main

import (
	c "alura_golang_oop/contas"
	"fmt"
)

func main() {
	contaRafael := c.ContaCorrente{"Rafael Lima", 1189, 1234, 1000}
	contaMitali := c.ContaCorrente{"MItali", 1189, 1234, 1000}

	status := contaRafael.Transferir(300, &contaMitali)

	fmt.Println(contaRafael)
	fmt.Println(contaMitali)

	fmt.Println(status)

}
