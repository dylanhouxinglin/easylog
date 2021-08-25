package config

import "time"

/*
 通用配置相关
 */
type CommonConf struct {
	ConnTimeOut		time.Duration		`yaml:"ConnTimeOut"`
	ReadTimeOut		time.Duration		`yaml:"ReadTimeOut"`
	WriteTimeOut	time.Duration		`yaml:"WriteTimeOut"`
	Retry			int					`yaml:"Retry"`
}
