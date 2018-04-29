package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		r           string
		sumandos    []int
		iteraciones int
		contador    int
	)
	fmt.Println("ingrese que desea realizar\n, suma : su, \n, resta: res,\n, multiplicacion: multi,\n,  divicion: div \n")
	fmt.Scanln(&r)

	for iteraciones == 0 {
		switch r {
		case "su", "suma":
			{

				for iteraciones == 0 {
					//imprimimos una sola vez el titulo al comienzo y almacenamos un dato
					if iteraciones == 0 {
						fmt.Println("ingrese numero")
					}
					fmt.Scanln(&r)
					if r != "sum" && r != "suma" {
						i, err := strconv.Atoi(r)
						if err == nil {
							fmt.Println("no podemos almacenar el dato ")
							iteraciones++
						} else {
							sumandos[contador] = i
							contador++
						}

					} else {
						contador = 0
						for i := 0; i < len(sumandos); i++ {
							contador += sumandos[i]
						}

						fmt.Println("la suma de los numeros es \n\t : ", contador)
						fmt.Scanln()
					}

				}

				iteraciones = 0
			}
		case "salir":
			{
				fmt.Println("gracias por utilizar el programa... :=) ")
				iteraciones++
			}
		default:
			fmt.Println("los datos ingresados no son validos por favor buelva" + "a intentarlo")
		}
	}

}
