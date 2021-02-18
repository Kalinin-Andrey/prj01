package car

import (
	"carizza/internal/domain/generation"
	"carizza/internal/domain/mark"
	"carizza/internal/domain/model"
	"carizza/internal/domain/modification"
	"carizza/internal/domain/serie"
	"time"
)

const (
	EntityName = "car"
	TableName  = "car"
)

// Post is the user entity
type Car struct {
	ID             uint                      `gorm:"primaryKey" json:"id"`
	ClientID       uint                      `sql:"type:int not null REFERENCES \"client\"(id)" gorm:"index" json:"clientId"`
	MarkID         uint                      `gorm:"type:int not null" json:"markId"`
	ModelID        uint                      `gorm:"type:int not null" json:"modelId"`
	GenerationID   uint                      `gorm:"type:int" json:"generationId"`
	SerieID        uint                      `gorm:"type:int" json:"serieId"`
	ModificationID uint                      `gorm:"type:int" json:"modificationId"`
	Mark           mark.Mark                 `json:"mark,omitempty"`
	Model          model.Model               `json:"model,omitempty"`
	Generation     generation.Generation     `json:"generation,omitempty"`
	Serie          serie.Serie               `json:"serie,omitempty"`
	Modification   modification.Modification `json:"modification,omitempty"`
	CreatedAt      time.Time                 `json:"createdAt"`
	UpdatedAt      time.Time                 `json:"updatedAt"`
	DeletedAt      *time.Time                `gorm:"index" json:"deletedAt,omitempty"`
}

func (e Car) TableName() string {
	return TableName
}

// New func is a constructor for the Post
func New() *Car {
	return &Car{}
}
