package queries

import (
	"github.com/kirin1994/HappyPets/domain/entities"
)

func (q *Queries) GetCompanies() []entities.Company {
	return q.Database.GetCompanies()
}
