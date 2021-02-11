package modification

const (
	EntityName = "modification"
	TableName  = "car_modification"
)

// Post is the user entity
type Modification struct {
	ID                  uint   `gorm:"column:id_car_modification"`
	ModelID             uint   `gorm:"column:id_car_model"`
	SerieID             uint   `gorm:"column:id_car_serie"`
	Name                string `gorm:"type:varchar(255)"`
	StartProductionYear uint
	EndProductionYear   uint
	TypeID              uint `gorm:"column:id_car_type"`
}

func (e Modification) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Modification {
	return &Modification{}
}
