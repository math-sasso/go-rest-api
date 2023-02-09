package main

import (
	"fmt"

	"github.com/math-sasso/go-rest-api/database"
	"github.com/math-sasso/go-rest-api/models"
	"github.com/math-sasso/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()
}
