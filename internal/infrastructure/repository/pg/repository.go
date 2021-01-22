package pg

import (
	"carizza/internal/domain"
	"carizza/internal/domain/model"
	"carizza/internal/domain/user"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"carizza/internal/pkg/db/pg"
	"carizza/internal/pkg/log"
)

// IRepository is an interface of repository
type IRepository interface{}

// repository persists albums in database
type repository struct {
	db         pg.IDB
	logger     log.ILogger
	Conditions domain.DBQueryConditions
}

const DefaultLimit = 100

// GetRepository return a repository
func GetRepository(logger log.ILogger, dbase pg.IDB, entity string) (repo IRepository, err error) {
	r := &repository{
		db:     dbase,
		logger: logger,
	}

	switch entity {
	case user.EntityName:
		repo, err = NewUserRepository(r)
	case model.EntityName:
		repo, err = NewModelRepository(r)
	default:
		err = errors.Errorf("Repository for entity %q not found", entity)
	}
	return repo, err
}

func (r *repository) SetDefaultConditions(defaultConditions domain.DBQueryConditions) {
	r.Conditions = defaultConditions

	if r.Conditions.Limit == 0 {
		r.Conditions.Limit = DefaultLimit
	}
}

func (r repository) dbWithDefaults() *gorm.DB {
	return r.applyConditions(r.db.DB(), r.Conditions)
}

func (r repository) applyConditions(db *gorm.DB, conditions domain.DBQueryConditions) *gorm.DB {

	if conditions.Where != nil {
		m := r.keysToSnakeCase(conditions.Where)
		db = db.Where(m)
	}

	if conditions.SortOrder != nil {
		m := r.keysToSnakeCaseStr(conditions.SortOrder)
		db = db.Order(m)
	}

	if conditions.Limit != 0 {
		db = db.Limit(conditions.Limit)
	}

	if conditions.Offset != 0 {
		db = db.Limit(conditions.Offset)
	}

	return db
}

func (r repository) keysToSnakeCase(in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(in))

	for key, val := range in {
		out[strcase.ToSnake(key)] = val
	}
	return out
}

func (r repository) keysToSnakeCaseStr(in map[string]string) map[string]interface{} {
	out := make(map[string]interface{}, len(in))

	for key, val := range in {
		out[strcase.ToSnake(key)] = val
	}
	return out
}
