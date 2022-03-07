package main

import (
	"fmt"
	"sync"
)

// 锁复习
var counter = &safeCount{}

type safeCount struct {
	totalNum         int
	totalLetterCount int
	totalWordCount   int
	//...
	sync.Mutex
}

func (sc *safeCount) AddNum(totalNum, totalLetterCount, totalWordCount int) {
	sc.Lock()
	defer sc.Unlock()
	sc.totalNum += totalNum
	//...
}

// 把一系列需要加锁计算的部分抽象成一个结构体，并且向外提供一个方法，外部只管调用方法
// 不关心锁的问题，然后这个方法内部自己实现锁的加减
func lockReview() {
	counter.AddNum(10, 20, 300)
}

// channel 复习

var sharedClient = &esClientWithBuffer{}
var batchSize = 20

type esClientWithBuffer struct {
	batchBuffer [][]interface{}
	messageChan chan interface{}
	shortBuffer []interface{}
}

func (cli *esClientWithBuffer) pushBatch() {
	// queue operation
	// ...
}

func (cli *esClientWithBuffer) prepareBatch() {
	for msg := range cli.messageChan {
		if len(cli.shortBuffer) == batchSize {
			fmt.Println("shortBuffer 满了")
			cli.batchBuffer = append(cli.batchBuffer, cli.shortBuffer)
			cli.shortBuffer = []interface{}{}
		}
		cli.shortBuffer = append(cli.shortBuffer, msg)
	}
}

func (cli *esClientWithBuffer) pushToESService(data interface{}) {
	cli.messageChan <- data
	fmt.Println("发送 data 到 chan 里面去")
}

func channelReview() {

}

