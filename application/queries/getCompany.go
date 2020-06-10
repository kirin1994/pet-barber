package queries

import "github.com/kirin1994/HappyPets/domain/entities"

func (q *Queries) GetCompany(id int) *entities.Company {
	return q.Database.GetCompany(id)
}
