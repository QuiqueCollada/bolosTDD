package main

func CalcularPuntuacion(partida [20]int) int{

 resultado := 0  
 	for _, v := range partida {  
 	 resultado += v  
 	}  
 return resultado  
}