package domain

import (
	"github.com/go-ozzo/ozzo-validation/v4"

	"carizza/internal/pkg/log"
)

const (
	MaxLIstLimit  = 1000
	SortOrderAsc  = "asc"
	SortOrderDesc = "desc"
)

var SortOrders = []string{"", SortOrderAsc, SortOrderDesc}

type Service struct {
	logger log.ILogger
}

type DBQueryConditions struct {
	Where     map[string]interface{}
	SortOrder map[string]string
	Limit     uint
	Offset    uint
}

func (e DBQueryConditions) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.SortOrder, validation.Each(validation.In(SortOrders))),
	)
}
