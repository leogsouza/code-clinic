package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("../Environmental_Data_Deep_Moor_2015.txt")

	if err != nil {
		panic(err)
	}
	// defer means that Close() will be called in the end of function
	defer f.Close()

	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		fmt.Println(row)
	}
}
