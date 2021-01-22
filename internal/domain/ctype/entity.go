package ctype

const (
	EntityName = "ctype"
	TableName  = "car_type"
	TypeIDCar  = 1
)

// Post is the user entity
type Type struct {
	ID   uint   `gorm:"primaryKey,column:id_car_type" json:"id"`
	Name string `gorm:"type:varchar(255)"`
}

func (e Type) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Type {
	return &Type{}
}
