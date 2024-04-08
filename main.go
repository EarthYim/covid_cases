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
	covidService := covid.New(apiAdapter)

	r.GET("/covid/summary", app.NewHandler(covidService.HandleRequest))
	app.Run(r, config.Server.Port)
}
