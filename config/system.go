package config

type System struct {
	Mode       string `yaml:"mode"`
	HttpPort   string `yaml:"httpPort"`
	upLoadMode string `yaml:"upLoadMode"`
	Host       string `yaml:"host"`
	Domain     string `yaml:"domain"`
	Version    string `yaml:"version"`
}
