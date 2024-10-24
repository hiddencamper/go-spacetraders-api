package APITools

import (
	"io"
	"net/http"
	"strings"
)

func GetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
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
	reader := strings.NewReader(data)
	resp, err := http.Post(url, "application/json", reader)
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
