package model

import (
	"context"

	"carizza/internal/domain"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	SetDefaultConditions(conditions domain.DBQueryConditions)
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id uint) (*Model, error)
	// Query returns the list of albums with the given offset and limit.
	Query(ctx context.Context, cond domain.DBQueryConditions) ([]Model, error)
	First(ctx context.Context, entity *Model) (*Model, error)
}