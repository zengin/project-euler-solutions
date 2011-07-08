package main

import (
	"io/ioutil"
	"strings"
)

var triangles []int

func main() {
	//read names
	fileBuf, err := ioutil.ReadFile("042_input.txt")
	if err != nil {
		panic(err.String())
	}
	fileStr := string(fileBuf)
	arr := strings.Split(fileStr, ",", -1)
	for i := range arr {
		arr[i] = strings.Split(arr[i], "\"", -1)[1]
	}

	//create an array of triangle numbers
	triangles = make([]int, 20)
	for i := range triangles {
		triangles[i] = (i * (i + 1)) / 2
	}

	//iterate over names to check
	result := 0
	for i := range arr {
		if isTriangle(alphabeticValue(arr[i])) {
			result++
		}
	}
	println(result)
}

func isTriangle(in int) bool {
	for i := range triangles {
		if triangles[i] == in {
			return true
		}
	}
	return false
}

func alphabeticValue(in string) int {
	sum := 0
	for i := range in {
		sum += int(in[i]) - int('A') + 1
	}
	return sum
}

