package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const URL_SERVER = "https://web.monde.com.br/ip"

type Info struct {
	Ip string `json:"ip"`
}

type ServerInterface interface {
	LoadData(url string) (string, error)
}

func main() {

	server := Server{}

	result := GetIP(server)

	fmt.Println("IP:", result)
}

func GetIP(s ServerInterface) string {

	response, err := s.LoadData(URL_SERVER)

	if err != nil {
		fmt.Println("ERROR: ", err)
		panic(err)
	}

	return response
}

type Server struct{}

func (s Server) LoadData(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return "", err
	}

	var info Info
    json.Unmarshal([]byte(string(body)), &info)

    return info.Ip, err
}