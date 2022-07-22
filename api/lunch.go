package api

import (
	"net/http"
	"time"

	"github.com/dtekcth/dtek-api/db"
	"github.com/dtekcth/dtek-api/ent"
	"github.com/dtekcth/dtek-api/ent/lunchmenu"
	"github.com/dtekcth/dtek-api/ent/lunchmenuitem"
	"github.com/dtekcth/dtek-api/ent/resturant"
	"github.com/dtekcth/dtek-api/model"
	"github.com/labstack/echo/v4"
)

func TodaysLunch(c echo.Context) error {
	ctx := c.Request().Context()
	db := db.Get()
	menu, err := db.Resturant.Query().
		Where(resturant.Slug("express"), resturant.CampusEQ(resturant.CampusJohanneberg)).
		WithMenu(func(lmq *ent.LunchMenuQuery) {
			lmq.Where(lunchmenu.Date(time.Now())).
				WithItems(func(lmiq *ent.LunchMenuItemQuery) {
					lmiq.Where(lunchmenuitem.LanguageEQ(lunchmenuitem.LanguageSe))
				})
		}).Only(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	lunchMenu := menu.Edges.Menu[0]
	items := lunchMenu.Edges.Items

	mdlItems := make([]model.LunchMenuItem, len(items))
	for i, item := range items {
		mdlItems[i] = model.LunchMenuItem{
			Title:        item.Title,
			Body:         item.Body,
			Preformatted: item.Preformatted,
			Allergens:    item.Allergens,
			Emission:     item.Emission,
			Price:        item.Price,
		}
	}

	mdlMenu := model.LunchMenu{
		Resturant: menu.Name,
		Campus:    menu.Campus.String(),
		Date:      lunchMenu.Date.Format(time.RFC3339),
		Language:  lunchmenuitem.LanguageSe.String(),
		Items:     mdlItems,
	}

	return c.JSON(http.StatusOK, mdlMenu)
}
