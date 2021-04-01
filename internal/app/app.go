package app

import (
	"carizza/internal/domain/address"
	"carizza/internal/domain/car"
	"carizza/internal/domain/client"
	"carizza/internal/domain/generation"
	"carizza/internal/domain/maintenance"
	"carizza/internal/domain/mark"
	"carizza/internal/domain/modification"
	"carizza/internal/domain/order"
	"carizza/internal/domain/serie"
	"carizza/internal/domain/supply"
	"carizza/internal/domain/work"
	golog "log"

	"github.com/pkg/errors"

	"carizza/internal/pkg/apperror"
	"carizza/internal/pkg/auth"
	"carizza/internal/pkg/config"
	"carizza/internal/pkg/jwt"

	pg "github.com/minipkg/db/gorm"
	"github.com/minipkg/db/redis"
	"github.com/minipkg/db/redis/cache"
	"github.com/minipkg/log"

	"carizza/internal/domain/ctype"
	"carizza/internal/domain/model"
	"carizza/internal/domain/user"
	pgrep "carizza/internal/infrastructure/repository/pg"
	redisrep "carizza/internal/infrastructure/repository/redis"
)

// App struct is the common part of all applications
type App struct {
	Cfg           config.Configuration
	Logger        log.ILogger
	IdentityDB    pg.IDB
	CarCatalogDB  pg.IDB
	MaintenanceDB pg.IDB
	Redis         redis.IDB
	Domain        Domain
	Auth          Auth
	Cache         cache.Service
}

type Auth struct {
	SessionRepository auth.SessionRepository
	TokenRepository   auth.TokenRepository
	Service           auth.Service
}

// Domain is a Domain Layer Entry Point
type Domain struct {
	User DomainUser
	//	CarCatalog
	Type         DomainType
	Mark         DomainMark
	Model        DomainModel
	Generation   DomainGeneration
	Serie        DomainSerie
	Modification DomainModification
	//	MaintenanceCatalog
	Maintenance DomainMaintenance
	Work        DomainWork
	Supply      DomainSupply
	//	Order
	Order DomainOrder
	//	Client
	Client  DomainClient
	Address DomainAddress
	Car     DomainCar
}

type DomainUser struct {
	Repository user.Repository
	Service    user.IService
}

type DomainType struct {
	Repository ctype.Repository
	Service    ctype.IService
}

type DomainMark struct {
	Repository mark.Repository
	Service    mark.IService
}

type DomainModel struct {
	Repository model.Repository
	Service    model.IService
}

type DomainGeneration struct {
	Repository generation.Repository
	Service    generation.IService
}

type DomainSerie struct {
	Repository serie.Repository
	Service    serie.IService
}

type DomainModification struct {
	Repository modification.Repository
	Service    modification.IService
}

type DomainMaintenance struct {
	Repository maintenance.Repository
	Service    maintenance.IService
}

type DomainWork struct {
	Repository work.Repository
	Service    work.IService
}

type DomainSupply struct {
	Repository supply.Repository
	Service    supply.IService
}

type DomainOrder struct {
	Repository order.Repository
	Service    order.IService
}

type DomainClient struct {
	Repository client.Repository
	Service    client.IService
}

type DomainAddress struct {
	Repository address.Repository
	Service    address.IService
}

type DomainCar struct {
	Repository car.Repository
	Service    car.IService
}

// New func is a constructor for the App
func New(cfg config.Configuration) *App {
	logger, err := log.New(cfg.Log)
	if err != nil {
		golog.Fatal(err)
	}

	IdentityDB, err := pg.New(cfg.DB.Identity, logger)
	if err != nil {
		golog.Fatal(err)
	}

	CarCatalogDB, err := pg.New(cfg.DB.CarCatalog, logger)
	if err != nil {
		golog.Fatal(err)
	}

	MaintenanceDB, err := pg.New(cfg.DB.Maintenance, logger)
	if err != nil {
		golog.Fatal(err)
	}

	rDB, err := redis.New(cfg.DB.Redis)
	if err != nil {
		golog.Fatal(err)
	}

	app := &App{
		Cfg:           cfg,
		Logger:        logger,
		IdentityDB:    IdentityDB,
		CarCatalogDB:  CarCatalogDB,
		MaintenanceDB: MaintenanceDB,
		Redis:         rDB,
	}

	err = app.Init()
	if err != nil {
		golog.Fatal(err)
	}

	return app
}

func (app *App) Init() (err error) {
	if err := app.SetupRepositories(); err != nil {
		return err
	}
	app.SetupServices()
	return nil
}

func (app *App) SetupRepositories() (err error) {
	var ok bool

	app.Domain.User.Repository, ok = app.getPgRepo(app.IdentityDB, user.EntityName).(user.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", user.EntityName, user.EntityName, app.getPgRepo(app.IdentityDB, user.EntityName))
	}
	//	CarCatalog
	app.Domain.Mark.Repository, ok = app.getPgRepo(app.CarCatalogDB, mark.EntityName).(mark.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", mark.EntityName, mark.EntityName, app.getPgRepo(app.CarCatalogDB, mark.EntityName))
	}

	app.Domain.Model.Repository, ok = app.getPgRepo(app.CarCatalogDB, model.EntityName).(model.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", model.EntityName, model.EntityName, app.getPgRepo(app.CarCatalogDB, model.EntityName))
	}

	app.Domain.Generation.Repository, ok = app.getPgRepo(app.CarCatalogDB, generation.EntityName).(generation.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", generation.EntityName, generation.EntityName, app.getPgRepo(app.CarCatalogDB, generation.EntityName))
	}

	app.Domain.Serie.Repository, ok = app.getPgRepo(app.CarCatalogDB, serie.EntityName).(serie.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", serie.EntityName, serie.EntityName, app.getPgRepo(app.CarCatalogDB, serie.EntityName))
	}

	app.Domain.Modification.Repository, ok = app.getPgRepo(app.CarCatalogDB, modification.EntityName).(modification.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", modification.EntityName, modification.EntityName, app.getPgRepo(app.CarCatalogDB, modification.EntityName))
	}
	//	MaintenanceCatalog
	app.Domain.Supply.Repository, ok = app.getPgRepo(app.MaintenanceDB, supply.EntityName).(supply.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", supply.EntityName, supply.EntityName, app.getPgRepo(app.MaintenanceDB, supply.EntityName))
	}

	app.Domain.Work.Repository, ok = app.getPgRepo(app.MaintenanceDB, work.EntityName).(work.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", work.EntityName, work.EntityName, app.getPgRepo(app.MaintenanceDB, work.EntityName))
	}

	app.Domain.Maintenance.Repository, ok = app.getPgRepo(app.MaintenanceDB, maintenance.EntityName).(maintenance.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", maintenance.EntityName, maintenance.EntityName, app.getPgRepo(app.MaintenanceDB, maintenance.EntityName))
	}
	//	Client
	app.Domain.Client.Repository, ok = app.getPgRepo(app.MaintenanceDB, client.EntityName).(client.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", client.EntityName, client.EntityName, app.getPgRepo(app.MaintenanceDB, client.EntityName))
	}

	app.Domain.Address.Repository, ok = app.getPgRepo(app.MaintenanceDB, address.EntityName).(address.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", address.EntityName, address.EntityName, app.getPgRepo(app.MaintenanceDB, address.EntityName))
	}

	app.Domain.Car.Repository, ok = app.getPgRepo(app.MaintenanceDB, car.EntityName).(car.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", car.EntityName, car.EntityName, app.getPgRepo(app.MaintenanceDB, car.EntityName))
	}
	//	Order
	app.Domain.Order.Repository, ok = app.getPgRepo(app.MaintenanceDB, order.EntityName).(order.Repository)
	if !ok {
		return errors.Errorf("Can not cast DB repository for entity %q to %vRepository. Repo: %v", order.EntityName, order.EntityName, app.getPgRepo(app.MaintenanceDB, order.EntityName))
	}

	if app.Auth.SessionRepository, err = redisrep.NewSessionRepository(app.Redis, app.Cfg.SessionLifeTime, app.Domain.User.Repository); err != nil {
		return errors.Errorf("Can not get new SessionRepository err: %v", err)
	}
	app.Auth.TokenRepository = jwt.NewRepository()

	app.Cache = cache.NewService(app.Redis, app.Cfg.CacheLifeTime)

	return nil
}

func (app *App) SetupServices() {
	app.Domain.User.Service = user.NewService(app.Logger, app.Domain.User.Repository)
	app.Auth.Service = auth.NewService(app.Cfg.JWTSigningKey, app.Cfg.JWTExpiration, app.Domain.User.Service, app.Logger, app.Auth.SessionRepository, app.Auth.TokenRepository)
	//	CarCatalog
	app.Domain.Mark.Service = mark.NewService(app.Logger, app.Domain.Mark.Repository)
	app.Domain.Model.Service = model.NewService(app.Logger, app.Domain.Model.Repository)
	app.Domain.Generation.Service = generation.NewService(app.Logger, app.Domain.Generation.Repository)
	app.Domain.Serie.Service = serie.NewService(app.Logger, app.Domain.Serie.Repository)
	app.Domain.Modification.Service = modification.NewService(app.Logger, app.Domain.Modification.Repository)
	//	MaintenanceCatalog
	app.Domain.Maintenance.Service = maintenance.NewService(app.Logger, app.Domain.Maintenance.Repository)
	app.Domain.Work.Service = work.NewService(app.Logger, app.Domain.Work.Repository)
	app.Domain.Supply.Service = supply.NewService(app.Logger, app.Domain.Supply.Repository)
	//	Order
	app.Domain.Order.Service = order.NewService(app.Logger, app.Domain.Order.Repository)
	//	Client
	app.Domain.Client.Service = client.NewService(app.Logger, app.Domain.Client.Repository)
	app.Domain.Address.Service = address.NewService(app.Logger, app.Domain.Address.Repository)
	app.Domain.Car.Service = car.NewService(app.Logger, app.Domain.Car.Repository)
}

// Run is func to run the App
func (app *App) Run() error {
	return nil
}

func (app *App) getPgRepo(dbase pg.IDB, entityName string) (repo pgrep.IRepository) {
	var err error

	if repo, err = pgrep.GetRepository(app.Logger, dbase, entityName); err != nil {
		golog.Fatalf("Can not get db repository for entity %q, error happened: %v", entityName, err)
	}
	return repo
}

func (app *App) Stop() error {
	errRedis := app.Redis.Close()
	errPg01 := app.IdentityDB.DB().Close()
	errPg02 := app.CarCatalogDB.DB().Close()
	errPg03 := app.MaintenanceDB.DB().Close()

	switch {
	case errPg01 != nil || errPg02 != nil || errPg03 != nil:
		return errors.Wrapf(apperror.ErrInternal, "pg close error: %v", errPg01)
	case errRedis != nil:
		return errors.Wrapf(apperror.ErrInternal, "redis close error: %v", errRedis)
	}

	return nil
}
