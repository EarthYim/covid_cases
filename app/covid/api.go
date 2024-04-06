package covid

import (
	"encoding/json"
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

type Adapter struct {
	http.Client
	endpoint string
}

func NewApiAdapter(endpoint string) *Adapter {
	return &Adapter{
		Client: http.Client{
			Timeout: 5 * time.Second,
		},
		endpoint: endpoint,
	}
}

func (a *Adapter) GetData() ([]CovidApiData, error) {
	req, err := http.NewRequest(http.MethodGet, a.endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var respData CovidApiResp
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, err
	}

	data := respData.Data

	// fmt.Printf("data: %v", data[0]) //DEBUG

	return data, nil
}
