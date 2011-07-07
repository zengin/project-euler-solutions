package main

import (
	"strconv"
)

var leftArr, rightArr []int
var leftIndex, rightIndex int

func main() {
	leftArr, rightArr, leftIndex, rightIndex = make([]int, 10000), make([]int, 10000), 0, 0
	arr := make([]bool, 1000000)
	var prime, leftTrunctable, rightTrunctable bool
	var tmp int
	arr[2] = true
	arr[3] = true
	sum := 0
	count := 0
	for i := 5; i < len(arr); i += 2 {
		prime = true
		for j := 3; j < i; j += 2 {
			if arr[j] && i%j == 0 {
				prime = false
			}
		}
		if prime {
			arr[i], rightTrunctable, leftTrunctable = true, true, true
			if !isRightTrunctable(i / 10) {
				rightTrunctable = false
			}
			tmp, _ = strconv.Atoi(strconv.Itoa(i)[1:])
			if !isLeftTrunctable(tmp) {
				leftTrunctable = false
			}
			if rightTrunctable {
				if leftTrunctable && i > 10 {
					println(i)
					count++
					sum += i
				}
				rightArr[rightIndex] = i
				rightIndex++
			}
			if leftTrunctable {
				leftArr[leftIndex] = i
				leftIndex++
			}
		}
	}
	println(count, sum)
}

func isLeftTrunctable(in int) bool {
	switch in {
	case 2, 3, 5, 7:
		return true
	}
	for i := 0; i < leftIndex; i++ {
		if leftArr[i] == in {
			return true
		}
	}
	return false
}

func isRightTrunctable(in int) bool {
	switch in {
	case 2, 3, 5, 7:
		return true
	}
	for i := 0; i < rightIndex; i++ {
		if rightArr[i] == in {
			return true
		}
	}
	return false
}

