package entities

//Company entity
type Company struct {
	ID   int    `json:"id"`
	Name string `json:"company_name" db:"company_name"`
}
