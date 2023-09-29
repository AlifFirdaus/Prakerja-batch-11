package repositories

import (
	"game-list-api/entities"

	"gorm.io/gorm"
)

// Define an interface for GameRepository with four methods
type GameRepositoryInterface interface {
	FindAll() (*[]entities.Game, error)
	FindById(id uint) (*entities.Game, error)
	Create(game entities.Game) (*entities.Game, error)
	Delete(id uint) error
}

// Define a struct for GameRepository with a single field db of type *gorm.DB
type GameRepository struct {
	db *gorm.DB
}

// Define a function NewGameRepository that returns a GameRepositoryInterface
// and takes a *gorm.DB as an argument
func NewGameRepository(db *gorm.DB) GameRepositoryInterface {
	// Return a pointer to a new GameRepository struct with the db field set to the passed in *gorm.DB
	return &GameRepository{
		db,
	}
}

// Define the FindAll method for the GameRepository struct
func (r *GameRepository) FindAll() (*[]entities.Game, error) {
	// Declare a slice of entities.Game
	var games []entities.Game

	// Query the database for all games and store the result in the games slice
	if err := r.db.Find(&games).Error; err != nil {
		return nil, err
	}

	// Return a pointer to the games slice and nil error
	return &games, nil
}

// Define the FindById method for the GameRepository struct
func (r *GameRepository) FindById(id uint) (*entities.Game, error) {
	// Declare a single entities.Game
	var game entities.Game

	// Query the database for a game with the passed in id and store the result in the game variable
	if err := r.db.First(&game, id).Error; err != nil {
		return nil, err
	}

	// Return a pointer to the game variable and nil error
	return &game, nil
}

// Define the Create method for the GameRepository struct
func (r *GameRepository) Create(game entities.Game) (*entities.Game, error) {
	// Insert the passed in game into the database
	if err := r.db.Create(&game).Error; err != nil {
		return nil, err
	}

	// Return a pointer to the game variable and nil error
	return &game, nil
}

// Define the Delete method for the GameRepository struct
func (r *GameRepository) Delete(id uint) error {
	// Delete the game with the passed in id from the database
	if err := r.db.Delete(&entities.Game{}, id).Error; err != nil {
		return err
	}

	// Return nil error
	return nil
}
