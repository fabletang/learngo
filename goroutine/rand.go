package goroutine

import (
	"fmt"
	"time"
)

//RandPrint rand print
func RandPrint() {
	go func() {
		fmt.Println("hello")
	}()
	go func() {
		fmt.Println("world")
	}()
	time.Sleep(50 * time.Millisecond)
	fmt.Println("!!!")
}
