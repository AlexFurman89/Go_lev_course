package main

import "fmt"

func generate(out chan int) {
	for i := 1; i <= 100; i++ {
		out <- i
	}
	close(out)
}
func filterEven(in chan int, out chan int) {
	for v := range in {
		if v%2 == 0 {
			out <- v
		}
	}
	close(out)
}
func square(in chan int, out chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}
func sum(in chan int, out chan int) {
	sum := 0
	for v := range in {
		sum += v
	}
	out <- sum
	close(out)
}

func main() {

	ch_generate := make(chan int, 100)
	ch_filterEven := make(chan int)
	ch_square := make(chan int)
	ch_sum := make(chan int)

	go generate(ch_generate)
	go filterEven(ch_generate, ch_filterEven)
	go square(ch_filterEven, ch_square)
	go sum(ch_square, ch_sum)
	result := <-ch_sum
	fmt.Println(result)

}
