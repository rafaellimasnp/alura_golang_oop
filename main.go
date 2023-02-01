package main

import (
	"fmt"

	c "https://github.com/rafaellimasnp/curso-golang-alura-oop"
)

func main() {
	contaRafael := c.ContaCorrente{"Rafael Lima", 1189, 1234, 1000}
	contaMitali := c.ContaCorrente{"MItali", 1189, 1234, 1000}

	status := contaRafael.Transferir(300, &contaMitali)

	fmt.Println(contaRafael)
	fmt.Println(contaMitali)

	fmt.Println(status)

}
