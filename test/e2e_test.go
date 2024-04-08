package test

import (
	"covid_cases/app/covid"
	"encoding/json"
	"net/http"
	"testing"
)

func TestE2E(t *testing.T) {

	t.Run("TestE2E", func(t *testing.T) {
		resp, err := http.Get("http://localhost:8080/covid/summary")
		if err != nil {
			t.Fatalf("failed calling API: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("http status code is not OK: %v", resp.StatusCode)
		}

		var data covid.CovidResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if data.Age.Age0To30 == 0 {
			t.Log("Age0To30 is 0")
		}
		if data.Province["Bangkok"] == 0 {
			t.Log("Bangkok is 0")
		}
	})
}
