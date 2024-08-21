package Models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RolesEnum string

const (
	AdminRole RolesEnum = "admin"
	UserRole  RolesEnum = "user"
)

type User struct {
	ID        string    `gorm:"type:char(36);primary_key"`
	Name      string    `gorm:"type:string;not null"`
	Username  string    `gorm:"type:string;not null;unique"`
	Email     string    `gorm:"type:string;not null;unique"`
	Password  string    `gorm:"type:string;not null"`
	Phone     string    `gorm:"type:string;not null;unique"`
	Roles     RolesEnum `gorm:"type:string;not null;default:'user'"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
}

// before create hook set default UUID if not set
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	// to set UUID if not set
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	// to set roles if not set
	if u.Roles == "" {
		u.Roles = UserRole
	}

	// hashed password
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}

	return

}
