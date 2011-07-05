package main

func main() {
	//allocate 4 MB of space at once
	arr := make([]int, 1000000)
	max, maxLoc, notFound, cur, steps := 0, 0, true, uint64(0), 0
	for i := 1; i < 1000000; i++ {
		notFound = true
		cur = uint64(i)
		steps = 0
		for notFound {
			if cur == 1 {
				notFound = false
				arr[i] = steps
				if steps > max {
					max = steps
					maxLoc = i
				}
			} else if cur < uint64(i) {
				notFound = false
				steps = steps + arr[cur]
				arr[i] = steps
				if steps > max {
					max = steps
					maxLoc = i
				}
			} else {
				if cur%2 == 0 {
					cur /= 2
				} else {
					cur = 3*cur + 1
				}
			}
			steps++
		}
	}
	println(max, maxLoc)
}

