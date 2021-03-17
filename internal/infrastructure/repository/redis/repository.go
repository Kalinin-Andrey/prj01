package redis

import (
	"carizza/pkg/selection_condition"

	"github.com/minipkg/go-app-common/db/redis"
)

// IRepository is an interface of repository
type IRepository interface{}

// repository persists albums in database
type repository struct {
	db         redis.IDB
	Conditions selection_condition.SelectionCondition
}

const DefaultLimit = 100

func (r *repository) SetDefaultConditions(defaultConditions selection_condition.SelectionCondition) {
	r.Conditions = defaultConditions

	if r.Conditions.Limit == 0 {
		r.Conditions.Limit = DefaultLimit
	}
}
