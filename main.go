package main

import (
	"game-list-api/config"
	"game-list-api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

// Define the main function, which is the entry point of the program
func main() {
	// Create a new Gin router with default middleware
	router := gin.Default()

	// Call the SetupRoutes function from the routes package, passing in the router
	// This function sets up all the routes for the API and returns a database connection
	db := routes.SetupRoutes(router)

	// Defer the CloseDatabaseConnection function from the config package, passing in the database connection
	// This function will be called when the main function exits, ensuring that the database connection is closed properly
	defer config.CloseDatabaseConnection(db)

	// Start the router and listen for incoming requests on the port specified in the APP_PORT environment variable
	router.Run(":" + os.Getenv("APP_PORT"))
}
