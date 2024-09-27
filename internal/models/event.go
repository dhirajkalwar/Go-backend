package models

import (
	"gorm.io/gorm"
)

// Event represents an event in the system.
type Event struct {
	gorm.Model           // Embeds ID, CreatedAt, UpdatedAt, DeletedAt
	Title       string   `gorm:"not null"`           // Title of the event
	Description string   `gorm:"type:text"`          // Detailed description of the event
	StartTime   string   `gorm:"not null"`           // Start time of the event
	EndTime     string   `gorm:"not null"`           // End time of the event
	Tickets     []Ticket `gorm:"foreignKey:EventID"` // One-to-many relationship with Ticket
}
