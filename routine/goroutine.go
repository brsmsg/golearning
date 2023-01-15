package main

import (
	"fmt"
	"runtime"
	"time"
)

// main退出后 所有goroutine也结束, main同样也是 go routine
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		// i 需要传参数，最后i会到10，out of range
		go func(i int) { // race codition 数据访问冲突
			for {
				// fmt.Printf("Hello from goroutine %d\n", i)
				a[i]++
				// 手动交出控制权，让别人有机会运行，一般不需要
				runtime.Gosched()
			}
		}(i)
	}
	// 防止main函数return
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

// Coroutine
// 轻量级线程
// 非抢占式多任务处理，coroutine主动交出控制权
// 抢占式 切换时需要记录更多上下文
// 编译器/解释器/虚拟机层面多任务，非OS级多任务
// 多个coroutine可以在一个线程上运行
