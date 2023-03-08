package handler

import (
	"net/http"
	"time"

	"github.com/dtekcth/dtek-api/model"
	"github.com/labstack/echo/v4"
)

type lunchRequest struct {
	Resturant []string `query:"resturant"`
	Lang      string   `query:"lang" validate:"required,oneof=se en"`
	Date      string   `query:"date" validate:"datetime"`
}

func (e *Env) GetLunch(c echo.Context) error {
	ctx := c.Request().Context()

	var (
		err  error
		req  lunchRequest
		date time.Time
	)

	if err = e.ParseQuery(c, &req); err != nil {
		return err
	}

	if req.Date == "" {
		date = time.Now()
	} else {
		date, err = time.Parse(time.RFC3339, req.Date)
		if err != nil {
			return err
		}
	}

	menuList := make([]*model.LunchMenu, 0, len(req.Resturant))
	for _, r := range req.Resturant {
		menu, err := e.LunchService.GetByDate(ctx, r, date, req.Lang)
		if err != nil {
			return err
		}

		items := make([]model.LunchMenuItem, 0, len(menu.Menu))
		for _, item := range menu.Menu {
			allergens := make([]model.Allergen, 0, len(item.Allergens))
			for _, a := range item.Allergens {
				allergens = append(allergens, model.Allergen{
					Code: a.Code, ImageUrl: a.ImageUrl,
				})
			}

			items = append(items, model.LunchMenuItem{
				Title:        item.Title,
				Body:         item.Body,
				Preformatted: item.Preformatted,
				Allergens:    allergens,
				Emission:     item.Emission,
				Price:        item.Price,
			})
		}

		mMenu := &model.LunchMenu{
			Resturant: menu.Resturant,
			Name:      menu.Name,
			FetchedAt: menu.UpdateTime,
			Items:     items,
		}

		menuList = append(menuList, mMenu)
	}

	return c.JSON(http.StatusOK, menuList)
}
