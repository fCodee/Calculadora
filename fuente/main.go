package main

import (
	"fmt"
)

func main() {
	var (
		r           string
		s           []byte
		iteraciones int
	)
	fmt.Println("ingrese que desea realizar\n, suma : su, \n, resta: res,\n, multiplicacion" +
		+" divicion: div \n")
	fmt.Scanln(&r)

	for iteraciones == 0 {
		switch r {

		case salir:
			{
				fmt.Println("gracias por utilizar el programa... :=) ")
				iteraciones++
			}
		default:
			fmt.Println("los datos ingresados no son validos por favor buelva" +
				+"a intentarlo")
		}
	}

}
