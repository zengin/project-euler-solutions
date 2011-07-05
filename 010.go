package main

func main() {
	arr := make([]int, 2000000)
	for i := range arr {
		arr[i] = i
	}
	i := 2
	sum := uint64(0)
	for i < 2000000 {
		sum += uint64(i)
		for j := i; j < 2000000; j += i {
			arr[j] = -1
		}
		var c int
		for c = i; c < 2000000 && arr[c] == -1; c += 1 {
		}
		i = c
	}
	println(sum)
}

