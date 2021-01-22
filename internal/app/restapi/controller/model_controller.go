package controller

import (
	"carizza/internal/domain"
	"carizza/internal/domain/ctype"

	"github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain/model"
)

type postController struct {
	Controller
	Service model.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/models/ - список всех моделей
//	GET /api/model/{MODEL_ID} - детали поста с комментами
func RegisterPostHandlers(r *routing.RouteGroup, service model.IService, logger log.ILogger, authHandler routing.Handler) {
	c := postController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/models", c.list)
	r.Get(`/model/<id>`, c.get)
}

// get method is for getting a one entity by ID
func (c postController) get(ctx *routing.Context) error {
	id := ctx.Param("id")

	entity, err := c.Service.Get(ctx.Request.Context(), id)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}

	ctx.Response.Header().Set("Content-Type", "application/json; charset=UTF-8")
	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
func (c postController) list(ctx *routing.Context) error {
	cond := domain.DBQueryConditions{
		Where: map[string]interface{}{
			"id_car_type": ctype.TypeIDCar,
		},
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
