// prime number generation code is taken from the Go tutorial:
// http://golang.org/doc/go_tutorial.html#tmp_361

package main

func main() {
	ch := make(chan int)
	go generate(ch)
	prime, result := 2, 1
	for prime <= 17 {
		//get new prime
		prime = <-ch
		newMultiple := prime
		for newMultiple <= 20 {
			newMultiple *= prime
		}
		newMultiple /= prime
		result *= newMultiple
		//filter the channel
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
	println(result)
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

