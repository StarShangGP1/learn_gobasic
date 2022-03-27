package main

import (
	"fmt"
	gobmi "github.com/armstrongli/go-bmi"
	"learn_gobasic/homework/class29/interface"
	"learn_gobasic/pkg/apis"
	"sort"
	"sync"
)

var _ _interface.ServerInterface = &CircleCache{}

type CircleItem struct {
	ID            uint32
	ReleaseTime   int64
	PID           uint32
	PName         string
	Content       string
	ByTimeTall    float32
	ByTimeWeight  float32
	ByTimeFatRate float32
	Visible       bool
}

type CircleCache struct {
	items     []CircleItem
	itemsLock sync.Mutex
}

func (c2 *CircleCache) DeleteStatus(id uint32) error {
	return nil
}

func (c2 *CircleCache) GetList() ([]*apis.TopPost, error) {
	//TODO implement me
	panic("implement me")
}

func (c2 *CircleCache) PostStatus(c *apis.Circle) error {
	bmi, _ := gobmi.BMI(float64(c.ByTimeTall), float64(c.ByTimeWeight))
	fatRate := gobmi.CalcFatRate(bmi, int(c.ByTimeAge), "男")
	c2.inputRecord(c, float32(fatRate))
	return nil
}

func (c2 *CircleCache) inputRecord(c *apis.Circle, fatRate float32) {
	c2.itemsLock.Lock()
	defer c2.itemsLock.Unlock()

	c2.items = append(c2.items, CircleItem{
		ID:            c.Id,
		ReleaseTime:   c.ReleaseTime,
		PID:           c.PId,
		PName:         c.PName,
		Content:       c.Content,
		ByTimeTall:    c.ByTimeTall,
		ByTimeWeight:  c.ByTimeWeight,
		ByTimeFatRate: fatRate,
		Visible:       c.Visible,
	})

}
func (c2 *CircleCache) DeletePost(pid uint32) error {
	for i, item := range c2.items {
		if item.PID == pid {
			c2.items[i].Visible = false
		}
	}
	return nil
}

func (c2 *CircleCache) ListPost() ([]*apis.TopPost, error) {
	c2.itemsLock.Lock()
	defer c2.itemsLock.Unlock()
	sort.Slice(c2.items, func(i, j int) bool {
		return c2.items[i].ReleaseTime > c2.items[j].ReleaseTime
	})

	fmt.Printf(" number : %d", len(c2.items))

	count := 0
	out := make([]*apis.TopPost, 0, 10)
	for _, item := range c2.items {
		if item.Visible {
			out = append(out, &apis.TopPost{
				ID:            item.ID,
				ReleaseTime:   item.ReleaseTime,
				PId:           item.PID,
				PName:         item.PName,
				Content:       item.Content,
				ByTimeTall:    item.ByTimeTall,
				ByTimeWeight:  item.ByTimeWeight,
				ByTimeFatRate: item.ByTimeFatRate,
			})
			count++
		}
		if count > 10 {
			break
		}
	}
	fmt.Printf(" Get post : %+v", out)
	return out, nil
}

func NewCircleCache() *CircleCache {
	return &CircleCache{
		items: make([]CircleItem, 0, 1000),
	}
}
