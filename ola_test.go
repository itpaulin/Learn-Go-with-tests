package main

import "testing"

func TestOla(t *testing.T) {

	verifyMessage := func(t testing.TB, result, expected string) {
		t.Helper() // deixando explicito que é um metodo auxiliar
		if result != expected {
			t.Errorf("resultado '%q', esperado '%q'", result, expected)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		result := Ola("paulin")
		expected := "Olá, paulin"

		verifyMessage(t, result, expected)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		result := Ola("")
		expected := "Olá, mundo"

		verifyMessage(t, result, expected)
	})
}
