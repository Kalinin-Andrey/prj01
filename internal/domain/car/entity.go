package car

const (
	EntityName = "car"
	TableName  = "car"
)

// Post is the user entity
type Car struct {
	ID      uint   `gorm:"column:id" json:"id"`
	MarkID  uint   `gorm:"column:id_car_mark" json:"markId"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type" json:"type"`
}

func (e Car) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Car {
	return &Car{}
}
