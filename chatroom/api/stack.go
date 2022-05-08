package api

import (
	"chat-project/proto"
)

type Stack struct {
	Data []*proto.ChatRecords
}

func (s *Stack) Push(data *proto.ChatRecords) {
	s.Data = append([]*proto.ChatRecords{data}, s.Data...)
}
