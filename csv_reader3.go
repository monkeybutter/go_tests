package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
)

type Columner interface {
	Mean() float64
	Var() float64
}

type Integer_Column struct { 
	col []int
}

func (c Integer_Column) Mean() float64 {
	sum := 0.0
	for _, el := range c.col {
		sum += float64(el)
	}
	return sum/float64(len(c.col))
}

func (c Integer_Column) Var() float64 {
	sum := 0.0
	mean := c.Mean()

	for _, el := range c.col {
		sum += (float64(el)-mean)*(float64(el)-mean)
	}
	return sum
}

type Float_Column struct { 
	col []float64
}

func (c Float_Column) Mean() float64 {
	sum := 0.0

	for _, el := range c.col {
		sum += el
	}
	return sum/float64(len(c.col))
}

func (c Float_Column) Var() float64 {
	sum := 0.0
	mean := c.Mean()

	for _, el := range c.col {
		sum += (el-mean)*(el-mean)
	}
	return sum
}

type Bool_Column struct { 
	col []bool
}

func (c Bool_Column) Mean() float64 {

	return .0
}

func (c Bool_Column) Var() float64 {

	return .0
}

type String_Column struct { 
	col []string
}

func (c String_Column) Mean() float64 {

	return .0
}

func (c String_Column) Var() float64 {

	return .0
}


func converter(slice []string) Columner {
	ints := make([]int, len(slice))
	int_col := Integer_Column{ints} 
	conv_succeed := true
	for i := range slice {
		t, err := strconv.Atoi(slice[i])
		if err == nil {
			int_col.col[i] = t
		} else {
			conv_succeed = false
			break
		}
	}
	if conv_succeed {
		return int_col
	}

	floats := make([]float64, len(slice))
	float_col := Float_Column{floats} 
	conv_succeed = true
	for i := range slice {
		t, err := strconv.ParseFloat(slice[i], 64)
		if err == nil {
			float_col.col[i] = t
		} else {
			conv_succeed = false
			break
		}
	}
	if conv_succeed {
		return float_col
	}

	bools := make([]bool, len(slice))
	bool_col := Bool_Column{bools} 
	conv_succeed = true
	for i := range slice {
		t, err := strconv.ParseBool(slice[i])
		if err == nil {
			bool_col.col[i] = t
		} else {
			conv_succeed = false
			break
		}
	}
	if conv_succeed {
		return bool_col
	}

	strings := make([]string, len(slice))
	string_col := String_Column{strings} 
	for i, el := range slice {
		string_col.col[i] = el
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
	
	m := make(map[string]Columner)

	for index,colName := range fields[0] {
		m[colName] = converter(columns[index][1:])
	}

	fmt.Println(m)
	fmt.Println(m["F"])
	fmt.Println(m["index"])
	fmt.Println(reflect.TypeOf(m["F"]))
	fmt.Println(reflect.TypeOf(m["index"]))
}

