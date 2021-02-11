package generation

const (
	EntityName = "generation"
	TableName  = "car_generation"
)

// Post is the user entity
type Generation struct {
	ID        uint   `gorm:"column:id_car_generation"`
	ModelID   uint   `gorm:"column:id_car_model"`
	Name      string `gorm:"type:varchar(255)"`
	Label     string `gorm:"-"`
	YearBegin string `gorm:"type:varchar(255)"`
	YearEnd   string `gorm:"type:varchar(255)"`
	TypeID    uint   `gorm:"column:id_car_type"`
}

func (e Generation) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Generation {
	return &Generation{}
}
