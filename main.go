package main

import (
	"fmt"

	"github.com/freneklopez/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	estado, texto := variables.ConviertoaTexto(100)
	fmt.Println(estado)
	fmt.Println(texto)

}
