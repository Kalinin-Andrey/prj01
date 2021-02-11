package supply

import (
	"time"
)

const (
	EntityName = "supply"
	TableName  = "supply"
)

// Supply is the user entity
type Supply struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(255);unique;index"`
	Description string `gorm:"type:text;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
}

func (e Supply) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Supply {
	return &Supply{}
}
