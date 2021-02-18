package ordered_maintenance

import (
	"time"
)

const (
	EntityName = "ordered_maintenance"
	TableName  = "ordered_maintenance"
)

// OrderedMaintenance is the service entity
type OrderedMaintenance struct {
	ID            uint       `gorm:"primaryKey" json:"id"`
	OrderID       uint       `sql:"type:int not null REFERENCES \"order\"(id)" gorm:"index" json:"orderId"`
	MaintenanceID uint       `sql:"type:int not null REFERENCES \"maintenance\"(id)" json:"maintenanceId"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

func (e OrderedMaintenance) TableName() string {
	return TableName
}

// New func is a constructor for the OrderedMaintenance
func New() *OrderedMaintenance {
	return &OrderedMaintenance{}
}
