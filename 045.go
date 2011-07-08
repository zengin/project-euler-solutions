package main

var arr []int64 //pentagonals

func main() {
	arr = make([]int64, 80000)
	for i := 1; i < len(arr); i++ {
		arr[i] = int64(i) * (3*int64(i) - 1) / 2
	}
	var tmp int64
	i := 144
loop:
	for {
		tmp = int64(i) * int64(2*i-1)
		if isPentagonal(tmp) { //every hexagonal is also triangle
			println(tmp)
			break loop
		}
		i++
	}
}

func isPentagonal(in int64) bool {
	return search(in, 1, len(arr)-1)
}

func search(value int64, low, high int) bool {
	if high < low {
		return false
	}
	mid := low + (high-low)/2
	if arr[mid] > value {
		return search(value, low, mid-1)
	} else if arr[mid] < value {
		return search(value, mid+1, high)
	}
	return true
}

