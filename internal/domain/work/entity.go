package work

import (
	"time"

	"carizza/internal/domain/supply"
)

const (
	EntityName = "work"
	TableName  = "work"
)

// Post is the user entity
type Work struct {
	ID          uint             `gorm:"primaryKey"`
	Name        string           `gorm:"type:varchar(255);unique;index"`
	Description string           `gorm:"type:text;"`
	Supplies    []*supply.Supply `gorm:"many2many:work2supply;"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `gorm:"index"`
}

func (e Work) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Work {
	return &Work{}
}
