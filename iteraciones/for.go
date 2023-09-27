package iteraciones

import (
	"fmt"
)

func Iterar() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 100; i += 5 {
		fmt.Println(i)
	}

	for i := 0; i > 10; i -= 5 {
		fmt.Println(i)
	}

	for i := 0; i > 1; i-- {
		fmt.Println(i)
	}

	for i := 0; i > 1; i-- {
		if i == 6 {
			break // Se utiliza para terminar el ciclo en el momento indicado
		}
		fmt.Println(i)
	}

	for i := 0; i > 1; i-- {
		if i == 6 {
			continue // Se utiliza para continuar con el ciclo saltandose la accion indicada en este caso se salta el 6
		}
		fmt.Println(i)
	}
}