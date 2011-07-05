package main

import (
	"io/ioutil"
	"strings"
	"sort"
)

func main() {
	fileBuf, err := ioutil.ReadFile("022_input.txt")
	if err != nil {
		panic(err.String())
	}
	fileStr := string(fileBuf)
	arr := strings.Split(fileStr, ",", -1)
	for i := range arr {
		arr[i] = strings.Split(arr[i], "\"", -1)[1]
	}

	sort.SortStrings(arr)

	result := uint64(0)
	for i := range arr {
		result += uint64(i+1) * uint64(alphabeticValue(arr[i]))
	}
	println(result)
}

func alphabeticValue(in string) int {
	sum := 0
	for i := range in {
		sum += int(in[i]) - int('A') + 1
	}
	return sum
}

