package request

import (
	"encoding/json"
	"io/ioutil"
	"kmid_checker/pkg/env"
	"net/http"
)

type ResponseData struct {
	StatusText string `json:"StatusText"`
}

func GetStatusFiveYears(code string) (string, error) {
	kdMidURL := env.Get("KD_MID_URL")

	url := kdMidURL + code
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var responseData ResponseData
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return "", err
	}

	return responseData.StatusText, nil
}
