package generation

const (
	EntityName = "generation"
	TableName  = "car_generation"
)

// Post is the user entity
type Generation struct {
	ID        uint   `gorm:"column:id_car_generation" json:"id"`
	ModelID   uint   `gorm:"column:id_car_model" json:"modelId"`
	Name      string `gorm:"type:varchar(255)" json:"name"`
	Label     string `gorm:"-" json:"label"`
	YearBegin string `gorm:"type:varchar(255)" json:"yearBegin"`
	YearEnd   string `gorm:"type:varchar(255)" json:"yearEnd"`
	TypeID    uint   `gorm:"column:id_car_type" json:"typeId"`
}

func (e Generation) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Generation {
	return &Generation{}
}
