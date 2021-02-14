package ordered_work

import (
	"time"
)

const (
	EntityName = "ordered_work"
	TableName  = "ordered_work"
)

// OrderedWork is the service entity
type OrderedWork struct {
	ID                   uint `gorm:"primaryKey"`
	OrderedMaintenanceID uint `sql:"type:int not null REFERENCES \"ordered_maintenance\"(id)" gorm:"index"`
	WorkID               uint `sql:"type:int not null REFERENCES \"work\"(id)"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time `gorm:"index"`
}

func (e OrderedWork) TableName() string {
	return TableName
}

// New func is a constructor for the OrderedMaintenance
func New() *OrderedWork {
	return &OrderedWork{}
}
