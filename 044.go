package main

var arr []int

func main() {
	//2500 is learned after experiment. The array includes the result
	arr = make([]int, 2500)
	for i := range arr {
		arr[i] = (i * (3*i - 1)) / 2
	}
	min := int(^uint(0) >> 1)
	for i := 1; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if isPentagonal(arr[j]-arr[i]) && isPentagonal(arr[j]+arr[i]) && arr[j]-arr[i] < min {
				min = arr[j] - arr[i]
			}
		}
	}
	println(min)
}

func isPentagonal(in int) bool {
	return search(in, 0, len(arr)-1)
}

func search(value, low, high int) bool {
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

