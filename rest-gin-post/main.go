package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Exemplo struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	exemplo := Exemplo{
		Nome:  "Exemplo",
		Idade: 30,
	}

	exemplo.Nome = uuid.New().String()

	// Converte o objeto Exemplo em JSON
	jsonData, err := json.Marshal(exemplo)
	if err != nil {
		log.Fatal(err)
	}

	// Cria uma requisição POST com os dados JSON
	request, err := http.NewRequest("POST", "http://localhost:8080/api/exemplo", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")

	// Envia a requisição POST
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Exibe a resposta
	fmt.Println("Status:", response.Status)
}
