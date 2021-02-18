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
	ID           uint                       `gorm:"primaryKey" json:"id"`
	ClientID     uint                       `sql:"type:int not null REFERENCES \"client\"(id)" gorm:"index" json:"clientId"`
	CarID        uint                       `sql:"type:int not null REFERENCES \"car\"(id)" json:"carId"`
	AddressID    uint                       `sql:"type:int not null REFERENCES \"address\"(id)" json:"addressId"`
	Date         time.Time                  `gorm:"type:date;index" json:"date"`
	PeriodID     uint                       `gorm:"type:smallint" json:"periodId"`
	Client       *client.Client             `json:"client,omitempty"`
	Car          *car.Car                   `json:"car,omitempty"`
	Address      *address.Address           `json:"address,omitempty"`
	Period       string                     `json:"period"`
	Maintenances []*maintenance.Maintenance `json:"maintenances,omitempty"`
	CreatedAt    time.Time                  `json:"createdAt"`
	UpdatedAt    time.Time                  `json:"updatedAt"`
	DeletedAt    *time.Time                 `gorm:"index" json:"deletedAt,omitempty"`
}

func (e Order) TableName() string {
	return TableName
}

// New func is a constructor for the Order
func New() *Order {
	return &Order{}
}
