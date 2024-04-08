package covid

import (
	"covid_cases/config"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

type CovidApiResp struct {
	Data []CovidApiData `json:"Data"`
}

type CovidApiData struct {
	ConfirmDate    string `json:"ConfirmDate"`
	No             int    `json:"No"`
	Age            int    `json:"Age"`
	Gender         string `json:"Gender"`
	GenderEn       string `json:"GenderEn"`
	Nation         string `json:"Nation"`
	NationEn       string `json:"NationEn"`
	Province       string `json:"Province"`
	ProvinceId     int    `json:"ProvinceId"`
	District       string `json:"District"`
	ProvinceEn     string `json:"ProvinceEn"`
	StatQuarantine int    `json:"StatQuarantine"`
}

type adapter struct {
	http.Client
	endpoint string
}

func NewApiAdapter(endpoint string, config config.Http) *adapter {
	return &adapter{
		Client: http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
		endpoint: endpoint,
	}
}

func (a *adapter) GetData() ([]CovidApiData, error) {
	req, err := http.NewRequest(http.MethodGet, a.endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK || resp.Body == nil {
		return nil, errors.New("failed to get data")
	}

	var respData CovidApiResp
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	data := respData.Data
	return data, nil
}
