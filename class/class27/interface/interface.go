package _interface

import "learn_gobasic/pkg/apis"

// ServerInterface 服务端接口
type ServerInterface interface {

	// RegisterPersonalInformation 注册个人信息
	RegisterPersonalInformation(pi *apis.PersonalInformation) error

	// UpdatePersonalInformation 更新个人信息
	UpdatePersonalInformation(pi *apis.PersonalInformation) (*apis.PersonalInformationFatRate, error)

	// GetFatRate 获取体脂率
	GetFatRate(name string) (*apis.PersonalRank, error)

	// GetTop 获取排名
	GetTop() ([]*apis.PersonalRank, error)
}

// RankInitInterface 排名初始化接口
type RankInitInterface interface {
	Init() error
}