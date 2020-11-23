package main

import (
	"testing"
)

func Test_suma_puntos_partida_cualquiera(t *testing.T) {

	partida := [20]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}

	resultado := CalcularPuntuacion(partida)

	expected := 20
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_suma_puntos_partida_cualquiera2(t *testing.T) {

	partida := [20]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,0}

	resultado := CalcularPuntuacion(partida)

	expected := 19
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_strike(t *testing.T) {

	partida := [20]int{10,0,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

	resultado := CalcularPuntuacion(partida)

	expected := 13
	got := resultado

	if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}
