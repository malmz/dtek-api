package model

type Allergen struct {
	Code     string `json:"code"`
	ImageUrl string `json:"imageUrl"`
}

type LunchMenu struct {
	Items []LunchMenuItem `json:"items"`
}

type LunchMenuItem struct {
	Title        string     `json:"title"`
	Body         string     `json:"body"`
	Preformatted bool       `json:"preformatted"`
	Allergens    []Allergen `json:"allergen"`
	Emission     float64    `json:"emission"`
	Price        string     `json:"price"`
}
