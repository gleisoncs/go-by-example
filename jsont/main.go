package main

import (
	"encoding/json"
	"fmt"
)

type Simple struct {
	Species      string
	Description  string
	Installments []Installments
}

type Installments struct {
	InstallmentId string
	Amount        string
}

func main() {
	simpleJson := `
			{
				"species":"pigeon",
				"decription":"likes to perch on rocks",
				"Installments":[
					{
						"InstallmentId":"4353-4444"
					},
					{
						"InstallmentId":"244"
					},
					{
						"InstallmentId":"355"
					}
				]
			}`
	var simple Simple
	json.Unmarshal([]byte(simpleJson), &simple)
	//	fmt.Printf("Birds : %+v", birds)

	for _, s := range simple.Installments {
		fmt.Println(s)
	}
}
