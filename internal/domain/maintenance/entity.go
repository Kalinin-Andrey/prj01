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
	ID          uint         `gorm:"primaryKey"`
	Name        string       `gorm:"type:varchar(255);unique;index"`
	Description string       `gorm:"type:text;"`
	Works       []*work.Work `gorm:"many2many:maintenance2work;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
}

func (e Maintenance) TableName() string {
	return TableName
}

// New func is a constructor for the Maintenance
func New() *Maintenance {
	return &Maintenance{}
}
