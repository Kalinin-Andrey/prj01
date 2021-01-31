package serie

import (
	"context"

	"carizza/internal/domain"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	SetDefaultConditions(conditions domain.DBQueryConditions)
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id uint) (*Serie, error)
	// Query returns the list of albums with the given offset and limit.
	Query(ctx context.Context, cond domain.DBQueryConditions) ([]Serie, error)
	First(ctx context.Context, entity *Serie) (*Serie, error)
}