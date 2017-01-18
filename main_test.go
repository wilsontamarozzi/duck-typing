package main

import(
	"testing"
)

func TestGetIP(t *testing.T) {

	ipExpected := "191.8.165.212"

	result := GetIP()

	if ipExpected != result {
		t.Error("Esperado %s, recebido %s", ipExpected, result)
	}
}