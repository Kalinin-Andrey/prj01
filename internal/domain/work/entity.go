package work

import (
	"carizza/internal/domain/supply"
	"time"
)

const (
	EntityName = "work"
	TableName  = "work"
)

// Post is the user entity
type Work struct {
	ID          uint             `gorm:"primaryKey" json:"id"`
	Name        string           `gorm:"type:varchar(255) not null;unique;index" json:"name"`
	Description string           `gorm:"type:text;" json:"description"`
	Supplies    []*supply.Supply `gorm:"many2many:work2supply" json:"supplies,omitempty"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
	DeletedAt   *time.Time       `gorm:"index" json:"deletedAt,omitempty"`
}

func (e Work) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Work {
	return &Work{}
}
