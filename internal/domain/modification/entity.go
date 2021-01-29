package modification

const (
	EntityName = "modification"
	TableName  = "car_modification"
)

// Post is the user entity
type Modification struct {
	ID      uint   `gorm:"column:id_car_modification" json:"id"`
	MarkID  uint   `gorm:"column:id_car_mark" json:"markId"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type" json:"type"`
}

func (e Modification) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Modification {
	return &Modification{}
}
