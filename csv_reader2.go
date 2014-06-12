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
	fields, err := csvReader.ReadAll()
	if err == io.EOF {
		fmt.Println("Error")
	} else if err != nil {
		panic(err)
	}
	fmt.Println(fields)
	fmt.Println(fields[0])
	fmt.Println(reflect.TypeOf(fields))
	column := []string{"A", "B", "C", "D"}
	fmt.Println(column)
	m := make(map[string][]string)
	m["primer"] = column
	fmt.Println(m)
}
