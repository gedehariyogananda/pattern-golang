package Seeder

import (
	"time"

	"github.com/gedehariyogananda/pattern-golang/Models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DatabaseSeeder(db *gorm.DB) {
	uuidUsers := []string{
		"d4b3f6e8-5b5f-4d2b-8f5e-5e7b6b2b5b2b",
		"e4a3c9f2-4d4e-4b3e-8d4e-4b3e4d4e5b5b",
		"f2d4b6f8-4b4f-4d4b-8f4e-5b4e4d4f5b5b",
		"a3e4f2c9-4d4e-4f3b-8f4e-5b4e4d4f5c5b",
		"b2e4f2a3-4d4e-4f3e-8d4e-5b4e4d4f5d5b",
	}

	uuidDivision := []string{
		"b3e4f2a3-4d4e-4f3e-8d4e-5b4e4d4f5d5c",
		"c3e4f2a3-4d4e-4f3e-8d4e-5b4e4d4f5d5d",
		"d3e4f2a3-4d4e-4f3e-8d4e-5b4e4d4f5d5e",
		"e3e4f2a3-4d4e-4f3e-8d4e-5b4e4d4f5d5f",
	}

	users := []Models.User{
		{
			ID:        uuid.MustParse(uuidUsers[0]).String(),
			Name:      "Aryyy",
			Username:  "aryyy",
			Email:     "aryyy@gmail.com",
			Password:  "password",
			Phone:     "08123456789",
			Roles:     Models.AdminRole,
			CreatedAt: time.Now(),
		},

		{
			ID:        uuid.MustParse(uuidUsers[1]).String(),
			Name:      "Bryyy",
			Username:  "bryyy",
			Email:     "bryyy@gmail.com",
			Password:  "password",
			Phone:     "08123456788",
			Roles:     Models.UserRole,
			CreatedAt: time.Now(),
		},

		{
			ID:        uuid.MustParse(uuidUsers[2]).String(),
			Name:      "Cryyy",
			Username:  "cryyy",
			Email:     "cryyy@gmail.com",
			Password:  "password",
			Phone:     "08123456787",
			Roles:     Models.UserRole,
			CreatedAt: time.Now(),
		},

		{
			ID:        uuid.MustParse(uuidUsers[3]).String(),
			Name:      "Dryyy",
			Username:  "dryyy",
			Email:     "dryyy@gmail.com",
			Password:  "password",
			Phone:     "08123456786",
			Roles:     Models.UserRole,
			CreatedAt: time.Now(),
		},

		{
			ID:        uuid.MustParse(uuidUsers[4]).String(),
			Name:      "Eryyy",
			Username:  "eryyy",
			Email:     "eryyy@gmail.com",
			Password:  "password",
			Phone:     "08123456785",
			Roles:     Models.UserRole,
			CreatedAt: time.Now(),
		},
	}

	divisions := []Models.Division{
		{
			ID:         uuid.MustParse(uuidDivision[0]).String(),
			NameDivisi: "Backend",
			CreatedAt:  time.Now(),
		},
		{
			ID:         uuid.MustParse(uuidDivision[1]).String(),
			NameDivisi: "Frontend",
			CreatedAt:  time.Now(),
		},
		{
			ID:         uuid.MustParse(uuidDivision[2]).String(),
			NameDivisi: "Mobile",
			CreatedAt:  time.Now(),
		},
		{
			ID:         uuid.MustParse(uuidDivision[3]).String(),
			NameDivisi: "DevOps",
			CreatedAt:  time.Now(),
		},
	}

	employees := []Models.Employee{
		{
			ID:         uuid.New().String(),
			DivisionId: uuid.MustParse(uuidDivision[0]).String(),
			Name:       "Aryyy",
			Phone:      "08123456789",
			Image:      "https://image.com/aryyy",
			Position:   "Ngoding Slurrrr",
		},
		{
			ID:         uuid.New().String(),
			DivisionId: uuid.MustParse(uuidDivision[1]).String(),
			Name:       "Bryyy",
			Phone:      "08123456788",
			Image:      "https://image.com/bryyy",
			Position:   "Ngoding Slurrrr",
		},
		{
			ID:         uuid.New().String(),
			DivisionId: uuid.MustParse(uuidDivision[2]).String(),
			Name:       "Cryyy",
			Phone:      "08123456787",
			Image:      "https://image.com/cryyy",
			Position:   "Ngoding Slurrrr",
		},
		{
			ID:         uuid.New().String(),
			DivisionId: uuid.MustParse(uuidDivision[3]).String(),
			Name:       "Dryyy",
			Phone:      "08123456786",
			Image:      "https://image.com/dryyy",
			Position:   "Ngoding Slurrrr",
		},
		{
			ID:         uuid.New().String(),
			DivisionId: uuid.MustParse(uuidDivision[0]).String(),
			Name:       "Eryyy",
			Phone:      "08123456785",
			Image:      "https://image.com/eryyy",
			Position:   "Ngoding Slurrrr",
		},
	}

	// insert seeders
	if err := db.Create(&users).Error; err != nil {
		panic(err)
	}

	if err := db.Create(&divisions).Error; err != nil {
		panic(err)
	}

	if err := db.Create(&employees).Error; err != nil {
		panic(err)
	}

}
