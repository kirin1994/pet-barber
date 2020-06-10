package ui

import (
	"github.com/gorilla/mux"
	"github.com/kirin1994/HappyPets/application/commands"
	"github.com/kirin1994/HappyPets/application/queries"
	"github.com/kirin1994/HappyPets/infrastructure/persistance"
)

type Server struct {
	Queries  *queries.Queries
	Commands *commands.Commands
	Database persistance.Database
	Handler  *mux.Router
}

func CreateNewServer(db persistance.Database) *Server {
	serv := Server{
		Database: db,
		Queries:  &queries.Queries{Database: db},
		Commands: &commands.Commands{Database: db},
		Handler:  mux.NewRouter(),
	}

	return &serv
}
