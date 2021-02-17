package controller

import (
	"errors"

	"carizza/internal/domain"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain/model"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type modelController struct {
	Controller
	Service model.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/models/ - список всех моделей
//	GET /api/model/{ID} - детали модели
func RegisterModelHandlers(r *routing.RouteGroup, service model.IService, logger log.ILogger, authHandler routing.Handler) {
	c := modelController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/models", c.list)
	r.Get(`/model/<id>`, c.get)
	r.Get(`/mark/<markId>/models`, c.list)
}

// get method is for getting a one entity by ID
func (c modelController) get(ctx *routing.Context) error {
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

	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
func (c modelController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	markId, err := c.parseUintParam(ctx, "markId")
	if errors.Is(err, apperror.ErrNotFound) {
		markId, err = c.parseUintQueryParam(ctx, "markId")
	}
	if err == nil && markId > 0 {
		cond.Where = map[string]interface{}{
			"id_car_mark": markId,
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
