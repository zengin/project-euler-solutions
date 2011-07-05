package main

var used []bool

func main() {
	used = make([]bool, 10)
	for i := range used {
		used[i] = false
	}
	remaining := 999999
	perm := 1
	index := 0
	for i := 9; i > 0; i-- {
		perm = 1
		for j := i; j > 0; j-- {
			perm *= j
		}
		index = remaining / perm
		remaining = remaining % perm
		j := 0
		for j = 0; j < index || used[j]; j++ {
			if used[j] {
				index++
			}
		}
		used[j] = true
		print(j)
	}
	println(remaining)
}

