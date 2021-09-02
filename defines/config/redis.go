package config

import "github.com/gomodule/redigo/redis"

type EasyRedis struct {
	General		*CommonConf `yaml:"General"`
	Host		string		`yaml:"Host"`
	Port		string		`yaml:"Port"`
	MaxIdle		int			`yaml:"MaxIdle"`
}

type RedisIns struct {
	RedisClient		redis.Conn
}
