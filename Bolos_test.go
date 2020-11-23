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

	expected := 14
	got := resultado

	if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_strike2(t *testing.T) {

	partida := [20]int{0,0,10,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

	resultado := CalcularPuntuacion(partida)

	expected := 12
	got := resultado

	if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_spare(t *testing.T) {

	partida := [20]int{0,10,1,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

	resultado := CalcularPuntuacion(partida)

	expected := 13
	got := resultado

	if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

func Test_spare_y_strike(t *testing.T) {

	partida := [20]int{0,10,10,0,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0}

	resultado := CalcularPuntuacion(partida)

	expected := 38
	got := resultado

	if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}

