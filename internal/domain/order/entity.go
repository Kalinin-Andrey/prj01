package order

import(
	"time"
)

const (
	EntityName = "service"
	TableName  = "service"
)

// Order is the service entity
type Order struct {
	ID        uint   `gorm:"PRIMARY_KEY" json:"id"`
	Name      string `gorm:"type:varchar(100);UNIQUE;INDEX" json:"username"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"INDEX"`
}

func (e Order) TableName() string {
	return TableName
}

// New func is a constructor for the Order
func New() *Order {
	return &Order{}
}
