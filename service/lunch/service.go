package lunch

import (
	"context"
	"time"

	"github.com/dtekcth/dtek-api/ent"
)

type Service struct {
	Db *ent.Client
}

func (s *Service) GetCurrent(resturant string, lang string) (*ent.LunchMenu, error) {
	return nil, nil
}

func (s *Service) GetByDate(ctx context.Context, resturant string, date time.Time, lang string) (*ent.LunchMenu, error) {
	menu, err := getCacheByDate(ctx, s.Db, resturant, date, lang)
	if err == nil {
		return menu, nil
	}

	result, err := fetchByDate(resturant, date, lang)
	if err != nil {
		return nil, err
	}
	return setCache(ctx, s.Db, resturant, lang, result)
}

func (s *Service) GetByWeek(ctx context.Context, resturant string, date time.Time, lang string) ([]*ent.LunchMenu, error) {
	menu, err := getCacheByWeek(ctx, s.Db, resturant, date, lang)
	if err == nil && len(menu) >= 5 {
		return menu, nil
	}

	result, err := fetchByWeek(resturant, date, lang)
	if err != nil {
		return nil, err
	}
	return setCacheMultiple(ctx, s.Db, resturant, lang, result)
}
