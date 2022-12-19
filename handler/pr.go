package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type eventPost struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
}

func (e *Env) CreateEvent(c echo.Context) error {
	//ctx := c.Request().Context()

	var (
		err error
		req eventPost
	)

	if err = e.ParseQuery(c, &req); err != nil {
		return err
	}

	var (
		start_time time.Time
		end_time   time.Time
	)
	if req.StartTime == "" {
		start_time = time.Now()
	} else {
		start_time, err = time.Parse(time.RFC3339, req.StartTime)
		if err != nil {
			return err
		}
	}

	if req.EndTime == "" {
		end_time = start_time.Add(time.Hour)
	} else {
		end_time, err = time.Parse(time.RFC3339, req.EndTime)
		if err != nil {
			return err
		}
	}

	event := e.PrService.CreateEvent(req.Title, req.Description, req.Location, start_time, end_time)
	return c.JSON(http.StatusOK, event)
}
