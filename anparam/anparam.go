package anparam

import "time"

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type Flags struct {
	ConfigPath string        `flag:"config_path" default:"data/config.json" usage:"Chemin du fichier de configuration"`
	Debug      bool          `flag:"debug" default:"true" usage:"Activer le mode debug"`
	LogPath    string        `flag:"log_path" default:"data/server_default.log" usage:"Port d'écoute"`
	Timeout    time.Duration `flag:"timeout" default:"30s" usage:"Timeout (ex: 10s, 1m)"`
}
