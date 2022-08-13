package lunch

import (
	"context"
	"errors"
	"time"

	"github.com/dtekcth/dtek-api/db"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/model"
	"github.com/rs/zerolog/log"
)

type Fetcher interface {
	Fetch(date time.Time, lang string) (*model.LunchMenu, error)
}

type Resturant struct {
	fetcher Fetcher
	name    string
}

var resturants map[string]Resturant = map[string]Resturant{
	"johanneberg-express": {name: "Express Johanneberg", fetcher: NewKarenFetcher("3d519481-1667-4cad-d2a3-08d558129279")},
	"karresturangen":      {name: "Kårresturangen", fetcher: NewKarenFetcher("21f31565-5c2b-4b47-d2a1-08d558129279")},
	"hyllan":              {name: "Hyllan", fetcher: NewKarenFetcher("a7f0f75b-c1cb-4fc3-d2a6-08d558129279")},
	"smak":                {name: "Smak", fetcher: NewKarenFetcher("3ac68e11-bcee-425e-d2a8-08d558129279")},
	"linsen":              {name: "Linsen", fetcher: NewLinsenFetcher()},
}

func getFromCache(ctx context.Context, resturant string, date time.Time, lang string) (*model.LunchMenu, error) {
	db := db.Get()
	menu, err := db.LunchMenu.Query().
		Where(
			lunchmenu.Resturant(resturant),
			lunchmenu.Date(date),
			lunchmenu.Or(
				lunchmenu.LanguageIsNil(),
				lunchmenu.LanguageEQ(lunchmenu.Language(lang)),
			),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return &model.LunchMenu{
		Name:      menu.Name,
		FetchedAt: menu.UpdateTime,
		Items:     menu.Menu,
	}, nil
}

func getEmptyMenu(resturant string) *model.LunchMenu {
	if r, ok := resturants[resturant]; ok {
		return &model.LunchMenu{
			Name:      r.name,
			FetchedAt: time.Now(),
			Items:     []model.LunchMenuItem{},
		}
	}
	return nil
}

func clearOldCache(ctx context.Context) error {
	db := db.Get()
	_, err := db.LunchMenu.Delete().
		Where(lunchmenu.UpdateTimeLT(time.Now().AddDate(0, 0, -2))).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func setCache(ctx context.Context, resturant string, date time.Time, lang string, menu *model.LunchMenu) (*model.LunchMenu, error) {
	db := db.Get()
	newMenu, err := db.LunchMenu.Create().
		SetResturant(resturant).
		SetDate(date).
		SetLanguage(lunchmenu.Language(lang)).
		SetMenu(menu.Items).
		SetName(menu.Name).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return &model.LunchMenu{
		Name:      newMenu.Name,
		FetchedAt: newMenu.UpdateTime,
		Items:     newMenu.Menu,
	}, nil
}

func getFromApi(resturant string, date time.Time, lang string) (*model.LunchMenu, error) {
	if resturant, ok := resturants[resturant]; ok {
		return resturant.fetcher.Fetch(date, lang)
	}
	log.Error().Str("resturant", resturant).Msg("No fetcher for resturant")
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
