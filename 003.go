// prime number generation code is taken from the Go tutorial:
// http://golang.org/doc/go_tutorial.html#tmp_361

package main

func main() {
	target := int64(600851475143)
	ch := make(chan int)
	go generate(ch)
	prime, result := 2, 0
	for int64(prime) < target {
		//get new prime
		prime := <-ch
		if target%int64(prime) == 0 {
			result = prime
			//update target
			for target%int64(prime) == 0 {
				target = target / int64(prime)
			}
		}
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

