package main

import (
	"fmt"

	"github.com/salgadoth/go-rest-api/database"
	"github.com/salgadoth/go-rest-api/routes"
)

func main() {
	database.ConectaComDB()

	fmt.Println("Iniciando o servidor REST com GO")
	routes.HandleRequest()
}
