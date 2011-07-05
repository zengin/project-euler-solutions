//same code with 018 except the input file and number of lines.

package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

var arr [][]int
var bests [][]int

func main() {
	fileBuf, err := ioutil.ReadFile("067_input.txt")
	if err != nil {
		panic(err.String())
	}
	fileStr := strings.Trim(string(fileBuf), "")
	oneDArrStr := strings.Split(fileStr, "\n", -1)
	var line []string
	arr = make([][]int, 100)
	bests = make([][]int, len(arr))
	for i := range arr {
		line = strings.Split(oneDArrStr[i], " ", -1)
		arr[i] = make([]int, len(line))
		bests[i] = make([]int, len(arr[i]))
		for j := range line {
			arr[i][j], _ = strconv.Atoi(line[j])
			bests[i][j] = -1
		}
	}
	println(bestSum(0, 0))
}

func bestSum(i, j int) int {
	var result int
	if bests[i][j] != -1 {
		result = bests[i][j]
	} else if i == len(arr)-1 {
		result = arr[i][j]
		return arr[i][j]
	} else {
		sumLeft, sumRight := bestSum(i+1, j), bestSum(i+1, j+1)
		if sumLeft > sumRight {
			result = arr[i][j] + sumLeft
		} else {
			result = arr[i][j] + sumRight
		}
	}
	bests[i][j] = result
	return result
}

