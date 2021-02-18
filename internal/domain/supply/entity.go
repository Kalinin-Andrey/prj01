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
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"type:varchar(255) not null;unique;index" json:"name"`
	Description string     `gorm:"type:text;" json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

func (e Supply) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Supply {
	return &Supply{}
}
