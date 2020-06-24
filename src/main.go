package main

import (
	"limakcv/src/app"
	"limakcv/src/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":8116")

}
