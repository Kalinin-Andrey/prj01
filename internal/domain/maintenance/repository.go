package maintenance

import (
	"context"
	"github.com/minipkg/selection_condition"
)

// Repository encapsulates the logic to access albums from the data source.
type Repository interface {
	SetDefaultConditions(conditions *selection_condition.SelectionCondition)
	// Get returns the album with the specified album ID.
	Get(ctx context.Context, id uint) (*Maintenance, error)
	// Query retrieves records with the specified conditions, offset and limit from the database.
	Query(ctx context.Context, cond *selection_condition.SelectionCondition) ([]Maintenance, error)
	First(ctx context.Context, entity *Maintenance) (*Maintenance, error)
	// Create saves a new Maintenance record in the database.
	Create(ctx context.Context, entity *Maintenance) error
	// Update saves a changed Maintenance record in the database.
	Update(ctx context.Context, entity *Maintenance) error
	// Save update value in database, if the value doesn't have primary key, will insert it
	Save(ctx context.Context, entity *Maintenance) error
	// Delete (soft) deletes a Maintenance record in the database.
	Delete(ctx context.Context, id uint) error
}
