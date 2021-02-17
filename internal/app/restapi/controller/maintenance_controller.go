package controller

import (
	"net/http"

	routing "github.com/go-ozzo/ozzo-routing/v2"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"carizza/internal/pkg/log"

	"carizza/internal/domain"
	"carizza/internal/domain/maintenance"
)

type maintenanceController struct {
	Controller
	Service maintenance.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
//	GET /api/maintenances/ - список всех моделей
//	GET /api/maintenance/{ID} - детали модели
func RegisterMaintenanceHandlers(r *routing.RouteGroup, service maintenance.IService, logger log.ILogger, authHandler routing.Handler) {
	c := maintenanceController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get("/maintenances", c.list)
	r.Get(`/maintenance/<id>`, c.get)

	r.Use(authHandler)

	r.Post("/posts", c.create)
	//r.Put("/posts", c.update)
	//r.Delete(`/post/<id>`, c.delete)

}

// get method is for getting a one entity by ID
func (c maintenanceController) get(ctx *routing.Context) error {
	id, err := c.parseUintParam(ctx, "id")
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
func (c maintenanceController) list(ctx *routing.Context) error {
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
	return ctx.Write(items)
}

func (c maintenanceController) create(ctx *routing.Context) error {
	entity := c.Service.NewEntity()
	if err := ctx.Read(entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	if err := entity.Validate(); err != nil {
		return errorshandler.BadRequest(err.Error())
	}

	if err := c.Service.Create(ctx.Request.Context(), entity); err != nil {
		c.Logger.With(ctx.Request.Context()).Info(err)
		return errorshandler.BadRequest(err.Error())
	}

	return ctx.WriteWithStatus(entity, http.StatusCreated)
}
