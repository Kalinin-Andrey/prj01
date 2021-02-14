package ordered_supply

import (
	"time"
)

const (
	EntityName = "ordered_supply"
	TableName  = "ordered_supply"
)

// OrderedSupply is the service entity
type OrderedSupply struct {
	ID            uint `gorm:"primaryKey"`
	OrderedWorkID uint `sql:"type:int not null REFERENCES \"ordered_work\"(id)" gorm:"index"`
	SupplyID      uint `sql:"type:int not null REFERENCES \"supply\"(id)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `gorm:"index"`
}

func (e OrderedSupply) TableName() string {
	return TableName
}

// New func is a constructor for the OrderedMaintenance
func New() *OrderedSupply {
	return &OrderedSupply{}
}
