package order

import (
	"time"

	"carizza/internal/domain/address"
	"carizza/internal/domain/car"
	"carizza/internal/domain/client"
	"carizza/internal/domain/maintenance"
)

const (
	EntityName = "order"
	TableName  = "order"
	Period0    = "с 10:00 до 20:00"
	Period1    = "с 10:00 до 15:00"
	Period2    = "с 15:00 до 20:00"
)

var Periods = []string{Period0, Period1, Period2}

// Order is the service entity
type Order struct {
	ID           uint      `gorm:"primaryKey"`
	ClientID     uint      `sql:"type:int not null REFERENCES \"client\"(id)" gorm:"index"`
	CarID        uint      `sql:"type:int not null REFERENCES \"car\"(id)"`
	AddressID    uint      `sql:"type:int not null REFERENCES \"address\"(id)"`
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

func (e Order) TableName() string {
	return TableName
}

// New func is a constructor for the Order
func New() *Order {
	return &Order{}
}
