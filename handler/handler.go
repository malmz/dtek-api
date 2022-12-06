package handler

import (
	"github.com/dtekcth/dtek-api/ent"
	"github.com/dtekcth/dtek-api/service/lunch"
)

type Env struct {
	Db           *ent.Client
	LunchService *lunch.Service
}
