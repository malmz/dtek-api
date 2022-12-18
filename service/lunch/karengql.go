package lunch

import (
	"context"
	"fmt"
	"time"

	"github.com/dtekcth/dtek-api/ent/schema"
	"github.com/hasura/go-graphql-client"
	"github.com/rs/zerolog/log"
)

type MenuDish struct {
	Recipes []struct {
		Allergens []struct {
			Code     string
			ImageUrl string
		}
	}
}

type MenuDushOccurrence struct {
	MenuBody []struct {
		Text     string `graphql:"name"`
		Language string `graphql:"categoryName"`
	} `graphql:displayNames`
	Date            time.Time `graphql:"startDate"`
	Nyckelhalsmarkt bool
	WwfApproved     bool
	Title           struct {
		Text string
	} `graphql:"dishType"`
	Dish      MenuDish
	Resturant struct {
		Name string `graphql:"mealProvidingUnitName"`
	} `graphql:"mealProvidingUnit"`
}

var client = graphql.NewClient("https://plateimpact-heimdall.azurewebsites.net/graphql", nil)

type KarenGQLFetcher struct {
	resturant string
}

func (f *KarenGQLFetcher) FetchByDate(ctx context.Context, date time.Time, lang string) (*LunchFetchResult, error) {
	var query struct {
		dishOccurrences []MenuDushOccurrence `graphql:"dishOccurrencesByDay(
			mealProvidingUnitID: $resturant 
			day: $day
		)"`
	}

	if err := client.Query(ctx, &query, map[string]interface{}{
		"resturant": graphql.String(f.resturant),
		"day":       graphql.String(date.Format("2006-01-02")),
	}); err != nil {
		return nil, err
	}

	menus, err := parseQuery(query.dishOccurrences, lang)
	if err != nil {
		return nil, err
	}

	if len(menus) != 1 {
		log.Warn().Msgf("Expected 1 menu, got %d", len(menus))
	}

	return &menus[0], nil
}

func (f *KarenGQLFetcher) FetchByWeek(ctx context.Context, date time.Time, lang string) ([]LunchFetchResult, error) {
	var query struct {
		dishOccurrences []MenuDushOccurrence `graphql:"dishOccurrencesByTimeRange(
			mealProvidingUnitID: $resturant 
			startDate: $startDate
			endDate: $endDate
		)"`
	}

	start, end := weekBounds(date)

	if err := client.Query(ctx, &query, map[string]interface{}{
		"resturant": graphql.String(f.resturant),
		"startDate": graphql.String(start.Format("2006-01-02")),
		"endDate":   graphql.String(end.Format("2006-01-02")),
	}); err != nil {
		return nil, err
	}

	return parseQuery(query.dishOccurrences, lang)
}

func parseQuery(query []MenuDushOccurrence, lang string) ([]LunchFetchResult, error) {
	days := map[time.Time]LunchFetchResult{}
	if len(query) == 0 {
		return nil, fmt.Errorf("No menu found")
	}
	for _, dish := range query {

		res, ok := days[dish.Date]
		if !ok {
			res = LunchFetchResult{
				Date:  dish.Date,
				Name:  dish.Resturant.Name,
				Items: []schema.LunchMenuItem{},
			}
		}

		for _, b := range dish.MenuBody {
			if b.Language != lang {
				continue
			}
			var allergens []schema.Allergen
			if len(dish.Dish.Recipes) > 0 {
				for _, allergen := range dish.Dish.Recipes[0].Allergens {
					allergens = append(allergens, schema.Allergen{
						Code:     allergen.Code,
						ImageUrl: allergen.ImageUrl,
					})
				}
			}
			res.Items = append(res.Items, schema.LunchMenuItem{
				Title:        dish.Title.Text,
				Body:         b.Text,
				Preformatted: false,
				Allergens:    allergens,
			})
		}
		days[dish.Date] = res
	}

	var returnDays []LunchFetchResult
	for _, day := range days {
		returnDays = append(returnDays, day)
	}

	return returnDays, nil
}
