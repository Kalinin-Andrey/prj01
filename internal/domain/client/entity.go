package client

import (
	"carizza/internal/domain/address"
	"carizza/internal/domain/car"
	"time"
)

const (
	EntityName = "client"
	TableName  = "client"
)

// Client is the service entity
type Client struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(255) not null;unique;index"`
	Phone     uint   `gorm:"type:smallint not null;unique;index"`
	Cars      []*car.Car
	Addresses []*address.Address
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

func (e Client) TableName() string {
	return TableName
}

// New func is a constructor for the Client
func New() *Client {
	return &Client{}
}
