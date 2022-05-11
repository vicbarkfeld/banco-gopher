package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDaVic := contas.ContaCorrente{Titular: "Vic", Saldo: 300}
	contaDaSandra := contas.ContaCorrente{Titular: "Sandra", Saldo: 100}

	status := contaDaVic.Transferir(-200, &contaDaSandra)

	fmt.Println(status)
	fmt.Println(contaDaVic)
	fmt.Println(contaDaSandra)
}
