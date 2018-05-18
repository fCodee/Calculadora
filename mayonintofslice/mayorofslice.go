package main

import (
	"fmt"
)

func min(a []int) int {
	temp := 0

	for i := 0; i < len(a); i++ {

		for p := i + 1; p < len(a); p++ {

			for j := 0; j < len(a); j++ {

				if a[i] > a[p] {
					temp = a[i]
					a[i] = a[p]
					a[p] = temp
				}
			}

		}
	}
	fmt.Println("slice despues de la funcion min : ", a)
	return a[0]
}

func max(a []int) int {
	temp := 0

	for i := 0; i < len(a); i++ {

		for p := i + 1; p < len(a); p++ {

			for j := 0; j < len(a); j++ {
				//lo unico que cambia con la funcion min es ell ">" por"<"
				if a[i] < a[p] {
					//this line do the same the three lines bellow
					//temp,a[i],a[p]=a[i],a[p],temp
					temp = a[i]
					a[i] = a[p]
					a[p] = temp
				}
			}

		}
	}
	fmt.Println("slice despues de la funcion max : ", a)
	return a[0]
}

func main() {
	slice := make([]int, 10)
	var input int

	fmt.Println("ingrese 10 numeros")

	for i := 0; i < 10; i++ {
		fmt.Scanln(&input)
		slice[i] = input
	}

	fmt.Println("los numeros ingresados son", slice)
	minimo := min(slice)
	fmt.Println("el menor del slice anterior es :", minimo)
	maximo := max(slice)
	fmt.Println("el mayor del slice anterior es :", maximo)
	//	max(slice)

}
