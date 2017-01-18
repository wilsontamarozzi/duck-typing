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

func TestLoadData(t *testing.T) {

	server := Server{}

	result, err := server.LoadData("https://web.monde.com.br/ip")

	if err != nil {
		t.Error("Não era esperado o erro: %d", err)
	}

	if len(result) <= 0 {
		t.Error("Esperado um IP, recebido vázio")
	}
}