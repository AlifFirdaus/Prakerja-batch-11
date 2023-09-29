package games

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Define an interface for the GameController
type GameControllerInterface interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

// Define the GameController struct
type GameController struct {
	service GameServiceInterface
}

// Define a function to create a new GameController
func NewGameController(service GameServiceInterface) GameControllerInterface {
	return &GameController{
		service,
	}
}

// Define a function to create a new game
func (c *GameController) Create(ctx *gin.Context) {
	// Bind the request to a CreateRequest struct
	var request CreateRequest

	if err := ctx.ShouldBind(&request); err != nil {
		// Return a 400 error if there is an error binding the request
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Call the Create function of the GameService
	result, err := c.service.Create(request)
	if err != nil {
		// Return a 500 error if there is an error creating the game
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return a 201 status code and the created game
	ctx.JSON(201, gin.H{
		"data": result,
	})
}

// Define a function to get all games
func (c *GameController) FindAll(ctx *gin.Context) {
	// Call the FindAll function of the GameService
	result, err := c.service.FindAll()
	if err != nil {
		// Return a 500 error if there is an error getting all games
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return a 200 status code and all games
	ctx.JSON(200, gin.H{
		"data": result,
	})
}

// Define a function to get a game by ID
func (c *GameController) FindById(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("id")

	// Convert the ID from a string to a uint
	idu64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err)
	}

	iduint := uint(idu64)

	// Call the FindById function of the GameService
	result, err := c.service.FindById(iduint)
	if err != nil {
		// Return a 500 error if there is an error getting the game by ID
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return a 200 status code and the game with the specified ID
	ctx.JSON(200, gin.H{
		"data": result,
	})
}

// Define a function to delete a game by ID
func (c *GameController) Delete(ctx *gin.Context) {
	// Get the ID from the request parameters
	id := ctx.Param("id")

	// Convert the ID from a string to a uint
	idu64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		panic(err)
	}

	iduint := uint(idu64)

	// Call the Delete function of the GameService
	err = c.service.Delete(iduint)
	if err != nil {
		// Return a 500 error if there is an error deleting the game by ID
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Return a 200 status code and a success message
	ctx.JSON(200, gin.H{
		"message": "Game deleted successfully",
	})
}
