package model

import "time"

type Allergen struct {
	Code     string `json:"code"`
	ImageUrl string `json:"imageUrl"`
}

type LunchMenu struct {
	Name      string          `json:"name"`
	FetchedAt time.Time       `json:"fetched_at"`
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
