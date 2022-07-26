package lunch

import (
	"context"
	"time"

	"github.com/dtekcth/dtek-api/db"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/model"
	"github.com/rs/zerolog/log"
)

type Fetcher interface {
	Fetch(date time.Time, lang string) (*model.LunchMenu, error)
}

var fetcher map[string]Fetcher = map[string]Fetcher{
	"johanneberg-express": NewKarenFetcher("3d519481-1667-4cad-d2a3-08d558129279"),
	"karresturangen":      NewKarenFetcher("21f31565-5c2b-4b47-d2a1-08d558129279"),
	"hyllan":              NewKarenFetcher("a7f0f75b-c1cb-4fc3-d2a6-08d558129279"),
	"smak":                NewKarenFetcher("3ac68e11-bcee-425e-d2a8-08d558129279"),
	"linsen":              NewLinsenFetcher(),
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
		Items: menu.Menu,
	}, nil
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

func setCache(ctx context.Context, resturant string, date time.Time, lang string, menu *model.LunchMenu) error {
	db := db.Get()
	_, err := db.LunchMenu.Create().
		SetResturant(resturant).
		SetDate(date).
		SetLanguage(lunchmenu.Language(lang)).
		SetMenu(menu.Items).
		Save(ctx)
	if err != nil {
		return err
	}

	return nil
}

func getFromApi(resturant string, date time.Time, lang string) (*model.LunchMenu, error) {
	return fetcher[resturant].Fetch(date, lang)
}

func GetFresh(ctx context.Context, resturant string, date time.Time, lang string) (*model.LunchMenu, error) {
	return getFromApi(resturant, date, lang)
}

func Get(ctx context.Context, resturant string, date time.Time, lang string) (*model.LunchMenu, error) {
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.Local)
	menu, err := getFromCache(ctx, resturant, date, lang)
	if err == nil {
		log.Debug().Str("resturant", resturant).Str("lang", lang).Time("date", date).Msg("Menu cache hit")
		return menu, nil
	}

	menu, err = getFromApi(resturant, date, lang)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("resturant", resturant).Str("lang", lang).Time("date", date).Msg("Got menu from API")

	err = setCache(ctx, resturant, date, lang, menu)
	if err != nil {
		return nil, err
	}
	log.Debug().Str("resturant", resturant).Str("lang", lang).Time("date", date).Msg("Set menu in cache")

	return menu, nil
}
