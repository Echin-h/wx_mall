package config

type Conf struct {
	System *System `yaml:"system"`
	Mysql  *Mysql  `yaml:"mysql"`
	Zap    *Zap    `yaml:"logger"`
	Redis  *Redis  `yaml:"redis"`
}
