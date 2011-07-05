package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBuf, err := ioutil.ReadFile("013_input.txt")
	if err != nil {
		panic(err.String())
	}
	fileStr := string(fileBuf)
	lines := strings.Split(fileStr, "\n", 100)
	result := ""
	carryout, digit := 0, 0
	for i := 49; i >= 0; i-- {
		sum := carryout
		for j := range lines {
			num, _ := strconv.Atoi(string(lines[j][i]))
			sum += num
		}
		digit = sum % 10
		carryout = sum / 10
		result = strconv.Itoa(digit) + result
	}
	result = strconv.Itoa(carryout) + result
	println(result[0:10])
}

