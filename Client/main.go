package main

import (
	"fmt"

	APITools "github.com/hiddencamper/go-spacetraders-api/APITools"
)

func main() {
	var url = "http://spacetraders.io/api/v2/"

	data, err := APITools.GetRequest(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(string(data))
}
