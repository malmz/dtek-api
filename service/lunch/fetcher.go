package lunch

import (
	"context"
	"errors"
	"time"

	"github.com/dtekcth/dtek-api/ent"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/schema"
	"github.com/dtekcth/dtek-api/model"
	"github.com/rs/zerolog/log"
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
	"johanneberg-express": {name: "Express Johanneberg", fetcher: NewKarenFetcher("3d519481-1667-4cad-d2a3-08d558129279")},
	"karresturangen":      {name: "Kårresturangen", fetcher: NewKarenFetcher("21f31565-5c2b-4b47-d2a1-08d558129279")},
	"hyllan":              {name: "Hyllan", fetcher: NewKarenFetcher("a7f0f75b-c1cb-4fc3-d2a6-08d558129279")},
	"smak":                {name: "Smak", fetcher: NewKarenFetcher("3ac68e11-bcee-425e-d2a8-08d558129279")},
	"linsen":              {name: "Linsen", fetcher: NewLinsenFetcher()},
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

func Get(ctx context.Context, resturant string, date time.Time, lang string) (*model.LunchMenu, error) {
	log.Debug().Str("resturant", resturant).Str("lang", lang).Time("date", date).Msg("Fetching lunch menu")
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)

	menu, err := getFromCache(ctx, resturant, date, lang)
	if err == nil {
		log.Debug().Str("resturant", resturant).Str("lang", lang).Time("date", date).Msg("Menu cache hit")
		return menu, nil
	}

	menu, err = getFromApi(resturant, date, lang)
	if err != nil {
		log.Debug().Err(err).Str("resturant", resturant).Str("lang", lang).Time("date", date).Msg("No menu from API")
		menu = getEmptyMenu(resturant)
	}

	log.Debug().Str("resturant", resturant).Str("lang", lang).Time("date", date).Msg("Got menu")

	menu, err = setCache(ctx, resturant, date, lang, menu)
	if err != nil {
		return nil, err
	}

	log.Debug().Str("resturant", resturant).Str("lang", lang).Time("date", date).Msg("Set menu in cache")

	return menu, nil
}

func clearOldCache(ctx context.Context, db *ent.Client) error {
	_, err := db.LunchMenu.Delete().
		Where(lunchmenu.UpdateTimeLT(time.Now().AddDate(0, 0, -2))).
		Exec(ctx)
	return err
}
