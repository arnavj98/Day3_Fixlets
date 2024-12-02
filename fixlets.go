package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strings"
)

var filePath = "Day-3-fixlets.csv"

func main() {
	entries, headers, err := readCSV(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. List data")
		fmt.Println("2. Query find")
		fmt.Println("3. Sort the data")
		fmt.Println("4. Add data")
		fmt.Println("5. Delete data")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			listEntries(headers, entries)
		case 2:
			queryEntries(headers, entries)
		case 3:
			sortEntries(headers, &entries)
		case 4:
			addEntry(headers, &entries)
		case 5:
			deleteEntry(headers, &entries)
		case 6:
			fmt.Println("Exiting...")
			writeCSV(filePath, headers, entries)
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func readCSV(filePath string) ([][]string, []string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	headers, err := reader.Read()
	if err != nil {
		return nil, nil, err
	}

	entries, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	return entries, headers, nil
}

func writeCSV(filePath string, headers []string, entries [][]string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write(headers)
	writer.WriteAll(entries)
}

func listEntries(headers []string, entries [][]string) {
	fmt.Println(strings.Join(headers, ", "))
	for _, entry := range entries {
		fmt.Println(strings.Join(entry, ", "))
	}
}

func queryEntries(headers []string, entries [][]string) {
	fmt.Println("Available columns:")
	for i, header := range headers {
		fmt.Printf("%d. %s\n", i+1, header)
	}
	fmt.Print("Enter column number to query: ")
	var col int
	fmt.Scan(&col)
	if col < 1 || col > len(headers) {
		fmt.Println("Invalid column number.")
		return
	}

	fmt.Print("Enter query value: ")
	var query string
	fmt.Scan(&query)

	fmt.Println(strings.Join(headers, ", "))
	for _, entry := range entries {
		if strings.Contains(entry[col-1], query) {
			fmt.Println(strings.Join(entry, ", "))
		}
	}
}

func sortEntries(headers []string, entries *[][]string) {
	fmt.Println("Available columns:")
	for i, header := range headers {
		fmt.Printf("%d. %s\n", i+1, header)
	}
	fmt.Print("Enter column number to sort by: ")
	var col int
	fmt.Scan(&col)
	if col < 1 || col > len(headers) {
		fmt.Println("Invalid column number.")
		return
	}

	sort.Slice(*entries, func(i, j int) bool {
		return (*entries)[i][col-1] < (*entries)[j][col-1]
	})

	fmt.Println("Entries sorted.")
}

func addEntry(headers []string, entries *[][]string) {
	newEntry := make([]string, len(headers))
	for i, header := range headers {
		fmt.Printf("Enter value for %s: ", header)
		fmt.Scan(&newEntry[i])
	}
	*entries = append(*entries, newEntry)
	fmt.Println("Entry added.")
}

func deleteEntry(headers []string, entries *[][]string) {
	fmt.Println("Available entries:")
	listEntries(headers, *entries)
	fmt.Print("Enter row number to delete: ")
	var row int
	fmt.Scan(&row)
	if row < 1 || row > len(*entries) {
		fmt.Println("Invalid row number.")
		return
	}

	*entries = append((*entries)[:row-1], (*entries)[row:]...)
	fmt.Println("Entry deleted.")
}
