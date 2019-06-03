package main

import "fmt"

func main() {
	mySlice := []string{"Adi", "Radili", "Sakata", "Gintoki"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	// fmt.Println(s[0])
	s[0] = "Jack"
}
