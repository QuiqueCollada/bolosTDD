package main

import (

	"fmt"
)

type Ronda struct {
	Intento1 int
	Intento2 int
}

func CalcularPuntuacion(partida [10]Ronda) int{

	var resultado int
	for i := 0; i < len(partida); i++ {

		resultado += partida[i].Intento1 + partida[i].Intento2

		if partida[i].Intento1 == 10 {
		
			resultado = resultado + partida[i+1].Intento1 + partida[i+1].Intento2

		} else if partida[i].Intento1 + partida[i].Intento2 == 10 {

			resultado = resultado + partida[i+1].Intento1

		} else if partida[i].Intento1 + partida[i].Intento2 > 10 {

			return -1
			fmt.Println("Ha introducido mal la puntuación. Roda/s con más de 10 puntos.")
			break
		} else if partida[i].Intento1 < 0 || partida[i].Intento2 < 0 {

			return -1
			fmt.Println("Ha introducido mal la puntuación. Intento/s menor que 0")
			break
		}
	}	
	return resultado
}