package domain

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"

	"carizza/pkg/log"
)

const (
	MaxLIstLimit  = 1000
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

var SortOrders = []interface{}{"", SortOrderAsc, SortOrderDesc}

type Service struct {
	logger log.ILogger
}

type DBQueryConditions struct {
	Where     interface{}
	SortOrder map[string]string
	Limit     uint
	Offset    uint
}

func (e DBQueryConditions) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.SortOrder, validation.Each(validation.In(SortOrders...))),
	)
}
