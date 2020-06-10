package main

import (
	"log"
	"net/http"

	"github.com/kirin1994/HappyPets/infrastructure/persistance"
	"github.com/kirin1994/HappyPets/ui"
	_ "github.com/lib/pq"
)

func main() {
	db := persistance.CreatePSQLConnection()
	s := ui.CreateNewServer(db)
	ui.AddRoutes(s)
	log.Panic(http.ListenAndServe(":8080", s.Handler))
}
