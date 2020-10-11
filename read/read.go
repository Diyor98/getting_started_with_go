package main

import (
	"bufio" // for file read
	f "fmt"
	"log"     // for console display??
	"os"      // for file access
	"strings" // stings manipulation
)

//Person struct
type Person struct {
	fname [20]byte
	lname [20]byte
}

func main() {
	var fileName string
	mySlice := make([]Person, 0)
	f.Println("Enter file name:")
	f.Scanln(&fileName)
	myfile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer myfile.Close() // why to use this?

	// Try with ReadLine
	// reader := bufio.NewReader(myfile)
	// for reader.ReadLine() {

	// }

	scanner := bufio.NewScanner(myfile)
	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Fields(text) // splits text by white space
		fname := words[0]
		lname := words[1]
		tempPerson := Person{}
		copy(tempPerson.fname[:], fname)
		copy(tempPerson.lname[:], lname)
		mySlice = append(mySlice, tempPerson)
	}

	for _, obj := range mySlice {
		fname := string(obj.fname[:])
		lname := string(obj.lname[:])
		f.Println(fname + " " + lname + "\n")

	}
}
