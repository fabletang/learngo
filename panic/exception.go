package panic

import "fmt"

func showIf() {
	if a := 1; a == 1 {
		fmt.Printf("%d\n", a)
	}
	return
}

func showPanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%s:%v\n", "恢复运行", err)
		}
	}()
	panic("恐慌!!!!")
}
