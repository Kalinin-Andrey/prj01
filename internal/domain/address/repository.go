package address

import (
	"carizza/pkg/selection_condition"
	"context"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	SetDefaultConditions(conditions selection_condition.SelectionCondition)
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id uint) (*Address, error)
	// Query returns the list of albums with the given offset and limit.
	Query(ctx context.Context, cond selection_condition.SelectionCondition) ([]Address, error)
	First(ctx context.Context, entity *Address) (*Address, error)
}
