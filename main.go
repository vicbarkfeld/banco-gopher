package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {
	clienteVic := clientes.Titular{"Vic", "123.123.122.12", "Desenvolvedora Go"}
	contaDaVic := contas.ContaCorrente{clienteVic, 152, 123552, 1005.5}
	fmt.Println(contaDaVic)
}
