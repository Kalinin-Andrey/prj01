package modification

const (
	EntityName = "modification"
	TableName  = "car_modification"
)

// Post is the user entity
type Modification struct {
	ID                  uint   `gorm:"column:id_car_modification" json:"id"`
	ModelID             uint   `gorm:"column:id_car_model" json:"modelId"`
	SerieID             uint   `gorm:"column:id_car_serie" json:"serieId"`
	Name                string `gorm:"type:varchar(255)" json:"name"`
	StartProductionYear uint   `json:"startProductionYear"`
	EndProductionYear   uint   `json:"endProductionYear"`
	TypeID              uint   `gorm:"column:id_car_type" json:"typeId"`
}

func (e Modification) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Modification {
	return &Modification{}
}
