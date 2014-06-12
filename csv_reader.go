package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	csvFile, err := os.Open("./table.csv")
	defer csvFile.Close()
	if err != nil {
		panic(err)
	}
	csvReader := csv.NewReader(csvFile)
	for {
		fields, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(fields)
		fmt.Println(reflect.TypeOf(fields))
	}
}
