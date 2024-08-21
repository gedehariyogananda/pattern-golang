package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Employee struct {
	ID         string    `gorm:"type:char(36);primary_key"`
	DivisionId string    `gorm:"type:char(36);not null"`
	Name       string    `gorm:"type:string;not null"`
	Phone      string    `gorm:"type:string;not null;unique"`
	Image      string    `gorm:"type:string;not null"`
	Position   string    `gorm:"type:string;not null"`
	Division   Division  `gorm:"foreignKey:DivisionId;references:ID"` // hasOne relationship divison by DivisionID
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (e *Employee) BeforeCreate(tx *gorm.DB) (err error) {

	// to set UUID if not set
	if e.ID == "" {
		e.ID = uuid.New().String()
	}

	return
}
