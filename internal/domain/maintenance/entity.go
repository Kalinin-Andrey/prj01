package maintenance

import(
	"time"
)

const (
	EntityName = "service"
	TableName  = "service"
)

// Maintenance is the service entity
type Maintenance struct {
	ID        uint   `gorm:"PRIMARY_KEY" json:"id"`
	Name      string `gorm:"type:varchar(255);UNIQUE;INDEX" json:"name"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"INDEX"`
}

func (e Maintenance) TableName() string {
	return TableName
}

// New func is a constructor for the Maintenance
func New() *Maintenance {
	return &Maintenance{}
}
