package serie

const (
	EntityName = "serie"
	TableName  = "car_serie"
)

// Post is the user entity
type Serie struct {
	ID      uint   `gorm:"column:id_car_serie" json:"id"`
	MarkID  uint   `gorm:"column:id_car_mark" json:"markId"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type" json:"type"`
}

func (e Serie) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Serie {
	return &Serie{}
}
