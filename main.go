package main

import (
	"fmt"

	c "github.com/rafaellimasnp/alura_golang_oop/tree/main/contas"
)

func main() {
	contaRafael := c.ContaCorrente{"Rafael Lima", 1189, 1234, 1000}
	contaMitali := c.ContaCorrente{"MItali", 1189, 1234, 1000}

	status := contaRafael.Transferir(300, &contaMitali)

	fmt.Println(contaRafael)
	fmt.Println(contaMitali)

	fmt.Println(status)

}
