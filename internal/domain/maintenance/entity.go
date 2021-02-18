package maintenance

import (
	"time"

	"carizza/internal/domain/work"
)

const (
	EntityName = "maintenance"
	TableName  = "maintenance"
)

// Maintenance entity
type Maintenance struct {
	ID          uint         `gorm:"primaryKey" json:"id"`
	Name        string       `gorm:"type:varchar(255) not null;unique;index" json:"name"`
	Description string       `gorm:"type:text;" json:"description"`
	Works       []*work.Work `gorm:"many2many:maintenance2work" json:"works,omitempty"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	DeletedAt   *time.Time   `gorm:"index" json:"deletedAt,omitempty"`
}

func (e Maintenance) TableName() string {
	return TableName
}

// New func is a constructor for the Maintenance
func New() *Maintenance {
	return &Maintenance{}
}

func (e Maintenance) Validate() error {
	return nil
}
