package main

import (
	"fmt"

	"github.com/FrenekLopez/godesde0/variables"
)

func main() {
	estado, texto := variables.ConvirtiendoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
