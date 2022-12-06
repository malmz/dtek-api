package lunch

import (
	"context"
	"time"

	"github.com/dtekcth/dtek-api/ent/schema"
	"github.com/hasura/go-graphql-client"
	"github.com/rs/zerolog/log"
)

type MenuDish struct {
	Recipes []struct {
		Alergens []struct {
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
	Resturant       struct {
		Name string
	} `graphql:"dishType"`
	Dish MenuDish
}

var client = graphql.NewClient("https://plateimpact-heimdall.azurewebsites.net/graphql", nil)

type KarenGQLFetcher struct {
	resturant string
}

func (f *KarenGQLFetcher) FetchByDate(ctx context.Context, date time.Time, lang string) (*LunchFetchResult, error) {
	var query struct {
		dishOccurrencesByTimeRange []MenuDushOccurrence `graphql:"dishOccurrencesByDay(
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

	menuDate := query.dishOccurrencesByTimeRange[0].Date
	resturantName := query.dishOccurrencesByTimeRange[0].Resturant.Name

	for _, dish := range query.dishOccurrencesByTimeRange {
		if !dish.Date.Equal(menuDate) {
			log.Warn().Str("resturant", f.resturant).
				Time("date", menuDate).
				Msg("multiple dates in menu when not expected, skipping")
			continue
		}
		if !(dish.Resturant.Name == resturantName) {
			log.Warn().Str("resturant", f.resturant).
				Time("date", menuDate).
				Msg("multiple resturants in menu when not expected, skipping")
			continue
		}
	}

	return &LunchFetchResult{
		Date: dishDate,
	}, nil
}

func parseQuery(query []MenuDushOccurrence) ([]LunchFetchResult, error) {
	days := map[time.Time]*LunchFetchResult{}
	for _, dish := range query {

		res, ok := days[dish.Date]
		if !ok {
			res = &LunchFetchResult{
				Date: dish.Date,
				Items: []schema.LunchMenuItem{},
			}
		}

		menu := schema.LunchMenuItem{
			dish.
		} 
		dish.

		if dish.Resturant.Name != days[dishDate].Resturant {
			log.Warn().Str("resturant", dish.Resturant.Name).
				Time("date", dishDate).
				Msg("multiple resturants in menu when not expected, skipping")
			continue
		}
		if dish.Nyckelhalsmarkt {
			days[dishDate].Nyckelhalsmarkt = true
		}
		if dish.WwfApproved {
			days[dishDate].WwfApproved = true
		}
	}
}
