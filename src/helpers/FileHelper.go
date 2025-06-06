package helpers

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// var path = "/Users/radily/www/goradili/src/github.com/iMbGenom/gorilla/functional_spec/fixtures/file_input_test.txt"
var fileName = "training_data.txt"
var path = "/../functional_spec/fixtures/" + fileName

// func CreateFile() {
// 	// detect if file exists
// 	var _, err = os.Stat(path)

// 	// create file if not exists
// 	if os.IsNotExist(err) {
// 		var file, err = os.Create(path)
// 		if isError(err) { return }
// 		defer file.Close()
// 	}

// 	fmt.Println("==> done creating file", path)
// }

func CreateFile(param string) {
	fmt.Println("Creating file", param)
	// return
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullPath := dir + path

	// detect if file exists
	_, err = os.Stat(fullPath)
	if err != nil {
		fmt.Println("Error osstat", err)
	}

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(fullPath)
		if err != nil {
			fmt.Println("Error create", err)
		}
		defer file.Close()
	}
	fmt.Println("==> done creating file", path)
	writeFile(param)
}

func writeFile(param string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullPath := dir + path

	// open file using READ & WRITE permission
	file, err := os.OpenFile(fullPath, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	count, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("Err strconv", err)
	}
	fmt.Println(count)
	i := 1
	for {
		i++
		_, err = file.WriteString("\n")
		if err != nil {
			return
		}
		if i == count {
			break
		}
	}

	// save changes
	err = file.Sync()
	if err != nil {
		return
	}

	fmt.Println("==> done writing to file")
}

func EditFile(count int, vehicle string) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fullPath := dir + path

	// open file using READ & WRITE permission
	file, err := os.OpenFile(fullPath, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	// count, err := strconv.Atoi(number)
	// if err != nil {
	// 	fmt.Println("Err strconv", err)
	// }
	// fmt.Println(count)
	_, err = file.WriteString(vehicle)
	if err != nil {
		fmt.Println("Error edit ", err)
	}

	// save changes
	// err = file.Sync()
	// if err != nil {
	// 	return
	// }

	fmt.Println("==> done writing to file")
}

func DeleteFile() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fullPath := dir + path
	var rmv = os.Remove(fullPath)
	if rmv != nil {
		fmt.Println("error", rmv)
	}
	fmt.Println("==> done deleting file")
}

// func CurrentPwd() {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		fmt.Println("Error pwd", err.Error())
// 	}
// 	fmt.Println(dir)
// 	// return ("")
// 	// return (dir + path)
// }
