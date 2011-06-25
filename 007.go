// prime number generation code is taken from the Go tutorial:
// http://golang.org/doc/go_tutorial.html#tmp_361

package main

func main() {
	ch := make(chan int)
	go generate(ch)
	prime, count := 2, 0
	for count < 10001 {
		//get new prime
		prime = <-ch
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
		count++
	}
	println(prime)
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

