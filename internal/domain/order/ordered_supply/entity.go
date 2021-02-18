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
	ID            uint       `gorm:"primaryKey" json:"id"`
	OrderedWorkID uint       `sql:"type:int not null REFERENCES \"ordered_work\"(id)" gorm:"index" json:"orderedWorkID"`
	SupplyID      uint       `sql:"type:int not null REFERENCES \"supply\"(id)" json:"supplyID"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

func (e OrderedSupply) TableName() string {
	return TableName
}

// New func is a constructor for the OrderedMaintenance
func New() *OrderedSupply {
	return &OrderedSupply{}
}
