package config

import "wxapp/pkg/log"

type Conf struct {
	System *System  `yaml:"system"`
	Mysql  *Mysql   `yaml:"mysql"`
	Logger *log.Zap `yaml:"logger"`
}

type System struct {
	Env        string `yaml:"env"`
	HttpPort   string `yaml:"httpPort"`
	upLoadMode string `yaml:"upLoadMode"`
	Host       string `yaml:"host"`
	Domain     string `yaml:"domain"`
	Version    string `yaml:"version"`
}

type Mysql struct {
	UserName  string `yaml:"username"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	DbName    string `yaml:"dbName"`
	Charset   string `yaml:"charset"`
	ParseTime string `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
}
