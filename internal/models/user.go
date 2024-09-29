package models

import (
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	gorm.Model          // Embeds ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string   `gorm:"not null"`             //  name
	Password   string   `gorm:"not null"`             // Hashed password
	Email      string   `gorm:"uniqueIndex;not null"` // Unique email address
	Tickets    []Ticket `gorm:"foreignKey:UserID"`    // One-to-many relationship with Ticket
}
