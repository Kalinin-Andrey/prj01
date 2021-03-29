package controller

import (
	"carizza/pkg/selection_condition"
	"errors"

	ozzo_handler "github.com/minipkg/go-app-common/ozzo_handler"

	"carizza/internal/pkg/apperror"

	"github.com/minipkg/go-app-common/log"
	"github.com/minipkg/go-app-common/ozzo_handler/errorshandler"

	"carizza/internal/domain/car"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type carController struct {
	Logger  log.ILogger
	Service car.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/cars/ - список всех работ
//	GET /api/car/<id> - детали модели
//	GET /client/<clientId>/cars - список работ для услуги
func RegisterCarHandlers(r *routing.RouteGroup, service car.IService, logger log.ILogger, authHandler routing.Handler) {
	c := carController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/cars", c.list)
	r.Get(`/car/<id>`, c.get)
	r.Get(`/client/<clientId>/cars`, c.list)
}

// get method is for getting a one entity by ID
func (c carController) get(ctx *routing.Context) error {
	id, err := ozzo_handler.ParseUintParam(ctx, "id")
	if err != nil {
		return errorshandler.BadRequest("ID is required to be uint")
	}

	entity, err := c.Service.Get(ctx.Request.Context(), id)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("not found")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
func (c carController) list(ctx *routing.Context) error {
	cond := &selection_condition.SelectionCondition{
		SortOrder: []map[string]string{{
			"name": "asc",
		}},
	}

	clientId, err := ozzo_handler.ParseUintParam(ctx, "clientId")
	if errors.Is(err, apperror.ErrNotFound) {
		clientId, err = ozzo_handler.ParseUintQueryParam(ctx, "clientId")
	}
	if err == nil && clientId > 0 {
		cond.Where = map[string]interface{}{
			"clientId": clientId,
		}
	}

	items, err := c.Service.Query(ctx.Request.Context(), cond)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}
	return ctx.Write(items)
}
