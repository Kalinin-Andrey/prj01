package restapi

import (
	"log"
	"net/http"
	"redditclone/internal/pkg/auth"
	"time"

	"github.com/go-ozzo/ozzo-routing/v2"
	"github.com/go-ozzo/ozzo-routing/v2/content"
	"github.com/go-ozzo/ozzo-routing/v2/cors"
	"github.com/go-ozzo/ozzo-routing/v2/file"
	"github.com/go-ozzo/ozzo-routing/v2/slash"

	"redditclone/internal/pkg/accesslog"
	"redditclone/internal/pkg/config"
	"redditclone/internal/pkg/errorshandler"

	commonApp "redditclone/internal/app"
	"redditclone/internal/app/restapi/controller"
)

// Version of API
const Version = "1.0.0"

// App is the application for API
type App struct {
	*commonApp.App
	Server *http.Server
}

// New func is a constructor for the ApiApp
func New(commonApp *commonApp.App, cfg config.Configuration) *App {
	app := &App{
		App:    commonApp,
		Server: nil,
	}

	// build HTTP server
	server := &http.Server{
		Addr:    cfg.Server.HTTPListen,
		Handler: app.buildHandler(),
	}
	app.Server = server

	return app
}

func (app *App) buildHandler() *routing.Router {
	router := routing.New()

	router.Use(
		accesslog.Handler(app.Logger),
		slash.Remover(http.StatusMovedPermanently),
		errorshandler.Handler(app.Logger),
		content.TypeNegotiator(content.JSON),
		cors.Handler(cors.AllowAll),
	)
	//router.NotFound(file.Content("website/index.html"))

	// serve index file
	router.Get("/", file.Content("website/index.html"))
	router.Get("/a/*", file.Content("website/index.html"))
	router.Get("/u/*", file.Content("website/index.html"))

	router.Get("/favicon.ico", file.Content("website/favicon.ico"))
	router.Get("/manifest.json", file.Content("website/manifest.json"))
	// serve files under the "static" subdirectory
	router.Get("/static/*", file.Server(file.PathMap{
		"/static/": "/website/static/",
	}))

	rg := router.Group("/api")

	authMiddleware := auth.Middleware(app.Logger, app.Auth.Service)

	auth.RegisterHandlers(rg.Group(""),
		app.Auth.Service,
		app.Logger,
	)

	app.RegisterHandlers(rg, authMiddleware)

	return router
}

// Run is func to run the ApiApp
func (app *App) Run() error {
	go func() {
		defer func() {
			if err := app.DB.DB().Close(); err != nil {
				app.Logger.Error(err)
			}

			err := app.Logger.Sync()
			if err != nil {
				log.Println(err.Error())
			}
		}()
		// start the HTTP server with graceful shutdown
		routing.GracefulShutdown(app.Server, 10*time.Second, app.Logger.Infof)
	}()
	app.Logger.Infof("server %v is running at %v", Version, app.Server.Addr)
	if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// RegisterHandlers sets up the routing of the HTTP handlers.
func (app *App) RegisterHandlers(rg *routing.RouteGroup, authMiddleware routing.Handler) {

	//controller.RegisterUserHandlers(rg, app.Domain.User.Service, app.Logger, authMiddleware)
	controller.RegisterPostHandlers(rg, app.Domain.Post.Service, app.Domain.User.Service, app.Logger, authMiddleware)
	controller.RegisterCommentHandlers(rg, app.Domain.Comment.Service, app.Domain.Post.Service, app.Logger, authMiddleware)
	controller.RegisterVoteHandlers(rg, app.Domain.Vote.Service, app.Domain.Post.Service, app.Logger, authMiddleware)

}
