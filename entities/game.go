package entities

import "time"

// Define a struct named Game
type Game struct {
	ID          uint      `gorm:"primaryKey" json:"id"`           // Define a field named ID of type uint with gorm and json tags
	Name        string    `gorm:"type:varchar(255)" json:"name"`  // Define a field named Name of type string with gorm and json tags
	Genre       string    `gorm:"type:varchar(255)" json:"genre"` // Define a field named Genre of type string with gorm and json tags
	ReleaseDate time.Time `gorm:"type:date" json:"release_date"`  // Define a field named ReleaseDate of type time.Time with gorm and json tags
}
