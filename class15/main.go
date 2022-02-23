package main

import (
	"fmt"
	//"crypto/rand"
	"math/rand"
	"sync"
	"time"
)

func main() {
	//countDictRW()
	//producerAndConsumer()
	//waitGroup()
	//syncOnce()
	//syncMap()
	productionGo100()
}

func syncPackage() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Done()
	wg.Wait()

	once = sync.Once{}
	once.Do(func() {
		fmt.Println("doSomething")
	})

	cond := sync.Cond{}
	cond.Wait()
	cond.Broadcast()
	cond.Signal()

	maps := sync.Map{}
	fmt.Println(maps)
}

func countDictRW() {
	fmt.Println("开始数")
	var totalCount int64 = 0
	lock := sync.RWMutex{}
	wg := sync.WaitGroup{}
	workerCount := 500
	wg.Add(workerCount)

	for p := 0; p < workerCount; p++ {
		go func(p int) { // 读锁可以多个go routine同时拿到。
			fmt.Println(p, "读锁开始时间：", time.Now())
			lock.RLock()
			fmt.Println(p, "读锁拿到锁时间：", time.Now())
			time.Sleep(1 * time.Second)
			lock.RUnlock()
		}(p)
	}
	for p := 0; p < workerCount; p++ {
		go func() {
			defer wg.Done()
			//fmt.Print("正在统计第", p, "页, ")
			//r, _ := rand.Int(rand.Reader, big.NewInt(100))
			//fmt.Println("有", r.Int64(), "字")
			lock.Lock()
			//totalCount += r.Int64()
			defer lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("总共有：", totalCount, "字")
}

func producerAndConsumer() {
	s := &Store{Max: 50}
	s.pCond = sync.NewCond(&s.lock)
	s.cCond = sync.NewCond(&s.lock)
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
	fmt.Println(s.DataCount)

}

func waitGroup() {
	runnerCount := 10
	var runners []Runner

	wg := sync.WaitGroup{}
	wg.Add(runnerCount)

	startPointWg := sync.WaitGroup{}
	startPointWg.Add(1)

	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("%d", i),
		})
	}

	for _, runnerItem := range runners {
		go runnerItem.Run(&startPointWg, &wg)
	}

	fmt.Println("各就位")
	time.Sleep(1 * time.Second)
	fmt.Println("预备：跑")

	startPointWg.Done()

	wg.Wait()
	fmt.Println("赛跑结束")
}

func syncMap() {
	m := sync.Map{}
	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Store(i, 1)
			for {
				v, ok := m.Load(i)
				if !ok {
					continue
				}
				m.Store(i, v.(int)+1)
				fmt.Println("i=", v)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
}

func syncOnce(standard []string) {
	once.Do(func() {
		globalRank.standard = standard
	})
}

type rank struct {
	standard []string
}

var globalRank = &rank{}
var once sync.Once = sync.Once{}

type Runner struct {
	Name string
}

func (r Runner) Run(startPointWg, wg *sync.WaitGroup) {
	defer wg.Done()
	startPointWg.Wait()

	start := time.Now()
	fmt.Println(r.Name, "开始跑@", start)
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	finish := time.Now()
	fmt.Println(r.Name, "跑到终点，用时：", finish.Sub(start))
}

type Store struct {
	DataCount int
	Max       int
	lock      sync.Mutex
	pCond     *sync.Cond
	cCond     *sync.Cond
}

type Producer struct{}

func (Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == s.Max {
		fmt.Println("生产者在等仓库拉货")
		s.pCond.Wait()
	}
	fmt.Println("开始生产+1")
	s.DataCount++
	s.pCond.Signal()
}

type Consumer struct{}

func (Consumer) consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DataCount == 0 {
		fmt.Println("消费者在等货")
		s.cCond.Wait()
	}
	fmt.Println("消费者消费-1")
	s.DataCount--
	s.cCond.Signal()
}
