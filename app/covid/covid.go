package covid

import (
	"covid_cases/app"

	"github.com/gin-gonic/gin"
)

type CovidResponse struct {
	Provinces []CovidByProvince `json:"Provinces"`
	Age       []CovidByAge      `json:"AgeGroups"`
}

type CovidByProvince struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type CovidByAge struct {
	Group string `json:"group"`
	Count int    `json:"count"`
}

func HandleCovid(c app.Context) {
	c.OK("Hello")
}

func New() gin.HandlerFunc {
	return app.NewHandler(HandleCovid)
}
