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

func main() {
	result := GetIP()

	fmt.Println("IP:", result)
}

func GetIP() string {

	response, err := LoadData(URL_SERVER)

	if err != nil {
		fmt.Println("ERROR: ", err)
		panic(err)
	}

	return response
}

func LoadData(url string) (string, error) {
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