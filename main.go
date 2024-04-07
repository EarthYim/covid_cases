package main

import (
	"covid_cases/app"
	"covid_cases/app/covid"
	"covid_cases/config"
)

func main() {
	config := config.GetConfig()
	r := app.NewRouter()

	apiAdapter := covid.NewApiAdapter(config.CovidApiEndpoint, config.Http)
	service := covid.New(apiAdapter)

	r.GET("/covid", app.NewHandler(service.HandleRequest))
	//TODO: gracefully shutdown
	app.Run(r, config.Server.Port)
}
