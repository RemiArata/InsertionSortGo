package main

import (
	"InsertionSortGo/helpers"
	"InsertionSortGo/is"
	"fmt"
)

func main() {
	data := helpers.GenerateData()
	fmt.Println(data)

	sortedData := is.InsertionSort(data)
	fmt.Println(sortedData)
}
