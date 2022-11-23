package lunch

import (
	"context"

	"github.com/hasura/go-graphql-client"
)

type MenuDish struct {
	recipes []struct {
		alergens []struct {
		}
	}
}

type MenuDushOccurrence struct {
	DisplayNames []struct {
		Name         string
		CategoryName string
	}
	Nyckelhalsmarkt bool
	WwfApproved     bool
	DishType        struct {
		Name string
	}
	Dish MenuDish
}

var client = graphql.NewClient("https://plateimpact-heimdall.azurewebsites.net/graphql", nil)

func Test() ([]MenuDushOccurrence, error) {
	var query struct {
		DishOccurrencesByDay []MenuDushOccurrence `graphql:"dishOccurrencesByDay(mealProvidingUnitID: $resturant day: $day)"`
	}
	variables := map[string]interface{}{
		"resturant": "3d519481-1667-4cad-d2a3-08d558129279",
		"day":       "2022-11-07",
	}

	err := client.Query(context.Background(), &query, variables)
	if err != nil {
		return nil, err
	}
	return query.DishOccurrencesByDay, nil
}
