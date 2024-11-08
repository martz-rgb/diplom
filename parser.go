package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func csv_analysis(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err, "failed to open file")
	}

	reader := csv.NewReader(file)

	count := 0
	k_map := map[int]int{}

	// header
	header, err := reader.Read()
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}

	index := 0
	for i := range header {
		if header[i] == "polynom" {
			index = i
			break
		}
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		count++
		k := strings.Count(record[index], "+") + 1
		k_map[k]++
	}

	fmt.Printf("filename: %s, total count: %d, k_map: %v\n", filename, count, k_map)
}
