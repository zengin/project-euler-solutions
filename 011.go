package main

import (
	"strconv"
	"io/ioutil"
	"strings"
)

func main() {
	fileBuf, err := ioutil.ReadFile("011_input.txt")
	if err != nil {
		panic(err.String())
	}
	fileStr := string(fileBuf)
	lines := strings.Split(fileStr, "\n", 20)
	arr := make([][]uint, 20)
	var strArr []string
	for i := range lines {
		arr[i] = make([]uint, 20)
		strArr = strings.Split(lines[i], " ", 20)
		for j := range arr[i] {
			intvalue, _ := strconv.Atoi(strArr[j])
			arr[i][j] = uint(intvalue)
		}
	}
	max := uint(0)
	//left to right
	for i := range arr {
		for j := 0; j < 17; j++ {
			mul := arr[i][j] * arr[i][j+1] * arr[i][j+2] * arr[i][j+3]
			if mul > max {
				max = mul
			}
		}
	}
	//up to down
	for i := 0; i < 17; i++ {
		for j := range arr[i] {
			mul := arr[i][j] * arr[i+1][j] * arr[i+2][j] * arr[i+3][j]
			if mul > max {
				max = mul
			}
		}
	}
	//first diagonal
	for i := 0; i < 17; i++ {
		for j := 0; j < 17; j++ {
			mul := arr[i][j] * arr[i+1][j+1] * arr[i+2][j+2] * arr[i+3][j+3]
			if mul > max {
				max = mul
			}
		}
	}
	//second diagonal
	for i := 3; i < 20; i++ {
		for j := 0; j < 17; j++ {
			mul := arr[i][j] * arr[i-1][j+1] * arr[i-2][j+2] * arr[i-3][j+3]
			if mul > max {
				max = mul
			}
		}
	}
	println(max)
}

