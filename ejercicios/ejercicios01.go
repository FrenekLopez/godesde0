package ejercicio01

import "strconv"

func ConvNumerico(texto string) (int, string) {
	num, err := strconv.Atoi(texto)
	if err != nil {
		return 0, "Hubo un error" + err.Error()
	}
	if num > 100 {
		return num, "El numero es mayor a 100"
	} else {
		return num, "El numero es menor a 100"

	}
}
