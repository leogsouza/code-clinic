package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("../../Environmental_Data_Deep_Moor_2015.txt")

	if err != nil {
		panic(err)
	}
	// defer means that Close() will be called in the end of function
	defer f.Close()

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("Total Records:", len(rows)-1)
	fmt.Println("Mean Air Temp", mean(rows, 1), median(rows, 1))
	fmt.Println("Mean Barometric", mean(rows, 2), median(rows, 2))
	fmt.Println("Mean Wind Speed", mean(rows, 7), median(rows, 7))
}

func mean(rows [][]string, idx int) float64 {
	var total float64

	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			total += val
		}
	}
	return total / float64(len(rows)-1)
}

func median(rows [][]string, idx int) float64 {
	// to hold data which will be sorted
	var sorted []float64

	// populate the sorted slice
	for i, row := range rows {
		if i != 0 {
			val, _ := strconv.ParseFloat(row[idx], 64)
			sorted = append(sorted, val)
		}
	}

	sort.Float64s(sorted)
	if len(sorted)%2 == 0 {
		// even number of items
		// for example 3, 5, 8, 9
		// median is (5+8) / 2
		middle := len(sorted) / 2
		higher := sorted[middle]
		lower := sorted[middle-1]
		return (higher + lower) / 2
	}
	// odd number of items
	// for example 3, 5, 8
	// median is 5
	middle := len(sorted) / 2
	return sorted[middle]
}
