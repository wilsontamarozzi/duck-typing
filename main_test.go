package main

import(
	"testing"
)

type ServerFake struct{}

func (s ServerFake) LoadData(url string) (string, error) {
    return "191.8.165.212", nil
}

func TestGetIP(t *testing.T) {

	ipExpected := "191.8.165.212"

	server := ServerFake{}

	result := GetIP(server)

	if ipExpected != result {
		t.Error("Esperado %s, recebido %s", ipExpected, result)
	}
}