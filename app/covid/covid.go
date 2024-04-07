package covid

import (
	"covid_cases/app"
	"encoding/json"
)

type CovidResponse struct {
	Provinces string     `json:"Provinces"`
	Age       CovidByAge `json:"AgeGroups"`
}

type CovidByProvince struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
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

func (s *service) CountAgeGroup() (*CovidResponse, error) {
	// 1. prepare data
	data, err := s.GetData()
	if err != nil {
		return nil, err
	}

	ageGroup := CovidByAge{}

	provMap := make(map[string]int)

	//2. count data age groups
	for _, d := range data {
		//count age group
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

		//count province
		provMap[d.ProvinceEn]++
	}

	provJsonResp, err := json.Marshal(provMap)
	if err != nil {
		return nil, err
	}
	provJsonString := string(provJsonResp)

	resp := CovidResponse{
		Provinces: provJsonString,
		Age:       ageGroup,
	}

	//3. response
	return &resp, nil
}

func (s *service) HandleRequest(ctx app.Context) {
	ageGroup, err := s.CountAgeGroup()
	if err != nil {
		ctx.InternalServerError(err)
		return
	}

	ctx.OK(ageGroup)
}
