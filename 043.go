package main

import (
	"strconv"
)

var result int64
var divisors []int
var tmp int
var tmp64 int64

func main() {
	result = int64(0)
	divisors = []int{2, 3, 5, 7, 11, 13, 17}
	allPerms("", "0123456789")
	println(result)
}

func allPerms(pre, s string) {
	if len(s) == 0 {
		checkTheProperty(pre)
	} else {
		for i := 0; i < len(s); i++ {
			allPerms(pre+s[i:i+1], s[0:i]+s[i+1:len(s)])
		}
	}
}

func checkTheProperty(s string) {
	if s[0] == '0' {
		return
	} else {
		for i := 1; i < 8; i++ {
			if tmp, _ = strconv.Atoi(s[i : i+3]); tmp%divisors[i-1] != 0 {
				return
			}
		}
	}
	tmp64, _ = strconv.Atoi64(s)
	result += tmp64
}

