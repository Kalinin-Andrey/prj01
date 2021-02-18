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
	ID                   uint       `gorm:"primaryKey" json:"id"`
	OrderedMaintenanceID uint       `sql:"type:int not null REFERENCES \"ordered_maintenance\"(id)" gorm:"index" json:"orderedMaintenanceID"`
	WorkID               uint       `sql:"type:int not null REFERENCES \"work\"(id)" json:"workID"`
	CreatedAt            time.Time  `json:"createdAt"`
	UpdatedAt            time.Time  `json:"updatedAt"`
	DeletedAt            *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

func (e OrderedWork) TableName() string {
	return TableName
}

// New func is a constructor for the OrderedMaintenance
func New() *OrderedWork {
	return &OrderedWork{}
}
