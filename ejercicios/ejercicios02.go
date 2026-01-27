package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {
	texto := ""

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Ingrese el numero : ")

		if !scanner.Scan() {
			return ("Error al ingresar el numero")
		}
		numero, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Debes ingresar un numero valido")
			continue
		} else {
			for i := 1; i <= 10; i++ {
				resultado := numero * i
				//Se contatena texto mas texto con +=
				texto += fmt.Sprintf("%d x %d = %d \n", numero, i, resultado)
			}
		}
		break
	}
	return texto
}
