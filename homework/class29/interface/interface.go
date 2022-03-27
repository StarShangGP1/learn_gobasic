package _interface

import "learn_gobasic/pkg/apis"

type ServerInterface interface {
	PostStatus(c *apis.Circle) error
	DeleteStatus(id uint32) error
	GetList() ([]*apis.TopPost, error)
}

type CircleInitInterface interface {
	Init() error
}
