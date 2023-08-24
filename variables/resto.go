package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string // Cadena de texto
var Estado bool // Booleano
var Sueldo float32 // Decimal
var Fecha time.Time // Tipo time

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}