package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/iMbGenom/gorilla/src/helpers"
)

var command string
var action string
var fileName = "training_data.txt"
var path = "/../functional_spec/fixtures/" + fileName
var trainDataIsExist int = 1
var emptySlot int = 0

func main() {
	/* Read File Input With Argument*/
	// count if with argument
	if len(os.Args) > 0 {
		// store argument into argue
		argue := os.Args[1]
		// check if .txt file
		checkFile, err := regexp.MatchString(".txt", argue)
		// if error checking file
		if err != nil {
			fmt.Println("Error", err)
		}
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
			data := string(fileInput)
			arrData := strings.Split(data, "\n")
			for i, v := range arrData {
				fmt.Println(i, v)
			}
		} else {
			command = os.Args[1]
			/* GET PWD */
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("ERROR getPWD", err)
			}
			fullPath := dir + path
			/* END GET PWD */

			/* CHECK IS DATA IS EXIST */
			_, err = os.Stat(fullPath)
			if err != nil {
				trainDataIsExist = 0
			}
			/* END CHECK IS DATA IS EXIST */

			/* CORE FLOW */
			if command == "create_parking_lot" && trainDataIsExist == 1 {
				fmt.Println("you must exit first")
				switch command {
				case "exit":
					fmt.Println("exitx")
					helpers.DeleteFile()
				}
			} else {
				if command == "create_parking_lot" {
					fmt.Println("create_parking_lot")
					action = os.Args[2]
					helpers.CreateFile(action)
				} else {
					// readfile and store argument into fileInput
					fileInput, err := ioutil.ReadFile(fullPath)
					// if error read file input
					if err != nil {
						fmt.Println("Can't read file:", fullPath)
						panic(err)
					}
					// fileInput is the file content
					// fmt.Println(string(fileInput))
					data := string(fileInput)
					arrData := strings.Split(data, "\n")
					// fmt.Println(len(arrData))
					for i, v := range arrData {
						fmt.Println(i, v)
						if v == "" {
							emptySlot = i
							break
						}
					}
					fmt.Println("Empty Slot = ", emptySlot)
					// return
					switch command {
					case "park":
						fmt.Println("parkx")
						for j, vv := range arrData {
							fmt.Println(j, vv)
							// if j == emptySlot {
							// 	fmt.Println(os.Args[2]+" "+os.Args[3], "rere")
							// 	// helpers.EditFile(emptySlot, os.Args[2]+" "+os.Args[3])
							// 	break
							// } else {
							helpers.EditFile(j, "1\n")
							// if j == emptySlot {
							// 	fmt.Println("Urutan ke", emptySlot)
							// 	helpers.EditFile(emptySlot, os.Args[2]+" "+os.Args[3]+"\n")
							// }
							// }
						}
						// fmt.Println(addPark)
					// case "park":
					// 	fmt.Println("parkx")
					// 	createTicket := TicketModel.CreateNew(os.Args[2])
					// 	fmt.Println(createTicket)

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
				}
			}
			/* END CORE FLOW */

			return
			checkFile, err := regexp.MatchString(".txt", fullPath)
			// if error checking file
			if err != nil {
				fmt.Println("Error", err)
			}
			fmt.Println(checkFile)
			// if file is .txt
			if checkFile {

			}
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
