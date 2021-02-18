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
	ID        uint               `gorm:"primaryKey" json:"id"`
	Name      string             `gorm:"type:varchar(255) not null;unique;index" json:"name"`
	Phone     uint               `gorm:"type:smallint not null;unique;index" json:"phone"`
	Cars      []*car.Car         `json:"cars,omitempty"`
	Addresses []*address.Address `json:"addresses,omitempty"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
	DeletedAt *time.Time         `gorm:"index" json:"deletedAt,omitempty"`
}

func (e Client) TableName() string {
	return TableName
}

// New func is a constructor for the Client
func New() *Client {
	return &Client{}
}
