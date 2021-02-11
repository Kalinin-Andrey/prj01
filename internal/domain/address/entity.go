package address

import "time"

const (
	EntityName = "address"
	TableName  = "address"
)

// Address is the user entity
type Address struct {
	ID        uint   `gorm:"primaryKey"`
	ClientID  uint   `gorm:"type:integer;index"`
	Value     string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (e Address) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Address {
	return &Address{}
}