package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type Column struct {
	rows []interface{}
}

func (col Column) GetSize() int {
	return len(col.rows)
}

func converter(slice []string) []interface{} {
	int_col := make([]interface{}, len(slice))
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

	float_col := make([]interface{}, len(slice))
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

	bool_col := make([]interface{}, len(slice))
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

	string_col := make([]interface{}, len(slice))
	for i, el := range slice {
		string_col[i] = el
	}
	return string_col

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
	
	m := make(map[string]Column)

	for index,colName := range fields[0] {
		m[colName] = Column{}
		fmt.Println(reflect.TypeOf(m[colName].rows))
		fmt.Println(reflect.TypeOf(converter(columns[index][1:])))
		//m[colName].rows = converter(columns[index])
	}

	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m["F"]))
	fmt.Println(reflect.TypeOf(m["index"]))
}
