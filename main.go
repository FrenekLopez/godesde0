package main

import (
	"fmt"

	"github.com/freneklopez/godesde0/ejercicios"
)

func main() {
	/* variables.MuestroEnteros()
	estado, texto := variables.ConviertoaTexto(100)
	fmt.Println(estado)
	fmt.Println(texto) */

	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows", os)
	} else {
		fmt.Println("Esto es Windows", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)

	}*/

	texto, numero := ejercicios.Conversion("110")
	fmt.Printf("%d \n", texto)
	fmt.Printf("%s", numero)

}
