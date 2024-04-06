package main

import (
	"covid_cases/app"
	"covid_cases/app/covid"
)

func main() {
	r := app.NewRouter()

	apiAdapter := covid.NewApiAdapter("https://static.wongnai.com/devinterview/covid-cases.json")
	service := covid.New(apiAdapter)

	r.GET("/covid", app.NewHandler(service.HandleRequest))
	//TODO: gracefully shutdown
	r.Run()
}
