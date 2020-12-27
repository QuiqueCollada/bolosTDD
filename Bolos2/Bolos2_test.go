package main

import (
	"testing"
)

func Test_suma_puntos_partida_cualquiera(t *testing.T) {

	r1 := Ronda{Intento1: 1, Intento2:0,}
	r2 := Ronda{Intento1: 0, Intento2:0,}

	partida := [10]Ronda{r1,r2,r2,r2,r2,r2,r2,r2,r2,r2}

	resultado := CalcularPuntuacion(partida)

	expected := 1
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_suma_puntos_partida_cualquiera2(t *testing.T) {

	r1 := Ronda{Intento1: 1, Intento2:1,}
	r2 := Ronda{Intento1: 0, Intento2:0,}

	partida := [10]Ronda{r1,r2,r2,r2,r2,r2,r2,r2,r2,r2}

	resultado := CalcularPuntuacion(partida)

	expected := 2
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_suma_puntos_strike(t *testing.T) {

	r1 := Ronda{Intento1: 10, Intento2:0,}
	r2 := Ronda{Intento1: 0, Intento2:0,}
	r3 := Ronda{Intento1: 2, Intento2:0,}

	partida := [10]Ronda{r1,r3,r2,r2,r2,r2,r2,r2,r2,r2}

	resultado := CalcularPuntuacion(partida)

	expected := 14
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_suma_puntos_strike2(t *testing.T) {

	r1 := Ronda{Intento1: 10, Intento2:0,}
	r2 := Ronda{Intento1: 0, Intento2:0,}
	r3 := Ronda{Intento1: 3, Intento2:6,}

	partida := [10]Ronda{r1,r3,r3,r2,r2,r2,r2,r2,r2,r2}

	resultado := CalcularPuntuacion(partida)

	expected := 37
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_suma_puntos_spare(t *testing.T) {

	r1 := Ronda{Intento1: 5, Intento2:5,}
	r2 := Ronda{Intento1: 0, Intento2:0,}
	r3 := Ronda{Intento1: 2, Intento2:0,}

	partida := [10]Ronda{r1,r3,r2,r2,r2,r2,r2,r2,r2,r2}

	resultado := CalcularPuntuacion(partida)

	expected := 14
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_suma_puntos_spare2(t *testing.T) {

	r1 := Ronda{Intento1: 4, Intento2:6,}
	r2 := Ronda{Intento1: 0, Intento2:0,}
	r3 := Ronda{Intento1: 3, Intento2:6,}

	partida := [10]Ronda{r1,r3,r3,r2,r2,r2,r2,r2,r2,r2}

	resultado := CalcularPuntuacion(partida)

	expected := 31
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_más_de_10_puntos_ronda(t *testing.T) {

	r1 := Ronda{Intento1: 6, Intento2:6,}
	r2 := Ronda{Intento1: 0, Intento2:0,}

	partida := [10]Ronda{r1,r2,r2,r2,r2,r2,r2,r2,r2,r2}

	resultado := CalcularPuntuacion(partida)

	expected := -1
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_intento_con_puntuación_negativa(t *testing.T) {

	r1 := Ronda{Intento1: -3, Intento2:6,}
	r2 := Ronda{Intento1: 0, Intento2:0,}

	partida := [10]Ronda{r1,r2,r2,r2,r2,r2,r2,r2,r2,r2}

	resultado := CalcularPuntuacion(partida)

	expected := -1
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}
