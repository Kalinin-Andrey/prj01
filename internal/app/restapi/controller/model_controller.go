package controller

import (
	"carizza/internal/domain"

	ozzo_handler "github.com/minipkg/go-app-common/ozzo_handler"

	"carizza/internal/pkg/apperror"

	"github.com/minipkg/go-app-common/log"
	"github.com/minipkg/go-app-common/ozzo_handler/errorshandler"

	"carizza/internal/domain/model"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type modelController struct {
	Logger  log.ILogger
	Service model.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/models/ - список всех моделей
//	GET /api/model/{ID} - детали модели
func RegisterModelHandlers(r *routing.RouteGroup, service model.IService, logger log.ILogger, authHandler routing.Handler) {
	c := modelController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/models", c.list)
	r.Get(`/model/<id>`, c.get)
	r.Get(`/mark/<markId>/models`, c.list)
}

// get method is for getting a one entity by ID
func (c modelController) get(ctx *routing.Context) error {
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
func (c modelController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
	}

	if len(ctx.Request.URL.Query()) > 0 {
		where := c.Service.NewEntity()
		err := ozzo_handler.ParseQueryParams(ctx, where)
		if err != nil {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.BadRequest("")
		}
		cond.Where = where
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
