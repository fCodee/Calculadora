package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		r           string
		sumandos    []int
		iteraciones int = 0
		suma        int
	)

	for iteraciones == 0 {
		fmt.Println("ingrese que desea realizar\n, suma : su, \n, resta: res,\n, multiplicacion: multi,\n,  divicion: div \n")
		fmt.Scanln(&r)
		switch r {

		case "su", "suma":
			{

				for iteraciones == 0 { //for-while interno case
					fmt.Println("+")

					fmt.Scanln(&r)
					if r != "sum" && r != "suma" && r != "su" {
						i, err := strconv.Atoi(r)
						if err != nil {
							fmt.Println("no podemos almacenar el dato ")
							iteraciones++
						}
						sumandos = append(sumandos, i)

					} else {

						for i := 0; i < len(sumandos); i++ {
							suma += sumandos[i]
						}

						fmt.Println("la suma de los numeros es \n\t : ", suma)
						iteraciones++ //for-while interno case finaliza
					}
				}

				//iteraciones = 1 //finaliza for-while externo
				iteraciones = 0
			}
		case "resta", "re", "res":
			{
				fmt.Println("Resta")

				for iteraciones == 0 {
					if r != "resta" && r != "re" && r != "r" {
						fmt.Println("ingrese minuendo")
						fmt.Scanln(&r)
						minuendo, err := strconv.Atoi(r)
						if err != nil {
							fmt.Println("no se admite el dato")
							iteraciones++
						} else {
							fmt.Println("ingrese sustraendo")
							fmt.Scanln(&r)
							sustraendo, er := strconv.Atoi(r)
							if er != nil {
								fmt.Println("no se admite el dato")
								iteraciones++
							}
							suma = minuendo - sustraendo
							fmt.Println("la resta de los numeros es", suma)
						}

					}
					iteraciones++ //for-while interno
				}
				iteraciones = 0 //for-while externo
			}
		case "salir", "sal":

			fmt.Println("gracias por utilizar el programa... :=) ")
			iteraciones++
		default:
			fmt.Println("los datos ingresados no son validos por favor buelva a intentarlo")
		}
	}

}
