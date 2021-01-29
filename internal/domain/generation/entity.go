package generation

const (
	EntityName = "generation"
	TableName  = "car_generation"
)

// Post is the user entity
type Generation struct {
	ID      uint   `gorm:"column:id_car_generation" json:"id"`
	MarkID  uint   `gorm:"column:id_car_mark" json:"markId"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type" json:"type"`
}

func (e Generation) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Generation {
	return &Generation{}
}
