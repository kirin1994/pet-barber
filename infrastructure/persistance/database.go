package persistance

import "github.com/kirin1994/HappyPets/domain/entities"

//Database interface
type Database interface {
	GetCompanies() []entities.Company
	GetCompany(id int) *entities.Company
}
