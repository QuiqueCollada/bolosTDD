package main

import (
	"testing"
)

func TestIngredienteBorradoConExito(t *testing.T) {

	expected := 0
    got := resultado

    if expected != got {
        t.Errorf("got '%v' expected '%v'", got, expected)
    }
}
