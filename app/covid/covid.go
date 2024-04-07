package covid

import (
	"covid_cases/app"
)

type CovidResponse struct {
	Province map[string]int `json:"Provinces"`
	Age      CovidByAge     `json:"AgeGroups"`
}

type CovidByAge struct {
	Age0To30   int `json:"0-30"`
	Age31To60  int `json:"31-60"`
	Age61Plus  int `json:"61+"`
	AgeUnknown int `json:"N/A"`
}

type service struct {
	apiAdapter
}

type apiAdapter interface {
	GetData() ([]CovidApiData, error)
}

func New(a apiAdapter) *service {
	return &service{
		apiAdapter: a,
	}
}

func (s *service) HandleRequest(ctx app.Context) {
	data, err := s.GetData()
	if err != nil {
		ctx.InternalServerError(err)
		return
	}

	ageGroup := s.CountAgeGroup(data)
	if err != nil {
		ctx.InternalServerError(err)
		return
	}

	ctx.OK(ageGroup)
}

func (s *service) CountAgeGroup(data []CovidApiData) *CovidResponse {
	ageGroup := CovidByAge{}
	provMap := make(map[string]int)

	for _, d := range data {
		switch {
		case d.Age == 0:
			ageGroup.AgeUnknown++
		case d.Age > 0 && d.Age <= 30:
			ageGroup.Age0To30++
		case d.Age > 30 && d.Age <= 60:
			ageGroup.Age31To60++
		case d.Age > 60:
			ageGroup.Age61Plus++
		}

		if d.ProvinceEn == "" {
			provMap["N/A"]++
		} else {
			provMap[d.ProvinceEn]++
		}
	}

	var lenMap int
	for _, v := range provMap {
		lenMap += v
	}

	resp := CovidResponse{
		Province: provMap,
		Age:      ageGroup,
	}

	return &resp
}
