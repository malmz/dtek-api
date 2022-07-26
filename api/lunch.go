package api

import (
	"net/http"
	"time"

	"github.com/dtekcth/dtek-api/lunch"
	"github.com/labstack/echo/v4"
)

func TodaysLunch(c echo.Context) error {
	ctx := c.Request().Context()
	resturant := c.Param("resturant")
	lang := c.Param("lang")
	if lang == "" {
		lang = "se"
	}
	date := time.Now()
	menu, err := lunch.Get(ctx, resturant, date, lang)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, menu)
}
