package gorm

import (
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"reflect"
	"strings"

	"carizza/pkg/selection_condition"
)

const DefaultLimit = 1000

func Conditions(db *gorm.DB, conditions selection_condition.SelectionCondition) *gorm.DB {
	if err := conditions.Validate(); err != nil {
		db.AddError(err)
		return db
	}

	db = Where(db, conditions.Where)
	db = SortOrder(db, conditions.SortOrder)
	db = Limit(db, conditions.Limit)
	db = Offset(db, conditions.Offset)

	return db
}

func SortOrder(db *gorm.DB, order map[string]string) *gorm.DB {
	if order == nil {
		return db
	}
	m := keysToSnakeCaseStr(order)
	s := strings.Builder{}

	for k, v := range m {
		s.WriteString(k + " " + v + ", ")
	}
	return db.Order(strings.Trim(s.String(), ", "))
}

func Offset(db *gorm.DB, value uint) *gorm.DB {
	if value == 0 {
		return db
	}
	return db.Offset(value)
}

func Limit(db *gorm.DB, value uint) *gorm.DB {
	if value == 0 {
		return db.Limit(DefaultLimit)
	}
	return db.Limit(value)
}

func Where(db *gorm.DB, conditions interface{}) *gorm.DB {
	if conditions == nil {
		return db
	}

	wcs, ok := conditions.(selection_condition.WhereConditions)
	if ok {
		return WhereConditions(db, wcs)
	}

	wc, ok := conditions.(selection_condition.WhereCondition)
	if ok {
		return WhereCondition(db, wc)
	}

	if !isStruct(conditions) {
		db.AddError(errors.Errorf("conditions must be a selection_condition.WhereConditions, selection_condition.WhereCondition or a struct"))
		return db
	}
	return db.Where(conditions)
}

func isStruct(e interface{}) bool {
	t := reflect.TypeOf(e)

	if t.Kind() == reflect.Ptr {
		t = reflect.Indirect(reflect.ValueOf(e)).Type()
	}
	return t.Kind() == reflect.Struct
}

func WhereConditions(db *gorm.DB, conditions selection_condition.WhereConditions) *gorm.DB {
	if err := conditions.Validate(); err != nil {
		db.AddError(err)
		return db
	}

	for _, condition := range conditions {
		db = WhereCondition(db, condition)
		if db.Error != nil {
			return db
		}
	}
	return db
}

func WhereCondition(db *gorm.DB, condition selection_condition.WhereCondition) *gorm.DB {
	if err := condition.Validate(); err != nil {
		db.AddError(err)
		return db
	}

	switch condition.Condition {
	case selection_condition.ConditionEq:
		db = db.Where(map[string]interface{}{condition.Field: condition.Value})
	case selection_condition.ConditionIn:
		conds, ok := condition.Value.([]interface{})
		if !ok {
			db.AddError(errors.Errorf("Can not assign value condition to slice"))
		}
		//db = db.Where(map[string]interface{}{condition.Field: condition.Value})
		db = db.Where(condition.Field+" IN ?", conds)
	case selection_condition.ConditionBt:
		conds, ok := condition.Value.([]interface{})
		if !ok {
			db.AddError(errors.Errorf("Can not assign value condition to slice"))
		}
		db = db.Where(condition.Field+" BETWEEN ? AND ?", conds[0], conds[1])
	case selection_condition.ConditionGt:
		db = db.Where(condition.Field+" > ?", condition.Value)
	case selection_condition.ConditionGte:
		db = db.Where(condition.Field+" >= ?", condition.Value)
	case selection_condition.ConditionLt:
		db = db.Where(condition.Field+" < ?", condition.Value)
	case selection_condition.ConditionLte:
		db = db.Where(condition.Field+" <= ?", condition.Value)
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
