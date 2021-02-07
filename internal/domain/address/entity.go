package address

const (
	EntityName = "address"
	TableName  = "address"
)

// Address is the user entity
type Address struct {
	ID      uint   `gorm:"column:id" json:"id"`
	MarkID  uint   `gorm:"column:id_car_mark" json:"markId"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type" json:"type"`
}

func (e Address) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Address {
	return &Address{}
}
