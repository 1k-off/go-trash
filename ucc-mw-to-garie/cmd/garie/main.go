package main

import (
	"ucc-mw-to-garie/internal/app"
)

func main() {
	env, err := app.Configure()
	app.HandleError(err)
	s, checkId := app.StartProcessing(env)
	app.GenerateGarieConfig("0 1 * * *", env.GarieConfigPath, s)
	app.StopProcessing(env, checkId)
}