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
		fmt.Println("ingrese que desea realizar\n suma : su \n resta: res\n multiplicacion: multi\n  divicion: div \nsalir : sal")
		fmt.Scanln(&r)
		switch r {
		case "su", "suma":
			{
				fmt.Println("bienvenido ingrese numeros")
				for iteraciones == 0 { //for-while interno case
					fmt.Scanln(&r)

					if r != "sum" && r != "suma" && r != "su" {
						i, err := strconv.Atoi(r)
						if err != nil {
							fmt.Println("no podemos almacenar el dato ")
							iteraciones++
						}
						fmt.Println("+")
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

				for iteraciones == 0 {
					if r != "resta" && r != "re" && r != "r" {
						fmt.Println("ingrese minuendo")
						fmt.Scanln(&r)

						minuendo, err := strconv.Atoi(r)

						if err != nil {
							fmt.Println("no se admite el dato")
							iteraciones++
						} else {
							fmt.Println("-")
							fmt.Scanln(&r)
							sustraendo, er := strconv.Atoi(r)
							if er != nil {
								fmt.Println("no se admite el dato")
								iteraciones++
							}
							//utilizamos la variable suma para no tener que decalarar mas variables
							suma = minuendo - sustraendo
							fmt.Println("la resta de los numeros es", suma)
						}

					}
					iteraciones++ //for-while interno
				}
				iteraciones = 0 //for-while externo
			}
		case "multi", "multiplicacion", "m":

			fmt.Println("ingrese 1o. factor")
			fmt.Scanln(&r)

			if r != "m" && r != "multi" && r != "multiplicacion" {

				fac1, err := strconv.Atoi(r)

				if err != nil {
					fmt.Println("dato no valido vuelva a intentarlo")
				} else {
					fmt.Println("*")            //imprimimos el signo *
					fmt.Scanln(&r)              //escaneamos la entrada
					fac2, er := strconv.Atoi(r) //la convertimos a entero

					if er != nil {
						fmt.Println("dato no valido vuelva a intentarlo")
					} else {
						suma = fac1 * fac2
						fmt.Println("la multiplicacion de los dos numeros es", suma)
					}
				}

			}
		case "div", "divicion", "d":
			for {
				fmt.Println("Que tipo de divicion desea realizar\n\t\t flotante : f \n\t\t entera : e\n\t\t salir : sal")
				fmt.Scanln(&r)

				if r == "f" || r == "e" {
					//declaramos dos variables de punto flotante para las diviciones
					//var divisor, dividendo float32
					var verificador bool

					if r == "f" {
						verificador = false
					} else {
						verificador = true
					}

					fmt.Println("ingrese  numero a dividir")
					fmt.Scanln(&r)

					dividendo, err := strconv.ParseFloat(r, 32)

					if err != nil {
						fmt.Println("dato no valido intente otra vez")
					}
					fmt.Println("ingrese divisor")
					fmt.Scanln(&r)

					divisor, err := strconv.ParseFloat(r, 32)

					if err != nil {
						fmt.Println("dato no valido intente otra vez")
					}

					if verificador == false {
						fmt.Println("la divicion entre ", dividendo, "  y  ", divisor, " es ", dividendo/divisor)
					} else {
						fmt.Println("la divicion entre ", int(dividendo), "  y  ", int(divisor), " es ", int(dividendo/divisor))
					}

				} else if r == "salir" || r == "sal" {
					break
				} else {
					fmt.Println("dato no valido intete otra vez")
				}
			}

		case "salir", "sal":
			fmt.Println("gracias por utilizar el programa... :=) ")
			iteraciones++
		default:
			fmt.Println("los datos ingresados no son validos por favor buelva a intentarlo")
		}
	}

}
