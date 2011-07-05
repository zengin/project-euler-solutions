package main

var arr [][]uint64

func main() {
	arr = make([][]uint64, 21)
	for i := range arr {
		arr[i] = make([]uint64, 21)
		for j := range arr[i] {
			arr[i][j] = uint64(0)
		}
	}
	arr[20][20] = uint64(1)
	println(noOfRoutes(0, 0))
}

func noOfRoutes(i, j int) uint64 {
	if arr[i][j] != uint64(0) {
		return arr[i][j]
	}
	var result uint64
	if i < 20 && j < 20 {
		result = noOfRoutes(i+1, j) + noOfRoutes(i, j+1)
	} else if i < 20 && j == 20 {
		result = noOfRoutes(i+1, j)
	} else {
		result = noOfRoutes(i, j+1)
	}
	arr[i][j] = result
	return result
}

