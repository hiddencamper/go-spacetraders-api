package main

import (
	"fmt"

	APITools "https://github.com/hiddencamper/go-spacetraders-api"
)

func main() {
	var url = "http://spacetraders.io/api/v2/"

	data, err := APITools.GetRequest(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(string(data))
}
