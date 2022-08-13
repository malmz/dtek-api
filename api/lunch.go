package api

import (
	"net/http"
	"time"

	"github.com/dtekcth/dtek-api/lunch"
	"github.com/dtekcth/dtek-api/model"
	"github.com/labstack/echo/v4"
)

type lunchRequest struct {
	Resturant []string `query:"resturant"`
	Lang      string   `query:"lang" validate:"required,oneof=se en"`
	Date      string   `query:"date" validate:"datetime"`
}

func GetLunch(c echo.Context) error {
	ctx := c.Request().Context()
	req := lunchRequest{}
	var err error
	if err = c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err = c.Validate(req); err != nil {
		return err
	}

	var date time.Time

	if req.Date == "" {
		date = time.Now()
	} else {
		date, err = time.Parse(time.RFC3339, req.Date)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}
	menuList := make([]*model.LunchMenu, 0, len(req.Resturant))
	for _, r := range req.Resturant {
		menu, err := lunch.Get(ctx, r, date, req.Lang)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}
		menuList = append(menuList, menu)
	}

	return c.JSON(http.StatusOK, menuList)
}
