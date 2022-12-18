package handler

import (
	"github.com/dtekcth/dtek-api/ent"
	"github.com/dtekcth/dtek-api/service/lunch"
	"github.com/labstack/echo/v4"
)

type Env struct {
	Db           *ent.Client
	LunchService *lunch.Service
}

func (e *Env) ParseQuery(c echo.Context, req interface{}) error {
	if err := c.Bind(&req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	return nil
}
