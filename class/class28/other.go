package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

import (
	"github.com/google/uuid"
)

func main1() {
	ms := &mainServer{workerCount: 100, requestCh: make(chan *request)}
	ms.startWorkers(context.TODO())

	sendConcurrency := 1000
	wg := sync.WaitGroup{}
	wg.Add(sendConcurrency)
	for i := 0; i < sendConcurrency; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 10000; i++ {
				func() {
					request := &request{
						ctx:      context.TODO(),
						id:       fmt.Sprintf("%d-%s", i, uuid.New().String()),
						resultCh: make(chan string),
					}
					defer close(request.resultCh)
					start := time.Now()
					resp := ms.HandleRequest(request)
					finish := time.Now()
					fmt.Printf("dur: %s, req: %s, resp: %s\n", finish.Sub(start), request.id, resp)
				}()
				time.Sleep(100 * time.Millisecond)
			}
		}()
	}
	wg.Wait()
}

type request struct {
	ctx      context.Context
	id       string
	resultCh chan string
}

type mainServer struct {
	workerCount int
	requestCh   chan *request
}

func (ms *mainServer) HandleRequest(r *request) string {
	ms.requestCh <- r
	output := <-r.resultCh
	// fmt.Println("output:", output)
	return output
}

func (ms *mainServer) startWorkers(ctx context.Context) {
	for i := 0; i < ms.workerCount; i++ {
		go ms.requestWorker(ctx)
	}
}

func (ms *mainServer) requestWorker(ctx context.Context) {
	for r := range ms.requestCh {
		select {
		case <-ctx.Done():
			return
		default:
		}
		ms.sendRequest(r)
	}
}

func (ms *mainServer) sendRequest(r *request) {
	ctx, cancel := context.WithTimeout(r.ctx, 1*time.Second)
	defer cancel()
	fmt.Println("sending request to ali server for request:", r.id, "with context:", ctx)
	time.Sleep(100 * time.Millisecond)
	fmt.Println("response for request", r.id, ": success")
	r.resultCh <- base64.RawStdEncoding.EncodeToString([]byte(r.id))
}
