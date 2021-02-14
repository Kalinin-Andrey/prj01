package work2supply

const (
	EntityName = "work2supply"
	TableName  = "work2supply"
)

// Work2Supply is the service entity
type Work2Supply struct {
	WorkID   uint `sql:"type:int not null REFERENCES \"work\"(id)" gorm:"primaryKey"`
	SupplyID uint `sql:"type:int not null REFERENCES \"supply\"(id)" gorm:"primaryKey"`
}

func (e Work2Supply) TableName() string {
	return TableName
}

// New func is a constructor for the Work2Supply
func New() *Work2Supply {
	return &Work2Supply{}
}
