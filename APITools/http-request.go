package APITools

import (
	"io"
	"net/http"
	"strings"
)

var auth string = ""

func SetAuthCode(code string) {
	auth = code
}

func GetRequest(url string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if auth != "" {
		req.Header.Add("Authorization", "Bearer "+auth)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func PostRequest(url string, data string) ([]byte, error) {

	client := &http.Client{}
	reader := strings.NewReader(data)
	req, err := http.NewRequest(http.MethodPost, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	if auth != "" {
		req.Header.Add("Authorization", "Bearer "+auth)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
