package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SolicNumero() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Ingrese el numero : ")

		if !scanner.Scan() {
			fmt.Println("Error al ingresar el numero")
			return
		}
		numero, err := strconv.Atoi(scanner.Text())

		if err != nil {
			fmt.Println("Debes ingresar un numero valido")
			continue
		} else {
			for i := 1; i <= 10; i++ {
				resultado := numero * i
				fmt.Printf("%d x %d = %d \n", numero, i, resultado)
			}
		}

		break
	}

}
