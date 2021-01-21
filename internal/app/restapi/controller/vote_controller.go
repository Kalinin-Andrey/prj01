package controller

import (
	"github.com/go-ozzo/ozzo-routing/v2"
	"carizza/internal/domain/post"
	"carizza/internal/domain/vote"
	"carizza/internal/pkg/log"
)

type voteController struct {
	Controller
	Service     vote.IService
	PostService post.IService
	Logger      log.ILogger
}

func RegisterVoteHandlers(r *routing.RouteGroup, service vote.IService, postService post.IService, logger log.ILogger, authHandler routing.Handler) {
	/*c := voteController{
		Service:     service,
		PostService: postService,
		Logger:      logger,
	}

	r.Use(authHandler)
	*/
}
