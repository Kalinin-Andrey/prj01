package controller

import (
	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/errorshandler"
	"strconv"

	routing "github.com/go-ozzo/ozzo-routing/v2"
	"github.com/pkg/errors"

	"carizza/internal/domain/user"
	"carizza/internal/pkg/log"
)

type userController struct {
	Controller
	Service user.IService
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func RegisterUserHandlers(r *routing.RouteGroup, service user.IService, logger log.ILogger, authHandler routing.Handler) {
	c := userController{
		Controller: Controller{
			Logger: logger,
		},
		Service: service,
	}

	r.Get(`/user/<id:\d+>`, c.get)
	//r.Get("/users", c.list)

}

// get method is for a getting a one enmtity by ID
func (c userController) get(ctx *routing.Context) error {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		c.Logger.With(ctx.Request.Context()).Info(errors.Wrapf(err, "Can not parse uint64 from %q", ctx.Param("id")))
		return errorshandler.BadRequest("id mast be a uint")
	}
	entity, err := c.Service.Get(ctx.Request.Context(), uint(id))
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}
	return ctx.Write(entity)
}

// list method is for a getting a list of all entities
/*func (c userController) list(ctx *routing.Context) error {
	rctx := ctx.Request.Context()
	items, err := c.Service.List(rctx)
	if err != nil {
		if err == apperror.ErrNotFound {
			c.Logger.With(ctx.Request.Context()).Info(err)
			return errorshandler.NotFound("")
		}
		c.Logger.With(ctx.Request.Context()).Error(err)
		return errorshandler.InternalServerError("")
	}
	return ctx.Write(items)
}*/
