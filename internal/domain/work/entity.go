package work

const (
	EntityName = "work"
	TableName  = "work"
)

// Post is the user entity
type Work struct {
	ID      uint   `gorm:"column:id_car_work" json:"id"`
	MarkID  uint   `gorm:"column:id_car_mark" json:"markId"`
	Name    string `gorm:"type:varchar(255)"`
	NameRus string `gorm:"type:varchar(255)"`
	TypeID  uint   `gorm:"column:id_car_type" json:"type"`
}

func (e Work) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Work {
	return &Work{}
}
