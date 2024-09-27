package models

import (
	"gorm.io/gorm"
)

// Ticket represents a ticket for an event purchased by a user.
type Ticket struct {
	gorm.Model       // Embeds ID, CreatedAt, UpdatedAt, DeletedAt
	UserID     uint  `gorm:"not null"` // Foreign key to User
	EventID    uint  `gorm:"not null"` // Foreign key to Event
	User       User  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Event      Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
