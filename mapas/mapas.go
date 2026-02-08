package mapas

import (
	"fmt"
)

// Este ejercicio son para la creacion de mapas con Go
func MostrarMapas() {

	paises := make(map[string]string)
	fmt.Println(paises)

	paises["Mexico "] = "D.F"
	paises["Argentina "] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina "])

	campeonato := map[string]int{
		"Barcelona":     39,
		"Real Madril":   38,
		"Chivas":        37,
		"Bocas Juniors": 30,
	}

	fmt.Println(campeonato)

	/*for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}*/
	delete(campeonato, "Real Madril")
	fmt.Println(campeonato)

	puntaje, existe := campeonato["Barcelona"]
	fmt.Printf("El puntaje capturado es %d y el equipo existe = %t", puntaje, existe)

}
