package main

var months []int
var count int

func main() {
	months = make([]int, 12)

	//calculate first days of the months in 1900
	daysInMonths := []int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30}
	months[0] = 0
	for i := range daysInMonths {
		months[i+1] = (months[i] + daysInMonths[i]) % 7
	}

	count = 0
	for year := 1900; year < 2000; year++ {
		for i := range months {
			if (year%4 == 0 && i <= 1) || ((year+1)%4 == 0 && i > 1) {
				months[i] = (months[i] + 366) % 7
			} else {
				months[i] = (months[i] + 365) % 7
			}
			if months[i] == 6 {
				count++
			}
		}
	}
	println(count)
}

