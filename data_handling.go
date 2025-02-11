package main

import (
	"fmt"
	"hello_there/repo"
	"strconv"
	"strings"
)

// ----------------------------------------------------------------------------
// reads the person data from a local CSV file and adds it to the repo.
// this emulates reading from a database, just with less overhead for the demo.
//
// caution: for brevity, this sample contains no real error handling
func loadDataFromCSV(file string, dataRepo *repo.DataRepo) {
	lines := strings.Split(file, "\n")

	for lineNumber := 1; lineNumber < len(lines); lineNumber++ {
		if len(lines[lineNumber]) > 10 {
			parts := strings.Split(lines[lineNumber], ",")

			age, err := strconv.Atoi(strings.Trim(parts[2], "\"\r"))
			if err != nil {
				fmt.Println("Error converting age string to a number", err)
			}
			person := repo.NewPerson(parts[0], parts[1], age)

			dataRepo.People = append(dataRepo.People, *person)
		}
	}
}

// ----------------------------------------------------------------------------
type PeopleCountResult struct {
	Count uint `json:"count"`
}
