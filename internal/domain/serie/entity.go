package serie

const (
	EntityName = "serie"
	TableName  = "car_serie"
)

// Post is the user entity
type Serie struct {
	ID           uint   `gorm:"column:id_car_serie"`
	ModelID      uint   `gorm:"column:id_car_model"`
	GenerationID uint   `gorm:"column:id_car_generation"`
	Name         string `gorm:"type:varchar(255)"`
	TypeID       uint   `gorm:"column:id_car_type"`
}

func (e Serie) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Serie {
	return &Serie{}
}
