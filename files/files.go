package files

import (
	"bufio"
	"fmt"

	"os"

	"github.com/freneklopez/godesde0/ejercicios"
)

var filename string = "./files/txt/tabla.txt"

func GrabaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	archivo, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error  al crear el archivo. " + err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumaTabla() {
	var texto string = ejercicios.TablaMultiplicar()
	if !Append(filename, texto) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(filen, texto string) bool {
	arch, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el Append" + err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el WriteString" + err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(filename)
	if err != nil {
		fmt.Println("Hubo un error leyendo el archivo " + err.Error())
		return
	}

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
