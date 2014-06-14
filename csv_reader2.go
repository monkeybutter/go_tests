package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
)

func converter(slice []string) interface{} {
	int_col := make([]int, len(slice))
	conv_succeed := true
	for i := range slice {
		t, err := strconv.Atoi(slice[i])
		if err == nil {
			int_col[i] = t
		} else {
			conv_succeed = false
			break
		}
	}
	if conv_succeed {
		return int_col
	}

	float_col := make([]float64, len(slice))
	conv_succeed = true
	for i := range slice {
		t, err := strconv.ParseFloat(slice[i], 64)
		if err == nil {
			float_col[i] = t
		} else {
			conv_succeed = false
			break
		}
	}
	if conv_succeed {
		return float_col
	}

	bool_col := make([]bool, len(slice))
	conv_succeed = true
	for i := range slice {
		t, err := strconv.ParseBool(slice[i])
		if err == nil {
			bool_col[i] = t
		} else {
			conv_succeed = false
			break
		}
	}
	if conv_succeed {
		return bool_col
	}

	return slice

}

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
	
	
	m := make(map[string]interface{})

	for index,colName := range fields[0] {
		m[colName] = converter(columns[index])
	}

	fmt.Println(m)
	//fmt.Println(reflect.TypeOf(m["primer"]))
	//fmt.Println(reflect.TypeOf(m["sencod"]))
}
