package maintenance2work

const (
	EntityName = "maintenance2work"
	TableName  = "maintenance2work"
)

// Maintenance2Work is the service entity
type Maintenance2Work struct {
	MaintenanceID uint `sql:"type:int not null REFERENCES \"maintenance\"(id)" gorm:"primaryKey"`
	WorkID        uint `sql:"type:int not null REFERENCES \"work\"(id)" gorm:"primaryKey"`
}

func (e Maintenance2Work) TableName() string {
	return TableName
}

// New func is a constructor for the Maintenance2Work
func New() *Maintenance2Work {
	return &Maintenance2Work{}
}
