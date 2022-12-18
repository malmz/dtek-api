package lunch

import (
	"context"
	"time"

	"github.com/dtekcth/dtek-api/ent"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
)

func getCacheByDate(ctx context.Context, db *ent.Client, resturant string, date time.Time, lang string) (*ent.LunchMenu, error) {
	return db.LunchMenu.Query().
		Where(
			lunchmenu.Resturant(resturant),
			lunchmenu.Date(truncateDate(date)),
			lunchmenu.Or(
				lunchmenu.LanguageIsNil(),
				lunchmenu.LanguageEQ(lunchmenu.Language(lang)),
			),
		).
		Only(ctx)
}

func getCacheByWeek(ctx context.Context, db *ent.Client, resturant string, date time.Time, lang string) ([]*ent.LunchMenu, error) {
	return db.LunchMenu.Query().
		Where(
			lunchmenu.Resturant(resturant),
			lunchmenu.DateGTE(startOfWeek(date)),
			lunchmenu.DateLT(endOfWeek(date)),
			lunchmenu.Or(
				lunchmenu.LanguageIsNil(),
				lunchmenu.LanguageEQ(lunchmenu.Language(lang)),
			),
		).Order(ent.Asc(lunchmenu.FieldDate)).
		All(ctx)
}

func setCache(ctx context.Context, db *ent.Client, resturant string, lang string, menu *LunchFetchResult) (*ent.LunchMenu, error) {
	return db.LunchMenu.Create().
		SetResturant(resturant).
		SetDate(truncateDate(menu.Date)).
		SetLanguage(lunchmenu.Language(lang)).
		SetMenu(menu.Items).
		SetName(menu.Name).
		Save(ctx)
}

func setCacheMultiple(ctx context.Context, db *ent.Client, resturant string, lang string, menu []LunchFetchResult) ([]*ent.LunchMenu, error) {
	bulk := make([]*ent.LunchMenuCreate, len(menu))
	for i, m := range menu {
		bulk[i] = db.LunchMenu.Create().
			SetResturant(resturant).
			SetDate(truncateDate(m.Date)).
			SetLanguage(lunchmenu.Language(lang)).
			SetMenu(m.Items).
			SetName(m.Name)
	}

	return db.LunchMenu.CreateBulk(bulk...).Save(ctx)
}
