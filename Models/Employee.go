package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID         string    `gorm:"type:char(36);primary_key" json:"id"`
	DivisionId string    `gorm:"type:char(36);not null" json:"division_id"` // foreignKey
	Name       string    `gorm:"type:string;not null" json:"name"`
	Phone      string    `gorm:"type:string;not null;unique" json:"phone"`
	Image      string    `gorm:"type:string;not null" json:"image"`
	Position   string    `gorm:"type:string;not null" json:"position"`
	Division   Division  `gorm:"foreignKey:DivisionId;references:ID" json:"division"` // hasOne relationship divison by DivisionID
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {

	// to set UUID if not set
	if e.ID == "" {
		e.ID = uuid.New().String()
	}

	return
}
