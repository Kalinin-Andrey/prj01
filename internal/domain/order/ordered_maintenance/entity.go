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
	ID            uint `gorm:"primaryKey"`
	OrderID       uint `sql:"type:int not null REFERENCES \"order\"(id)" gorm:"index"`
	MaintenanceID uint `sql:"type:int not null REFERENCES \"maintenance\"(id)"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `gorm:"index"`
}

func (e OrderedMaintenance) TableName() string {
	return TableName
}

// New func is a constructor for the OrderedMaintenance
func New() *OrderedMaintenance {
	return &OrderedMaintenance{}
}
