package config

import (
	"game-list-api/entities"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// SetupDatabase function returns a pointer to a gorm.DB instance
func SetupDatabase() *gorm.DB {
	// switch statement to determine which database driver to use based on the value of the DB_CONNECTION environment variable
	switch os.Getenv("DB_CONNECTION") {
	case "mysql":
		return setupMysql() // call setupMysql function if DB_CONNECTION is set to "mysql"
	case "postgres":
		return setupPostgres() // call setupPostgres function if DB_CONNECTION is set to "postgres"
	default:
		panic("DB_CONNECTION is not defined") // panic if DB_CONNECTION is not defined or set to an unsupported value
	}
}

// setupMysql function returns a pointer to a gorm.DB instance configured to use MySQL
func setupMysql() *gorm.DB {
	// construct the DSN (Data Source Name) string to connect to the MySQL database using environment variables
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"

	// open a connection to the MySQL database using the gorm.Open function and the mysql driver
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database") // panic if there is an error connecting to the database
	}

	// automatically migrate the Game entity to the database
	db.AutoMigrate(&entities.Game{})

	return db // return the gorm.DB instance
}

// setupPostgres function returns a pointer to a gorm.DB instance configured to use PostgreSQL
func setupPostgres() *gorm.DB {
	// construct the DSN (Data Source Name) string to connect to the PostgreSQL database using environment variables
	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USERNAME") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_DATABASE") + " port=" + os.Getenv("DB_PORT") + " sslmode=disable TimeZone=Asia/Shanghai"

	// open a connection to the PostgreSQL database using the gorm.Open function and the postgres driver
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database") // panic if there is an error connecting to the database
	}

	// automatically migrate the Game entity to the database
	db.AutoMigrate(&entities.Game{})

	return db // return the gorm.DB instance
}

// CloseDatabaseConnection function closes the database connection
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("failed to close database connection") // panic if there is an error closing the database connection
	}

	dbSQL.Close() // close the database connection
}
