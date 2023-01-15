package main

import (
	"bufio"
	"fmt"
	"golearning/functional/fib"
	"os"
)

// defer到return前打印, panic也能打印
// 参数在defer语句时计算
// defer顺序先进后出
func tryDefer() {
	defer fmt.Println(1)
	// defer fmt.Println(2)
	// fmt.Println(3)
	// panic("error occurred")
	// fmt.Println(4)
}

func writeFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		// panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// 写入io
	writer := bufio.NewWriter(file)
	// io 写入文件
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// tryDefer()
	writeFile("fib.txt")
}
