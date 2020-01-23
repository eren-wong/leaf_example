package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ida-wong/leaf/config"

	"server/const/log"
)

var Server struct {
	LogLevel    string
	LogPath     string
	WSAddr      string
	CertFile    string
	KeyFile     string
	TCPAddr     string
	MaxConnNum  int
	ConsolePort int
	ProfilePath string
}

func init() {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	dirPath := filepath.Dir(exePath)
	configPath := filepath.Join(dirPath, "server.json")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Server)
	if err != nil {
		panic(err)
	}

	config.LogFlag = log_const.Flag
	config.LogLevel = Server.LogLevel
	config.LogPath = Server.LogPath
	config.ConsolePort = Server.ConsolePort
	config.ProfilePath = Server.ProfilePath
}
