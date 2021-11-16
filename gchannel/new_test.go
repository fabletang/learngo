package gchannel

import (
	"fmt"
	"testing"
	"time"
)

func TestNewChannel(t *testing.T) {
	newChannel()
}

func TestBlock(t *testing.T) {
	ch := make(chan int, 1)
	//ch <- 1
	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
		ch <- 2
		<-ch
		<-ch
	}()

	after := time.After(3 * time.Second)
	tick := time.Tick((1*1000 - 10) * time.Millisecond)
	for {
		select {
		case <-after:
			fmt.Printf("超时退出,at %s\n", time.Now().Format("2006-01-02 15:04:05.000"))
			close(ch) //channel不需要通过close释放资源，只要没有goroutine持有channel，相关资源会自动释放
			return
		case <-time.Tick((1*1000 - 10) * time.Millisecond):
			fmt.Printf("....tick wrong,at %s\n", time.Now().Format("2006-01-02 15:04:05.000"))
		case <-tick:
			fmt.Printf("....tick right,at %s\n", time.Now().Format("2006-01-02 15:04:05.000"))

		case rs := <-ch:

			go func() {
				fmt.Printf("run block()...%d,at %s\n", rs, time.Now().Format("2006-01-02 15:04:05.000"))
				block()
			}()
		}
	}
}
