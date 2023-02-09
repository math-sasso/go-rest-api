package main

import (
	"fmt"

	"github.com/math-sasso/go-rest-api/database"
	"github.com/math-sasso/go-rest-api/routes"
)

func main() {

	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando o servidor Rest em Go")
	routes.HandleRequest()
}
