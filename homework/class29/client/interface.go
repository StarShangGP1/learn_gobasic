package main

import (
	"learn_gobasic/pkg/apis"
	"sync"
)

type ClientInterface interface {
	ReadPostInformation() apis.Circle
	GetPId() uint32
}

var _ ClientInterface = &fakeCircleInterface{}

type fakeCircleInterface struct {
	pId          uint32
	pName        string
	content      string
	byTimeTall   float32
	byTimeWeight float32
	byTimeAge    uint32
	sync.Mutex
}

func (f *fakeCircleInterface) ReadPostInformation() apis.Circle {
	c := apis.Circle{
		PId:          f.pId,
		PName:        f.pName,
		Content:      f.content,
		ByTimeTall:   f.byTimeTall,
		ByTimeWeight: f.byTimeWeight,
		ByTimeAge:    f.byTimeAge,
	}
	return c
}
func (f *fakeCircleInterface) GetPId() uint32 {
	return f.pId
}
