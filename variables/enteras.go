package variables

import "fmt"

func MuestroEnteros() {
	var intcomun int
	intde32 := int32(10) // asignaciòn
	intde64 := int64(100)

	fmt.Println("incomun = ", intcomun)
	fmt.Println("intde32 = ", intde32)
	fmt.Println("intde64 = ", intde64)
}