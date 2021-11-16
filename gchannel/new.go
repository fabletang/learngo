package gchannel

import "fmt"

func newChannel() {
	ch := make(chan int, 1)
	ch <- 1
	rs := <-ch * 2
	fmt.Printf("in:%d,out:%d\n", 1, rs)
}

func block() (b bool) {
	b = true
	ch := make(chan string, 2)
	ch <- "a"
	ch <- "b"
	ch <- "c"
	b = false
	return b

}
