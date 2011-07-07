package main

var total int
var coins []int

func main() {
	total, coins = 200, []int{1, 2, 5, 10, 20, 50, 100, 200}
	println(count(total, len(coins)))
}

func count(n, m int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	} else if m <= 0 && n >= 1 {
		return 0
	}
	return count(n, m-1) + count(n-coins[m-1], m)
}

