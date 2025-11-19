package cmd

import (
	"social-backend/config"
	"social-backend/conn"
	"social-backend/controller"
	"social-backend/repository/db"
	"social-backend/router"
	"social-backend/server"
	"social-backend/services"

	"github.com/spf13/cobra"

	"github.com/labstack/echo/v4"
)

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	config.LoadConfig()
	dbClient:= conn.Db()
	dbRepo := db.NewRepository(dbClient)
	usrSvc := services.ForNowUserService(*dbRepo)

	userController := controller.ForNowUserController(usrSvc)
	var echo_= echo.New()
	routes := router.NewRouter(echo_, userController)
	server := server.NewServer(echo_)

	routes.SetupRoutes()
	server.Start()
}