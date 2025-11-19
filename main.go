package main

import (
	"social-backend/cmd"
	"social-backend/config"
	"social-backend/conn"
)

func main() {
	config.LoadConfig()
	conn.ConnectDB()

	cmd.ServeCmd.Execute()

}
