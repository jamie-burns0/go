package main

import (
	"encoding/json"
	"fmt"

	"github.com/jamie-burns0/quicksort-go/quicksort"
)

type RequestJson struct {
	UnsortedData []int `json:"data"`
}

type ResponseJson struct {
	SortedData []int `json:"data"`
}

func main() {

	unsorted := []int{5,10,1,3,2,4}
	sorted := quicksort.Sort(unsorted)
	fmt.Printf("sorted: %v\n", sorted)

	fmt.Println("Sorting with JSON data...")

	requestBody := `{"data":[5,10,1,3,2,4]}`
	fmt.Printf("requestBody: %v\n", requestBody)

	var requestJson RequestJson

	err := json.Unmarshal([]byte(requestBody), &requestJson)
	fmt.Printf("unmarshal err: %v\n", err)

	sorted2 := quicksort.Sort(requestJson.UnsortedData)
	fmt.Printf("sorted2: %v\n", sorted2)

	responseJson := ResponseJson{
		SortedData: sorted2,
	}

	jsonBytes, err := json.Marshal(responseJson)
	fmt.Printf("marshal err: %v\n", err)

	responseBody := string(jsonBytes)

	fmt.Printf("responseBody: %v\n", responseBody)
	fmt.Println(err)
}
