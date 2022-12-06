package handler

import (
	"net/http"
	"strconv"

	"github.com/dtekcth/dtek-api/model"
	"github.com/labstack/echo/v4"
)

func (e *Env) CreateNews(c echo.Context) error {
	ctx := c.Request().Context()
	news := &model.News{}
	if err := c.Bind(news); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	n, err := e.Db.News.Create().SetTitle(news.Title).SetContent(news.Content).Save(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	mdlNews := model.News{
		ID:      n.ID,
		Title:   n.Title,
		Content: n.Content,
		Updated: n.UpdateTime,
		Created: n.CreateTime,
	}
	return c.JSON(http.StatusOK, mdlNews)
}

func (e *Env) ListNews(c echo.Context) error {
	ctx := c.Request().Context()
	news, err := e.Db.News.Query().All(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	mdlNews := make([]model.News, len(news))
	for i, n := range news {
		mdlNews[i] = model.News{
			ID:      n.ID,
			Title:   n.Title,
			Content: n.Content,
			Updated: n.UpdateTime,
			Created: n.CreateTime,
		}
	}
	return c.JSON(http.StatusOK, mdlNews)
}

func (e *Env) GetNews(c echo.Context) error {
	ctx := c.Request().Context()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	n, err := e.Db.News.Get(ctx, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	mdlNews := model.News{
		ID:      n.ID,
		Title:   n.Title,
		Content: n.Content,
		Updated: n.UpdateTime,
		Created: n.CreateTime,
	}
	return c.JSON(http.StatusOK, mdlNews)
}

func (e *Env) UpdateNews(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	news := &model.News{}
	if err := c.Bind(news); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	n, err := e.Db.News.UpdateOneID(id).SetTitle(news.Title).SetContent(news.Content).Save(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	mdlNews := model.News{
		ID:      n.ID,
		Title:   n.Title,
		Content: n.Content,
		Updated: n.UpdateTime,
		Created: n.CreateTime,
	}
	return c.JSON(http.StatusOK, mdlNews)
}
