package main

import (
	"fmt"

	APITools "github.com/hiddencamper/go-spacetraders-api/APITools"
)

func main() {

	status, err := APITools.API_GetStatus()
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	fmt.Println(status)

	fmt.Println(status.ResetDate)
	fmt.Println(status.Description)
	fmt.Println(status.Status)
}
