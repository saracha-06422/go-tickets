package main

import (
	"fmt"

	"github.com/saracha-06422/go-tickets/controller/routes"
	"github.com/saracha-06422/go-tickets/databases"
)

func main() {
	fmt.Println("saracha")
	databases.ConnectPostgres()

	r := routes.SetupRouter()
	r.Run(":8088")
}
