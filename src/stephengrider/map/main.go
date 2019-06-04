package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// var colors map[string]string

	// colors := make(map[string]string) // MAKE is built-in function .. key string, value string
	// colorsx := make(map[int]string)   // MAKE is built-in function .. key int, value string

	// colors["white"] = "#ffffff"
	// colorsx[10] = "#ffffff"

	// delete(colorsx, 10) // DELETE is built-in function for delete element map

	// fmt.Println(colors)
	// fmt.Println(colorsx)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
