package main

import (
	"math"
	"container/list"
)

var writable []bool
var limit int
var abunList *list.List

func main() {
	abunList = list.New()
	limit = 28124
	writable = make([]bool, limit)
	for i := 1; i < limit; i++ {
		writable[i] = false
		if i < sumOfProperDivisors(i) {
			abunList.PushBack(i)
		}
	}

	checkWritables()

	sum := 0
	for i := 1; i < limit; i++ {
		if !writable[i] {
			sum += i
		}
	}
	println(sum)
}

func checkWritables() {
	for i := abunList.Front(); i != nil; i = i.Next() {
		for j := i; j != nil; j = j.Next() {
			if sum := i.Value.(int) + j.Value.(int); sum < limit {
				writable[sum] = true
			}
		}
	}
}

func sumOfProperDivisors(input int) int {
	//checking until the square root is enough
	//since a divisor less than the square root corresponds
	//the other divisor greater than the square root
	limit, sum := int(math.Sqrt(float64(input))), 0
	for i := 1; i <= limit; i++ {
		if input%i == 0 {
			if i == input/i {
				sum += i
			} else {
				sum += i + (input / i)
			}
		}
	}
	return sum - input
}

