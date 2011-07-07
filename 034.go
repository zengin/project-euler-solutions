package main

func main() {
	factorials := make([]int, 10)
	factorials[0] = 1
	for i := 1; i <= 9; i++ {
		factorials[i] = factorials[i-1] * i
	}
	//find limit
	powerOfTen := 10
	limit := factorials[9]
	for limit > powerOfTen {
		limit *= 2
		powerOfTen *= 10
	}

	sum := 0
	var localSum, j int
	for i := 10; i < limit; i++ {
		localSum, j = 0, i
		for j > 9 {
			localSum += factorials[j%10]
			j /= 10
		}
		localSum += factorials[j]

		if i == localSum {
			sum += i
		}
	}
	println(sum)
}

