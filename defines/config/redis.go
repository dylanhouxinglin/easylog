package config

type EasyRedis struct {
	Host		string		`yaml:"Host"`
	Port		string		`yaml:"Port"`
	MaxIdle		int			`yaml:"MaxIdle"`
}
