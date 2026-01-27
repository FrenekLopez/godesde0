package ejercicios

import (
	"strconv"
)

func Conversion(texto string) (int, string) {
	numero, err := strconv.Atoi(texto)

	// Para manejar errores se emplea la palabra nil.
	if err != nil {
		return 0, "Existe un error" + err.Error()
	}

	if numero > 100 {
		return numero, "El numero es mayor a 100"
	} else {
		return numero, "El numero es menor a 100"
	}

}
