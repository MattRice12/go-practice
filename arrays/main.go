package main

import "fmt"

func main() {
	var langsArray [3]string
	langsArray[0] = "Go"
	langsArray[1] = "Matt"
	langsArray[2] = "Go"
	fmt.Println(langsArray)

	var langsSlice []string
	langsSlice = append(langsSlice, "Go")
	langsSlice = append(langsSlice, "Matt")
	langsSlice = append(langsSlice, "Go")
	langsSlice = append(langsSlice, "!")
	fmt.Println(langsSlice)
}
