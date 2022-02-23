package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//producerAndConsumer()
	prime()
}

func prime() {
	startTime := time.Now()

	maxNum := 1000
	workerNum := 4

	result := make(chan int, maxNum/4)
	baseNumCh := make(chan int, 10)

	wg := sync.WaitGroup{}
	wg.Add(workerNum)

	for num := 0; num < workerNum; num++ {
		go func() {
			defer wg.Done()

			for oNum := range baseNumCh {
				if isPrime(oNum) {
					result <- oNum
				}
			}
		}()
	}
	for num := 2; num <= maxNum; num++ {
		baseNumCh <- num
	}
	close(baseNumCh)
	defer close(result)
	wg.Wait()
	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}

func isPrime(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func producerAndConsumer() {
	s := &Store{Max: 50}
	s.instrument()
	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(500 * time.Millisecond)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			for {
				time.Sleep(500 * time.Millisecond)
				Consumer{}.consume(s)
			}
		}()
	}

	time.Sleep(1 * time.Second)

}

type Store struct {
	init  sync.Once
	Max   int
	store chan int
}

func (s *Store) instrument() {
	s.init.Do(func() {
		s.store = make(chan int, s.Max)
	})
}

type Producer struct{}

func (Producer) Produce(s *Store) {
	fmt.Println("开始生产+1")
	s.store <- rand.Int()
}

type Consumer struct{}

func (Consumer) consume(s *Store) {
	fmt.Println("消费者消费-1", <-s.store)
}
