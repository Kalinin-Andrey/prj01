package car

import (
	"time"
)

const (
	EntityName = "car"
	TableName  = "car"
)

// Post is the user entity
type Car struct {
	ID             uint `gorm:"primaryKey"`
	ClientID       uint `gorm:"type:integer;index"`
	MarkID         uint `gorm:"type:integer"`
	ModelID        uint `gorm:"type:integer"`
	GenerationID   uint `gorm:"type:integer"`
	SerieID        uint `gorm:"type:integer"`
	ModificationID uint `gorm:"type:integer"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `gorm:"index"`
}

func (e Car) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Car {
	return &Car{}
}
