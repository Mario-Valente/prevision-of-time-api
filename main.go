package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Response struct {
	Extremes   []ExtremesData `json:"extremes"`
	Heights    []HeightsData  `json:"heights"`
	Disclaimer string         `json:"disclaimer"`
	Status     int            `json:"status"`
	Latitude   float64        `json:"latitude"`
	Longitude  float64        `json:"longitude"`
}

type ExtremesData struct {
	Timestamp int     `json:"timestamp"`
	Datetime  string  `json:"datetime"`
	Height    float64 `json:"height"`
	State     string  `json:"state"`
}
type HeightsData struct {
	Timestamp int     `json:"timestamp"`
	Datetime  string  `json:"datetime"`
	Height    float64 `json:"height"`
	State     string  `json:"state"`
}

func main() {
	api_key := ""
	apiUrl := "https://api.marea.ooo/v2/tides"

	parameters := url.Values{}
	// parameters.Add("x-marea-api-token", api_key)
	parameters.Add("latitude", "-16.593873")
	parameters.Add("longitude", "-39.089267")
	parameters.Add("interval", "60")
	parameters.Add("model", "FES2014")
	parameters.Add("duration", "14400")

	QueryString := parameters.Encode()
	fmt.Println(QueryString)

	url := fmt.Sprintf("%s?%s", apiUrl, QueryString)
	fmt.Println(url)

	// response, err := http.Get(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("erro ao ao criar reuisição", err)
		return
	}
	fmt.Println(req)
	req.Header.Add("x-marea-api-token", api_key)
	if err != nil {
		fmt.Println("erro ao fazer a solicitação http", err)
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao realizar a requisição:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler response: ", err)
	}

	var resposta Response

	if err := json.Unmarshal(body, &resposta); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}
	fmt.Println("Disclaimer:", resposta.Disclaimer)
	fmt.Println("Status:", resposta.Status)
	fmt.Println("Latitude:", resposta.Latitude)
	fmt.Println("Longitude:", resposta.Longitude)

	for i, extreme := range resposta.Extremes {
		fmt.Printf("Extremo %d:\n", i+1)
		fmt.Printf("  Timestamp: %d\n", extreme.Timestamp)
		fmt.Printf("  Datetime: %s\n", extreme.Datetime)
		fmt.Printf("  Altura: %f\n", extreme.Height)
		fmt.Printf("  Estado: %s\n", extreme.State)
	}
	// 	for i, height := range resposta.Heights {
	// 		fmt.Printf(" Heights%d:\n", i+1)
	// 		fmt.Printf("  Timestamp: %d\n", height.Timestamp)
	// 		fmt.Printf("  Datetime: %s\n", height.Datetime)
	// 		fmt.Printf("  Altura: %f\n", height.Height)
	// 		fmt.Printf("  Estado: %s\n", height.State)
	// 	}
}
