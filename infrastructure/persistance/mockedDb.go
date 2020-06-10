package persistance

import "github.com/kirin1994/HappyPets/domain/entities"

var companies = []entities.Company{
	{ID: 1, Name: "Moj Pupil"},
	{ID: 2, Name: "New Jork"},
	{ID: 3, Name: "Happy pets"},
	{ID: 4, Name: "Four legs friends"},
}

type MockedDb struct{}

func (db *MockedDb) GetCompanies() []entities.Company {
	return companies
}

func (db *MockedDb) GetCompany(id int) *entities.Company {
	for _, c := range companies {
		if c.ID == id {
			return &c
		}
	}

	return nil
}
