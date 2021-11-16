package gcontext

import (
	"bytes"
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

type keyType string

const (
	traceID         = "abc123"
	traceid keyType = "traceid"
)

var ctx context.Context

func init() {

	ctx = context.Background()
}

//Inherit show example
func Inherit() {
	ctx := context.Background()
	var cancel context.CancelFunc
	//ctx, cancel = context.WithCancel(ctx)
	//print(ctx, "root", traceid)
	//watch(ctx)
	ctx = context.WithValue(ctx, traceid, traceID)
	//print(ctx, "WithValue", traceid)
	//watch(ctx)
	//print(ctx, "WithCancel", traceid)
	watch(ctx)
	ctx, cancel = context.WithCancel(ctx)
	watch(ctx)
	ctx = context.WithValue(ctx, traceid, "xyz456")
	//print(ctx, "WithValue", traceid)
	watch(ctx)
	//time.Sleep(2 * time.Second)
	//cancel()
	//time.Sleep(1 * time.Second)

	var stop context.CancelFunc
	ctx, stop = signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	ctx = context.WithValue(ctx, traceid, "oooooooooo")
	watch(ctx)
	ctx = context.WithValue(ctx, traceid, "xxxxxxxx")
	watch(ctx)
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
	watch(ctx)
	ctx = context.WithValue(ctx, traceid, "yyyyyyyy")
	watch(ctx)

	defer stop()
	time.Sleep(6 * time.Second)
	cancel()
	fmt.Println("------graceful shuwdown")
	time.Sleep(1 * time.Second)

}
func print(ctx context.Context, remark string, key keyType) {
	//var rs string
	//if ctx.Value(key) != nil {
	//rs = ctx.Value(key).(string)
	//}
	rs, ok := ctx.Value(key).(string)
	if ok {
		fmt.Printf("remark:%s,key:%s,value:%s,at:%s\n", remark, key, rs, time.Now().Format("2006-01-02 15:04:05.000"))
	} else {
		fmt.Printf("remark:%s,key:%s,value:nil\n", remark, key)
	}

}
func watch(ctx context.Context) {
	go func() {
		gid := fmt.Sprint(getGID())
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("======== goroutin-id:%s,end at:%s\n", gid, time.Now().Format("2006-01-02 15:04:05.000"))
				return
			default:

				print(ctx, "watch goroutine:"+gid, traceid)
				time.Sleep(1 * time.Second)

			}
		}
	}()
}
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
