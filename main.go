package main

import (
	"github.com/freneklopez/godesde0/defer_panic"
)

/*
// Version SIN puntero → No cambia el valor original

	func incrementarMal(n int) {
		n++
	}

// Version CON puntero → SI cambia el valor original

	func incrementarBien(n *int) {
		*n++ // modifica el valor en la memoria original
	}

// Punteros a structs

	type Persona struct {
		Nombre string
		Edad   int
	}

	func cumpleaños(p *Persona) {
		p.Edad++ //No necesitas *p.Edad Go lo "desreferencia" autenticamente.
	}
*/

func main() {

	//1
	/* variables.MuestroEnteros()
	estado, texto := variables.ConviertoaTexto(100)
	fmt.Println(estado)*/
	//fmt.Println("texto")

	//2
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

	//3
	/*texto, numero := ejercicios.Conversion("110")
	fmt.Printf("%d \n", texto)
	fmt.Printf("%s", numero)

	teclado.IngresoNumero()

	iteraciones.Iterar()

	ejercicios.SolicNumero()
	*/
	//fmt.Println(ejercicios.TablaMultiplicar())
	//files.GrabaTabla()
	//files.SumaTabla()
	//files.LeoArchivo()
	//funciones.Calculos()
	//funciones.Exponencia(5)

	/*//Ejemplos sobre como utilizar los punteros en Go
	edad := 25
	incrementarMal(edad)
	fmt.Println("Despues de incrememntarMaL:", edad) // Sigue siendo 25

	incrementarBien(&edad)
	fmt.Println("Despues de incrementarBien:", edad) //ahora es 26

	Eric := Persona{Nombre: "Eric", Edad: 30}
	cumpleaños(&Eric)
	fmt.Printf("%s ahora tiene %d años", Eric.Nombre, Eric.Edad)
	// Eric ahora tiene 31
	*/
	//arreglos_slices.MuestroArreglos()
	//arreglos_slices.MuestroSlice()
	//arreglos_slices.Capacidad()
	//mapas.MostrarMapas()
	//users.AltaUser()

	/*pedro := new(modelos.Hombre)
	e.HumanoRespirando(pedro)

	maria := new(modelos.Mujer)
	e.HumanoRespirando(maria)*/
	//defer_panic.VemosDerfer()
	defer_panic.EjemploPanic()
}
