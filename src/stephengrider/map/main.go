package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// var colors map[string]string

	colors := make(map[string]string) // MAKE is built-in function
	colorsx := make(map[int]string)   // MAKE is built-in function

	colors["white"] = "#ffffff"
	colorsx[10] = "#ffffff"

	delete(colorsx, 10) // DELETE is built-in function for delete element map

	fmt.Println(colors)
	fmt.Println(colorsx)
}
