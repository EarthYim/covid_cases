package covid_test

import (
	"covid_cases/app/covid"
	"errors"
	"testing"
)

type mockAdapter struct {
	err bool
}

func (m *mockAdapter) GetData() ([]covid.CovidApiData, error) {
	if m.err {
		return nil, errors.New("error")
	} else {
		return []covid.CovidApiData{
			{
				Age:        0,
				ProvinceEn: "",
			},
			{
				Age:        25,
				ProvinceEn: "Bangkok",
			},
			{
				Age:        35,
				ProvinceEn: "Bangkok",
			},
			{
				Age:        45,
				ProvinceEn: "Bangkok",
			},
			{
				Age:        55,
				ProvinceEn: "Bangkok",
			},
			{
				Age:        65,
				ProvinceEn: "Bangkok",
			},
			{
				Age:        75,
				ProvinceEn: "Bangkok",
			},
			{
				Age:        85,
				ProvinceEn: "Bangkok",
			},
			{
				Age:        95,
				ProvinceEn: "Bangkok",
			},
			{
				Age:        105,
				ProvinceEn: "Bangkok",
			},
		}, nil
	}
}

type mockContext struct {
	*testing.T
	flagInternalError bool
	flagOK            bool
}

func (m mockContext) InternalServerError(err error) {
	if !m.flagInternalError {
		m.Errorf("expected internal server error")
	}
}

func (m mockContext) OK(data interface{}) {
	if !m.flagOK {
		m.Errorf("expected ok")
	}
}

func TestHandleCovid(t *testing.T) {

	t.Run("Test HandleCovid with error", func(t *testing.T) {
		s := covid.New(&mockAdapter{err: true})
		s.HandleRequest(mockContext{flagInternalError: true, flagOK: false})
	})

	t.Run("Test HandleCovid with success", func(t *testing.T) {
		s := covid.New(&mockAdapter{err: false})
		s.HandleRequest(mockContext{flagInternalError: false, flagOK: true})
	})
}
