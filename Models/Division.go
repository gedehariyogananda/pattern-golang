package Models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Division struct {
	ID         string     `gorm:"type:char(36);primary_key" json:"id"` // primary
	NameDivisi string     `gorm:"type:string;not null" json:"name_divisi"`
	Employees  []Employee `gorm:"-" json:"-"` // many employyes foreignKey Divison_id
	CreatedAt  time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (d *Division) BeforeCreate(tx *gorm.DB) (err error) {

	// to set UUID if not set
	if d.ID == "" {
		d.ID = uuid.New().String()
	}

	return
}
