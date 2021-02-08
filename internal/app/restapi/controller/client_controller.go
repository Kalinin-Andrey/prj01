package controller

import (
	"github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain"
	"carizza/internal/domain/client"
)

type clientController struct {
	Controller
	Service client.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/clients/ - список всех работ
//	GET /api/client/<id> - детали модели
func RegisterClientHandlers(r *routing.RouteGroup, service client.IService, logger log.ILogger, authHandler routing.Handler) {
	c := clientController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/clients", c.list)
	r.Get(`/client/<id>`, c.get)
}

// get method is for getting a one entity by ID
func (c clientController) get(ctx *routing.Context) error {
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
func (c clientController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		SortOrder: map[string]string{
			"name": "asc",
		},
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
