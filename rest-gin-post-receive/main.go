package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Exemplo struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	r := gin.Default()

	r.POST("/api/exemplo", func(c *gin.Context) {
		var exemplo Exemplo

		if err := c.ShouldBindJSON(&exemplo); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erro na requisição"})
			return
		}

		fmt.Println("Nome:", exemplo.Nome+", Idade:", exemplo.Idade)

		c.JSON(http.StatusOK, gin.H{"message": "Requisição POST recebida com sucesso"})
	})

	r.Run(":8080")
}
