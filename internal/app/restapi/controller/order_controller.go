package controller

import (
	"carizza/internal/pkg/apperror"
	"carizza/pkg/selection_condition"

	"github.com/minipkg/go-app-common/log"
	ozzo_handler "github.com/minipkg/go-app-common/ozzo_handler"
	"github.com/minipkg/go-app-common/ozzo_handler/errorshandler"

	"carizza/internal/domain/order"

	routing "github.com/go-ozzo/ozzo-routing/v2"
)

type orderController struct {
	Logger  log.ILogger
	Service order.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/orders/ - список всех моделей
//	GET /api/order/<id> - детали модели
func RegisterOrderHandlers(r *routing.RouteGroup, service order.IService, logger log.ILogger, authHandler routing.Handler) {
	c := orderController{
		Logger:  logger,
		Service: service,
	}

	r.Get("/orders", c.list)
	r.Get(`/order/<id>`, c.get)
}

// get method is for getting a one entity by ID
func (c orderController) get(ctx *routing.Context) error {
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
func (c orderController) list(ctx *routing.Context) error {
	cond := &selection_condition.SelectionCondition{
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
	return ctx.Write(items)
}
