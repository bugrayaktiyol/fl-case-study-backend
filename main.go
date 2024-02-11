package main

import (
	"log"

	"github.com/bugrayaktiyol/fl-case-study-backend/internal/users"
	"github.com/bugrayaktiyol/fl-case-study-backend/internal/utils"
)

func main() {

	// Setting up environment (dotenv load, gorm open etc)
	db, err := utils.SetupEnvironment()
	if err != nil {
		log.Fatal("Error setting up environment:", err)
	}

	utils.AutoMigrateModels(db)

	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userHandler := users.NewHandler(userService)

	app := utils.SetupFiberApp(userHandler)

	app.Listen(":3001")
}
