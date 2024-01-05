package main

import (
	"log"

	"github.com/bugrayaktiyol/fl-case-study-backend/api/users"
)

func main() {
	// Setting up environment (dotenv load, gorm open etc)
	db, err := SetupEnvironment()
	if err != nil {
		log.Fatal("Error setting up environment:", err)
	}

	AutoMigrateModels(db)

	userRepository := users.NewUserRepository(db)

	app := SetupFiberApp(userRepository)

	app.Listen(":3001")
}
