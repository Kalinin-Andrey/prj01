package client

import(
	"time"
)

const (
	EntityName = "client"
	TableName  = "client"
)

// Client is the service entity
type Client struct {
	ID        uint   `gorm:"PRIMARY_KEY" json:"id"`
	Name      string `gorm:"type:varchar(100);UNIQUE;INDEX" json:"username"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"INDEX"`
}

func (e Client) TableName() string {
	return TableName
}

// New func is a constructor for the Client
func New() *Client {
	return &Client{}
}
