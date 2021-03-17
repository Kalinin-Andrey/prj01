package controller

import (
	"errors"

	"carizza/internal/pkg/apperror"

	"github.com/minipkg/go-app-common/log"
	ozzo_handler "github.com/minipkg/go-app-common/ozzo_handler"
	"github.com/minipkg/go-app-common/ozzo_handler/errorshandler"

	"carizza/pkg/selection_condition"

	"carizza/internal/domain/modification"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type modificationController struct {
	Logger  log.ILogger
	Service modification.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/modifications/ - список всех моделей
//	GET /api/modification/{ID} - детали модели
func RegisterModificationHandlers(r *routing.RouteGroup, service modification.IService, logger log.ILogger, authHandler routing.Handler) {
	c := modificationController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/modifications", c.list)
	r.Get(`/modification/<id>`, c.get)
	r.Get(`/serie/<id>/modifications`, c.list)
}

// get method is for getting a one entity by ID
func (c modificationController) get(ctx *routing.Context) error {
	id, err := ozzo_handler.ParseUintParam(ctx, "id")
	if err != nil {
		errorshandler.BadRequest("ID is required to be uint")
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
func (c modificationController) list(ctx *routing.Context) error {
	cond := selection_condition.SelectionCondition{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	serieId, err := ozzo_handler.ParseUintParam(ctx, "serieId")
	if errors.Is(err, apperror.ErrNotFound) {
		serieId, err = ozzo_handler.ParseUintQueryParam(ctx, "serieId")
	}
	if err == nil && serieId > 0 {
		cond.Where = map[string]interface{}{
			"id_car_serie": serieId,
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
