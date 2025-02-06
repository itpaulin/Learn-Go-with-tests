package main

import "testing"

func TestOla(t *testing.T) {
	result := Ola("Paulin")
	expected := "Olá, Paulin"

	if result != expected {
		t.Errorf("Result '%s', expected '%s'", result, expected)
	}
}
