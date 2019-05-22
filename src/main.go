package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/iMbGenom/gorilla/src/helpers"
	TicketModel "github.com/iMbGenom/gorilla/src/models/ticket"
)

var action string
var subject string

func main() {
	// fmt.Println(os.Args[1]) // get exit code

	/* Read File Input With Argument*/
	// count if with argument
	if len(os.Args) > 0 {
		// fmt.Println("Any argument")
		// store argument into argue
		argue := os.Args[1]
		// check if .txt file
		checkFile, err := regexp.MatchString(".txt", argue)
		// if error checking file
		if err != nil {
			fmt.Println("Error", err)
		}
		// fmt.Println(checkFile)
		// if file is .txt
		if checkFile {
			// readfile and store argument into fileInput
			fileInput, err := ioutil.ReadFile(os.Args[1])
			// if error read file input
			if err != nil {
				fmt.Println("Can't read file:", os.Args[1])
				panic(err)
			}
			// fileInput is the file content
			// fmt.Println(string(fileInput))
			data := string(fileInput)
			arrData := strings.Split(data, "\n")
			// fmt.Println(len(arrData))
			for i, v := range arrData {
				fmt.Println(i, v)
			}
		} else {
			// if not .txt file
			// fmt.Println("Single Command")
			// data := os.Args[1] + " " + os.Args[2] + " " + os.Args[3]
			action = os.Args[1]

			switch action {
			case "create_parking_lot":
				fmt.Println("create_parking_lot")
				helpers.CreateFile()
			case "park":
				fmt.Println("parkx")
				createTicket := TicketModel.CreateNew(os.Args[2])
				fmt.Println(createTicket)
			case "leave":
				fmt.Println("leavex")
			case "status":
				fmt.Println("statusx")
			case "registration_numbers_for_cars_with_colour":
				fmt.Println("registration_numbers_for_cars_with_colourx")
			case "slot_numbers_for_cars_with_colour":
				fmt.Println("slot_numbers_for_cars_with_colourx")
			case "slot_number_for_registration_number":
				fmt.Println("slot_number_for_registration_numberx")
			case "exit":
				fmt.Println("exitx")
				helpers.DeleteFile()
			}
			// ticket.CreateNew()
			// fmt.Println(data)
		}
	} else {
		fmt.Println(os.Args)
	}
	/* END Read File Input With Argument*/

	/* Read File Input Without Argument*/
	// fileInput, err := ioutil.ReadFile("../functional_spec/fixtures/file_input.txt") // just pass the file name
	// if err != nil {
	// 	fmt.Print("Error", err)
	// }
	// // fmt.Println(b) // print the content as 'bytes'
	// data := string(fileInput) // convert content to a 'string'
	// // fmt.Println(data) // print the content as a 'string'

	// arrData := strings.Split(data, "\n")
	// // fmt.Println(len(arrData))

	// for i, v := range arrData {
	// 	fmt.Println(i, v)
	// }
	/* Read File Input Without Argument*/
}
