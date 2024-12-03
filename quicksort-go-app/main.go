package main

import (
	"encoding/json"
	"fmt"

	"github.com/jamie-burns0/quicksort-go/quicksort"
)

type RequestBody struct {
	Data []int `json:"data"`
}

func main() {

	unsorted := []int{5,10,1,3,2,4}
	sorted := quicksort.Sort(unsorted)
	fmt.Printf("sorted: %v\n", sorted)

	body := `{"data":[5,10,1,3,2,4]}`
	var requestBody RequestBody

	err := json.Unmarshal([]byte(body), &requestBody)

	fmt.Println(requestBody.Data)
	fmt.Println(err)

	sorted2 := quicksort.Sort(requestBody.Data)
	fmt.Printf("sorted2: %v\n", sorted2)
}
