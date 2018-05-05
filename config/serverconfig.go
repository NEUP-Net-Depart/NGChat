package config

// 服务器配置结构体。
type ServerConfig struct {
	ListenAddr	string	`toml:"listen_addr"`
	GoMaxProcs	int		`toml:"go_maxprocs"`
}

