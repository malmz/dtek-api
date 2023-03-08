package lunch

import (
	"context"
	"time"

	"github.com/dtekcth/dtek-api/db"
	"github.com/dtekcth/dtek-api/ent"
	"github.com/jackc/pgtype"
)

type Service struct {
	Db *db.Queries
}

func (s *Service) GetCurrent(resturant string, lang string) (*ent.LunchMenu, error) {
	return nil, nil
}

func (s *Service) GetByDate(ctx context.Context, resturant string, date time.Time, lang string) (*db.LunchMenu, error) {
	menu, err := s.Db.GetLunchByDate(ctx, db.GetLunchByDateParams{
		resturant,
		date,
		db.Language(lang),
	})

	if err == nil {
		return &menu, nil
	}

	result, err := fetchByDate(ctx, resturant, date, lang)
	if err != nil {
		return nil, err
	}

	var items pgtype.JSONB
	items.Set(result.Items)

	s.Db.CreateLunchMenus(ctx, []db.CreateLunchMenusParams{
		db.CreateLunchMenusParams{
			Resturant: resturant,
			Date:      date,
			Language:  db.Language(lang),
			Name:      result.Name,
			Menu:      items,
		},
	})

	menu, err = s.Db.GetLunchByDate(ctx, db.GetLunchByDateParams{
		resturant,
		date,
		db.Language(lang),
	})

	return &menu, err
}

func (s *Service) GetByWeek(ctx context.Context, resturant string, date time.Time, lang string) ([]*ent.LunchMenu, error) {
	start := startOfWeek(date)
	end := endOfWeek(date)

	menu, err := s.Db.GetLunchByDateRange(ctx, db.GetLunchByDateRangeParams{
		resturant,
		start,
		end,
		db.Language(lang),
	})

	if err == nil && len(menu) >= 5 {
		return menu, nil
	}

	results, err := fetchByWeek(ctx, resturant, date, lang)
	if err != nil {
		return nil, err
	}

	var create []db.CreateLunchMenusParams

	for _, result := range results {
		var items pgtype.JSONB
		items.Set(result.Items)

		create = append(create, db.CreateLunchMenusParams{
			Resturant: resturant,
			Date:      result.Date,
			Language:  db.Language(lang),
			Name:      result.Name,
			Menu:      items,
		})
	}

	s.Db.CreateLunchMenus(ctx, create)

	menu, err = s.Db.GetLunchByDateRange(ctx, db.GetLunchByDateRangeParams{
		resturant,
		start,
		end,
		db.Language(lang),
	})

	return menu, err
}
