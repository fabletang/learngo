package goroutine

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

//Join fork-join
func Join() {
	var wg sync.WaitGroup

	str := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		str = "world"
	}()
	wg.Wait()
	time.Sleep(50 * time.Millisecond)
	fmt.Println(str)

}

//JoinErr err example for join
func JoinErr() {
	var wg sync.WaitGroup
	str := "hello"
	var i int
	for i, str = range []string{"a", "b", "c"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("index:%d value:%s,goroutine-id:%d\n", i, str, getGID())
		}()

	}
	wg.Wait()
	fmt.Printf("JoinErr:last:%s\n", str)
}

//JoinEnum enum example for join,the result: disorder but no-repeat
func JoinEnum() {
	var wg sync.WaitGroup
	str := "hello"
	var i int
	for i, str = range []string{"a", "b", "c"} {
		wg.Add(1)
		go func(index int, s string) {
			defer wg.Done()
			//fmt.Printf("index:%d value:%s\n", index, s)
			fmt.Printf("index:%d value:%s,goroutine-id:%d\n", i, str, getGID())

		}(i, str)
	}
	wg.Wait()
	fmt.Printf("JoinErr:last:%s\n", str)
}
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
