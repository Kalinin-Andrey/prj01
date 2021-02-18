package address

import "time"

const (
	EntityName = "address"
	TableName  = "address"
)

// Address is the user entity
type Address struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	ClientID  uint       `sql:"type:int not null REFERENCES \"client\"(id)" gorm:"index" json:"clientId"`
	Value     string     `gorm:"type:varchar(255) not null" json:"value"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

func (e Address) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Address {
	return &Address{}
}
