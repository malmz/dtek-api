package model

type Allergen struct {
	Code     string `json:"code"`
	ImageUrl string `json:"imageUrl"`
}

type LunchMenu struct {
	Resturant string          `json:"resturant"`
	Campus    string          `json:"campus"`
	Date      string          `json:"date"`
	Language  string          `json:"language"`
	Items     []LunchMenuItem `json:"items"`
}

type LunchMenuItem struct {
	Title        string     `json:"title"`
	Body         string     `json:"body"`
	Preformatted bool       `json:"preformatted"`
	Allergens    []Allergen `json:"allergen"`
	Emission     float64    `json:"emission"`
	Price        string     `json:"price"`
}
