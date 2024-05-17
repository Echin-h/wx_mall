package config

type Zap struct {
	Prefix       string `default:"gin-vue-admin" yaml:"prefix" json:"prefix"`
	Level        string `default:"panic" yaml:"level" json:"level"`
	Director     string `default:"director" yaml:"director" json:"director"`
	EncoderLevel string `default:"LowercaseLevelEncoder" yaml:"encoderLevel" json:"encoderLevel"`
	MaxAge       int    `default:"30" yaml:"maxAge" json:"maxAge"`
	MaxSize      int    `default:"600" yaml:"maxSize" json:"maxSize"`
	MaxBackups   int    `default:"3" yaml:"maxBackups" json:"maxBackups"`
	Compress     bool   `default:"true" yaml:"compress" json:"compress"`
	ShowLine     bool   `default:"true" yaml:"showLine" json:"showLine"`
	LogInConsole bool   `default:"true" yaml:"logInConsole" json:"logInConsole"`
	Encoder      string `default:"console" yaml:"encoder" json:"encoder"`
}
