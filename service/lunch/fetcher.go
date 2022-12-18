package lunch

import (
	"context"
	"errors"
	"time"

	"github.com/dtekcth/dtek-api/ent"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/schema"
)

type LunchFetchResult struct {
	Name  string
	Date  time.Time
	Items []schema.LunchMenuItem
}

type Fetcher interface {
	FetchByDate(ctx context.Context, date time.Time, lang string) (*LunchFetchResult, error)
	FetchByWeek(ctx context.Context, date time.Time, lang string) ([]LunchFetchResult, error)
}

type resturant struct {
	fetcher Fetcher
	name    string
}

var resturants map[string]resturant = map[string]resturant{
	"johanneberg-express": {name: "Express Johanneberg", fetcher: &KarenGQLFetcher{"3d519481-1667-4cad-d2a3-08d558129279"}},
	"karresturangen":      {name: "Kårresturangen", fetcher: &KarenGQLFetcher{"21f31565-5c2b-4b47-d2a1-08d558129279"}},
	"hyllan":              {name: "Hyllan", fetcher: &KarenGQLFetcher{"a7f0f75b-c1cb-4fc3-d2a6-08d558129279"}},
	"smak":                {name: "Smak", fetcher: &KarenGQLFetcher{"3ac68e11-bcee-425e-d2a8-08d558129279"}},
	"linsen":              {name: "Linsen", fetcher: &KarenGQLFetcher{"b672efaf-032a-4bb8-d2a5-08d558129279"}},
}

func fetchByDate(ctx context.Context, resturant string, date time.Time, lang string) (*LunchFetchResult, error) {
	if resturant, ok := resturants[resturant]; ok {
		return resturant.fetcher.FetchByDate(ctx, date, lang)
	}
	return nil, errors.New("no fetcher for resturant")
}

func fetchByWeek(ctx context.Context, resturant string, date time.Time, lang string) ([]LunchFetchResult, error) {
	if resturant, ok := resturants[resturant]; ok {
		return resturant.fetcher.FetchByWeek(ctx, date, lang)
	}
	return nil, errors.New("no fetcher for resturant")
}

func clearOldCache(ctx context.Context, db *ent.Client) error {
	_, err := db.LunchMenu.Delete().
		Where(lunchmenu.UpdateTimeLT(time.Now().AddDate(0, 0, -2))).
		Exec(ctx)
	return err
}
