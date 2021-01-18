package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func writeLog(i int, ch chan int) {
	file, err := os.OpenFile("config.txt", os.O_RDWR|os.O_APPEND, os.ModeAppend)
	defer file.Close()
	if err != nil {
		panic(err)
	}
	file.WriteString("Write File--" + strconv.Itoa(i) + "\n")
	ch <- i
}
func main() {
	fmt.Println("Start------")
	start := time.Now().UnixNano()

	ch := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		go writeLog(i, ch)
	}
	// 不能用 range 因为range在通道不关闭的时候自己也不会停止读取
	for j := 0; j < 5; j++ {
		fmt.Println("j=%d, ch=%d", j, <-ch)
	}
	end := time.Now().UnixNano()
	fmt.Println("End------", (end-start)/1e6)
}

// ch <- v    // 发送值v到Channel ch中
// v := <-ch  // 从Channel ch中接收数据，并将数据赋值给v
