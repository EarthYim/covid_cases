package main

import (
	"covid_cases/app"
	"covid_cases/app/covid"
)

func main() {
	r := app.NewRouter()

	r.GET("/covid", covid.New())
	//TODO: gracefully shutdown
	r.Run()
}
