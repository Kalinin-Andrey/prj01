package gorm

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"strings"

	"carizza/pkg/selection_condition"
)

const DefaultLimit = 1000

func ApplyConditions(db *gorm.DB, conditions selection_condition.SelectionCondition) *gorm.DB {
	if err := conditions.Validate(); err != nil {
		db.AddError(err)
	}

	if conditions.Where != nil {
		db = db.Where(conditions.Where)
	}

	if conditions.SortOrder != nil {
		m := keysToSnakeCaseStr(conditions.SortOrder)
		s := strings.Builder{}

		for k, v := range m {
			s.WriteString(k + " " + v + ", ")
		}
		db = db.Order(strings.Trim(s.String(), ", "))
	}

	if conditions.Limit != 0 {
		db = db.Limit(conditions.Limit)
	} else {
		db = db.Limit(DefaultLimit)
	}

	if conditions.Offset != 0 {
		db = db.Limit(conditions.Offset)
	}

	return db
}

func ApplyWhereConditions(db *gorm.DB, conditions selection_condition.WhereConditions) *gorm.DB {
	if err := conditions.Validate(); err != nil {
		db.AddError(err)
	}

	for _, condition := range conditions {
		switch condition.Condition {
		case selection_condition.ConditionEq:
		case selection_condition.ConditionIn:
		case selection_condition.ConditionBt:
		case selection_condition.ConditionGt:
		case selection_condition.ConditionGte:
		case selection_condition.ConditionLt:
		case selection_condition.ConditionLte:
		}
	}
	return db
}

func keysToSnakeCase(in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(in))

	for key, val := range in {
		out[strcase.ToSnake(key)] = val
	}
	return out
}

func keysToSnakeCaseStr(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))

	for key, val := range in {
		out[strcase.ToSnake(key)] = val
	}
	return out
}
