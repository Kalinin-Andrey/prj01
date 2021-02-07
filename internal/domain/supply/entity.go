package supply

const (
	EntityName = "supply"
	TableName  = "supply"
)

// Supply is the user entity
type Supply struct {
	ID      uint   `gorm:"column:id_car_model" json:"id"`
	MarkID  uint   `gorm:"column:id_car_mark" json:"markId"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type" json:"type"`
}

func (e Supply) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Supply {
	return &Supply{}
}
