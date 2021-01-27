package mark

const (
	EntityName = "mark"
	TableName  = "car_mark"
)

// Mark is the mark entity
type Mark struct {
	ID      uint   `gorm:"column:id_car_mark" json:"id"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type" json:"type"`
}

func (e Mark) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Mark {
	return &Mark{}
}
