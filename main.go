package main

import (
	// "fmt"
	// "runtime"
	// "github.com/diego0096/go-structures/variables"
	// "fmt"

	// "github.com/diego0096/go-structures/ejercicios"
	// "github.com/diego0096/go-structures/files"
	// "github.com/diego0096/go-structures/funciones"
	// "github.com/diego0096/go-structures/arreglos_slices"
	// "github.com/diego0096/go-structures/mapas"
	// "github.com/diego0096/go-structures/users"
	// "github.com/diego0096/go-structures/interfaces"
	e "github.com/diego0096/go-structures/ejer_interfaces"
	"github.com/diego0096/go-structures/modelos"
	// "github.com/diego0096/go-structures/teclado"
	// "github.com/diego0096/go-structures/iteraciones"
)

func main() {
	// variables.MuestroEnteros()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoATexto(158)
	// fmt.Println(estado)
	// fmt.Println(texto)

	// if os := runtime.GOOS; os == "Linux." || os == "OS X." { // runtime.GOOS se utiliza para obtener el sistema operativo que se esta usando
	// 	fmt.Println("Esto no es Windows, es ", os)
	// } else {
	// 	fmt.Println("Esto si es Windows", os)
	// }

	// switch os := runtime.GOOS; os {
	// case "linux":
	// 	fmt.Println("Esto es Linux")
	// case "darwin":
	// 	fmt.Println("Esto es Darwin")
	// default:
	// 	fmt.Printf("%s \n", os) // %s significa que el valor se va a formatera como string
	// }

	// numero, text := ejercicios.Ejercicio1("200")
	// fmt.Println(numero)
	// fmt.Println(text)

	// teclado.IngresoNumeros()

	// iteraciones.Iterar()

	// ejercicios.Multiply()

	// fmt.Println(ejercicios.Multiply())

	// files.SumaTabla()
	// files.LeoArchivo()

	// funciones.Calculos()
	// funciones.LlamarClosure()

	// arreglos_slices.Capacidad();

	// mapas.MostrarMapas();

	// users.AltaUsuario()

	Pedro:=new(modelos.Hombre)
	e.HumanosRespirando(Pedro)

	Maria:=new(modelos.Mujer)
	e.HumanosRespirando(Maria)

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
