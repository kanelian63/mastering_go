package main

import (
	"fmt"
)

func main() {
	willClose := make(chan int, 10)
	// 채널에 데이터 쓰기
	willClose <- -1
	fmt.Println(willClose)
	willClose <- 0
	willClose <- 2

	// 채널에 데이터 비우기
	<-willClose
	fmt.Println(willClose)
	<-willClose
	fmt.Println(willClose)
	<-willClose
	fmt.Println(willClose)

	close(willClose)
	read := <-willClose
	fmt.Println(read)
}