package model

const (
	EntityName = "model"
	TableName  = "car_model"
)

// Post is the user entity
type Model struct {
	ID      uint   `gorm:"column:id_car_model"`
	MarkID  uint   `gorm:"column:id_car_mark"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type"`
}

func (e Model) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Model {
	return &Model{}
}
