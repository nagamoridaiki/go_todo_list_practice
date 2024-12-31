package main

import (
	"log"
	"myapp/controllers"
	"myapp/db"
	"myapp/repositories"
	"myapp/router"
	"myapp/services"
	"myapp/utils/logic"
	"myapp/utils/validation"
)

func main() {
	// DB connection
	dbCon := db.Init()

	// logic layer
	//userLogic := logic.NewUserLogic()
	todoLogic := logic.NewTodoLogic()
	responseLogic := logic.NewResponseLogic()

	// validation
	todoValidate := validation.NewTodoValidation()

	// repository layer
	//userRepo := repositories.NewUserRepository(dbCon)
	todoRepo := repositories.NewTodoRepository(dbCon)
	// service layer
	todoService := services.NewTodoService(todoRepo, todoLogic, responseLogic, todoValidate)
	// controller layer
	appController := controllers.NewAppController()
	//authController := controllers.NewAuthController(authService)
	todoController := controllers.NewTodoController(todoService)

	// router configuration
	appRouter := router.NewAppRouter(appController)
	//authRouter := router.NewAuthRouter(authController)
	todoRouter := router.NewTodoRouter(todoController)
	mainRouter := router.NewMainRouter(appRouter, todoRouter)

	// API startup
	if err := mainRouter.StartWebServer(); err != nil {
		log.Printf("error start web server")
	}
}
