package main

import (
	"fmt"
	"github.com/diego0096/go-structures/variables"
	"runtime"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(158)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "Linux." || os == "OS X." { // runtime.GOOS se utiliza para obtener el sistema operativo que se esta usando
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto si es Windows", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os) // %s significa que el valor se va a formatera como string
	}
}

// var (
// 	name string
// 	age int
// 	location string
// )

// var (
// 	name1, location1 string
// 	age1 int
// )
