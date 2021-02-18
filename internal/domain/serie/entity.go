package serie

const (
	EntityName = "serie"
	TableName  = "car_serie"
)

// Post is the user entity
type Serie struct {
	ID           uint   `gorm:"column:id_car_serie" json:"id"`
	ModelID      uint   `gorm:"column:id_car_model" json:"modelId"`
	GenerationID uint   `gorm:"column:id_car_generation" json:"generationId"`
	Name         string `gorm:"type:varchar(255)" json:"name"`
	TypeID       uint   `gorm:"column:id_car_type" json:"typeId"`
}

func (e Serie) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Serie {
	return &Serie{}
}
