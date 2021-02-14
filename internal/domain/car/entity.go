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
	ClientID       uint `sql:"type:int not null REFERENCES \"client\"(id)" gorm:"index"`
	MarkID         uint `gorm:"type:int not null"`
	ModelID        uint `gorm:"type:int not null"`
	GenerationID   uint `gorm:"type:int"`
	SerieID        uint `gorm:"type:int"`
	ModificationID uint `gorm:"type:int"`
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
