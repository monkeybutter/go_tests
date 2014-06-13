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

	/*
	Transpose the 2d string array
	*/
	columns := make([][]string, len(fields[0]))
	for i:= range(fields[0]) {
		columns[i] = make([]string, len(fields))
		for j := range fields {
			columns[i][j] = fields[j][i]
		}
	}

	fmt.Println(columns)
	fmt.Println(columns[0])
	fmt.Println(reflect.TypeOf(columns))
	
	
	colNames := fields[0]
	m := make(map[string]interface{})
	//ia := []interface{}{1,"df",3}
	ia := [][]int{{1,2,3},{1,2,3},{1,2,3},{1,2,3}}
	m["primer"] = colNames
	m["sencod"] = ia
	//columns[0]
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m["primer"]))
	fmt.Println(reflect.TypeOf(m["sencod"]))
}
