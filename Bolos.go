package main

func CalcularPuntuacion(partida [20]int) int{

 resultado := 0  

  for i := 0; i < len(partida); i++ { 

  	if partida[i] == 10 && i%2 == 0 {
  		resultado += partida[i]  
 		resultado = resultado + partida[i+1] + partida[i+2]
  	} else {
  		resultado += partida[i] 
  	}
  }

 return resultado  
}