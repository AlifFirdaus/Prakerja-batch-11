package routes

import (
	"game-list-api/config"
	"game-list-api/pkg/games"
	"game-list-api/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes function takes a pointer to a gin.Engine as an argument and returns a pointer to a gorm.DB object.
// It sets up the database connection, creates a new instance of the game repository, creates a new instance of the game controller,
// creates a new route group, and sets up the routes for the game controller.
// Finally, it returns the database connection object.
func SetupRoutes(router *gin.Engine) *gorm.DB {
	// Load environment variables from .env file
	config.LoadEnv()

	// Set up database connection
	db := config.SetupDatabase()

	// Create a new instance of the game repository
	GamesRepository := repositories.NewGameRepository(db)

	// Create a new instance of the game controller
	gameController := games.NewGameController(
		games.NewGameService(
			GamesRepository,
		),
	)

	// Create a new route group
	routeGroup := router.Group("/api/v1")

	// Create a new route group for games
	routerGroupGames := routeGroup.Group("/games")

	// Set up routes for the game controller
	routerGroupGames.POST("", gameController.Create)
	routerGroupGames.GET("", gameController.FindAll)
	routerGroupGames.GET("/:id", gameController.FindById)
	routerGroupGames.DELETE("/:id", gameController.Delete)

	// Return the database connection object
	return db
}
