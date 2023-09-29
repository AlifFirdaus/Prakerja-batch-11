package games

import (
	"errors"
	"game-list-api/entities"
	"game-list-api/repositories"
	"time"

	"gorm.io/gorm"
)

// Define an interface for the GameService
type GameServiceInterface interface {
	Create(CreateRequest) (*entities.Game, error)
	FindAll() (*[]entities.Game, error)
	FindById(id uint) (*entities.Game, error)
	Delete(id uint) error
}

// Define the GameService struct
type GameService struct {
	gameRepository repositories.GameRepositoryInterface
}

// Define a constructor function for the GameService
func NewGameService(
	gameRepository repositories.GameRepositoryInterface,
) GameServiceInterface {
	return &GameService{
		gameRepository,
	}
}

// Define the Create method for the GameService
func (s *GameService) Create(request CreateRequest) (*entities.Game, error) {
	// Parse the release date from the request
	releaseDate, err := time.Parse("2006-01-02", request.ReleaseDate)
	if err != nil {
		panic(err)
	}

	// Create a new game entity using the game repository
	game, err := s.gameRepository.Create(entities.Game{
		Name:        request.Name,
		Genre:       request.Genre,
		ReleaseDate: releaseDate,
	})
	if err != nil {
		panic(err)
	}

	return game, nil
}

// Define the FindAll method for the GameService
func (s *GameService) FindAll() (*[]entities.Game, error) {
	// Find all games using the game repository
	games, err := s.gameRepository.FindAll()
	if err != nil {
		panic(err)
	}

	return games, nil
}

// Define the FindById method for the GameService
func (s *GameService) FindById(id uint) (*entities.Game, error) {
	// Find a game by ID using the game repository
	game, err := s.gameRepository.FindById(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}

	// If the game is not found, return an error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("game not found")
	}

	return game, nil
}

// Define the Delete method for the GameService
func (s *GameService) Delete(id uint) error {
	// Find a game by ID using the game repository
	_, err := s.gameRepository.FindById(id)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		panic(err)
	}

	// If the game is not found, return an error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("game not found")
	}

	// Delete the game using the game repository
	return s.gameRepository.Delete(id)
}
