package pg

import (
	"carizza/internal/domain"
	"carizza/internal/domain/generation"
	"carizza/internal/domain/mark"
	"carizza/internal/domain/model"
	"carizza/internal/domain/modification"
	"carizza/internal/domain/serie"
	"carizza/internal/domain/user"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"strings"

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

const DefaultLimit = 1000

// GetRepository return a repository
func GetRepository(logger log.ILogger, dbase pg.IDB, entity string) (repo IRepository, err error) {
	r := &repository{
		db:     dbase,
		logger: logger,
	}

	switch entity {
	case user.EntityName:
		repo, err = NewUserRepository(r)
	case mark.EntityName:
		repo, err = NewMarkRepository(r)
	case model.EntityName:
		repo, err = NewModelRepository(r)
	case generation.EntityName:
		repo, err = NewGenerationRepository(r)
	case serie.EntityName:
		repo, err = NewSerieRepository(r)
	case modification.EntityName:
		repo, err = NewModificationRepository(r)
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
	db, _ := r.applyConditions(r.db.DB(), r.Conditions)
	return db
}

func (r repository) applyConditions(db *gorm.DB, conditions domain.DBQueryConditions) (*gorm.DB, error) {

	if err := conditions.Validate(); err != nil {
		return nil, err
	}

	if conditions.Where != nil {
		m := r.keysToSnakeCase(conditions.Where)
		db = db.Where(m)
	}

	if conditions.SortOrder != nil {
		m := r.keysToSnakeCaseStr(conditions.SortOrder)
		s := strings.Builder{}

		for k, v := range m {
			s.WriteString(k + " " + v + ", ")
		}
		db = db.Order(strings.Trim(s.String(), ", "))
	}

	if conditions.Limit != 0 {
		db = db.Limit(conditions.Limit)
	}

	if conditions.Offset != 0 {
		db = db.Limit(conditions.Offset)
	}

	return db, nil
}

func (r repository) keysToSnakeCase(in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(in))

	for key, val := range in {
		out[strcase.ToSnake(key)] = val
	}
	return out
}

func (r repository) keysToSnakeCaseStr(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))

	for key, val := range in {
		out[strcase.ToSnake(key)] = val
	}
	return out
}
