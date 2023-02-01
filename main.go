package main

import (
	"curso-golang-alura-oop/Contas",
	"fmt"
)

func main() {
	contaRafael := Contas.ContaCorrente{"Rafael Lima", 1189, 1234, 1000}
	contaMitali := Contas.ContaCorrente{"MItali", 1189, 1234, 1000}

	status := contaRafael.Transferir(300, &contaMitali)

	fmt.Println(contaRafael)
	fmt.Println(contaMitali)

	fmt.Println(status)

}
