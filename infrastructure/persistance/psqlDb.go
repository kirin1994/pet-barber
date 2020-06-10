package persistance

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/kirin1994/HappyPets/domain/entities"
)

type PsqlDb struct {
	conn *sqlx.DB
}

func (db *PsqlDb) GetCompanies() []entities.Company {
	companies := []entities.Company{}
	err := db.conn.Select(&companies, "SELECT * FROM company")
	if err != nil {
		fmt.Println(err, "zmienic na zwracanie erroru wyzej jako 2 arg")
	}
	return companies
}

func (db *PsqlDb) GetCompany(id int) *entities.Company {
	company := entities.Company{}
	err := db.conn.Get(&company, "SELECT * FROM company WHERE id=$1", id)
	if err != nil {
		fmt.Println(err, id)
	}
	return &company
}

func CreatePSQLConnection() *PsqlDb {
	db := PsqlDb{}
	var err error
	db.conn, err = sqlx.Connect("postgres", fmt.Sprintf("user=postgres password=%s dbname=happypet sslmode=disable", os.Args[1]))
	if err != nil {
		fmt.Println(err)
	}
	db.conn.Ping()
	if err != nil {
		fmt.Println(err)
	}
	return &db
}
