package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error
var texto string

func Multiply() string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ingrese un número")
		if scanner.Scan() {
			numero, err = strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("El dato ingresado es incorrecto")
				continue
			} else {
				break
			}
		}
	}
	for i := 1; i <= 10; i++{
		texto += fmt.Sprintf("%d x %d = %d \n", numero, i, i * numero)
	}
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println("\n", numero*i)
	// }	
	return texto
}
