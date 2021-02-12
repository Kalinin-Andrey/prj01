package ordered_maintenance

import (
	"carizza/internal/domain/maintenance"
	"time"

	"carizza/internal/domain/address"
	"carizza/internal/domain/car"
	"carizza/internal/domain/client"
)

const (
	EntityName = "ordered_maintenance"
	TableName  = "ordered_maintenance"
)

// OrderedMaintenance is the service entity
type OrderedMaintenance struct {
	ID           uint      `gorm:"primaryKey"`
	ClientID     uint      `gorm:"type:integer;index"`
	CarID        uint      `gorm:"type:integer"`
	AddressID    uint      `gorm:"type:integer"`
	Date         time.Time `gorm:"type:date;index"`
	PeriodID     uint      `gorm:"type:smallint"`
	Client       *client.Client
	Car          *car.Car
	Address      *address.Address
	Period       string
	Maintenances []*maintenance.Maintenance
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `gorm:"index"`
}

func (e OrderedMaintenance) TableName() string {
	return TableName
}

// New func is a constructor for the OrderedMaintenance
func New() *OrderedMaintenance {
	return &OrderedMaintenance{}
}
