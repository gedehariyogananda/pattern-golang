package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Division struct {
	ID         string     `gorm:"type:char(36);primary_key"`
	NameDivisi string     `gorm:"type:string;not null"`
	Employees  []Employee `gorm:"-"` // many employyes foreignKey Divison_id
	CreatedAt  time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt  time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (d *Division) BeforeCreate(tx *gorm.DB) (err error) {

	// to set UUID if not set
	if d.ID == "" {
		d.ID = uuid.New().String()
	}

	return
}
