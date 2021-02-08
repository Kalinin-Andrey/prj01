package controller

import (
	"errors"

	"github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain"
	"carizza/internal/domain/car"
)

type carController struct {
	Controller
	Service car.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/cars/ - список всех работ
//	GET /api/car/<id> - детали модели
//	GET /client/<clientId>/cars - список работ для услуги
func RegisterCarHandlers(r *routing.RouteGroup, service car.IService, logger log.ILogger, authHandler routing.Handler) {
	c := carController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/cars", c.list)
	r.Get(`/car/<id>`, c.get)
	r.Get(`/client/<clientId>/cars`, c.list)
}

// get method is for getting a one entity by ID
func (c carController) get(ctx *routing.Context) error {
	id, err := c.parseUintParam(ctx, "id")
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

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
func (c carController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	clientId, err := c.parseUintParam(ctx, "clientId")
	if errors.Is(err, apperror.ErrNotFound) {
		clientId, err = c.parseUintQueryParam(ctx, "clientId")
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
	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.Write(items)
}
