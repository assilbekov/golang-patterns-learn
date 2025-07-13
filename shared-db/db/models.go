package db

import (
	"time"

	"gorm.io/gorm"
)

// User represents the user information
type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:100;not null"`
	Surname   string         `gorm:"size:100;not null"`
	Grade      string         `gorm:"size:50;not null"`
	RatePerHour float64       `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// UserWorkHours represents the user's working hours per day
type UserWorkHours struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `gorm:"not null;index"`
	User        User           `gorm:"foreignKey:UserID"`
	HoursPerDay float64        `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
