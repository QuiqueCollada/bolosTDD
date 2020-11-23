package main

func CalcularPuntuacion(partida [20]int) int{

 resultado := 0  
 	for _, v := range partida {  
 	 resultado += v  
 	} 

 	if partida[0] == 10 {
 		resultado = resultado + partida[1] + partida[2]
 	}

 return resultado  
}