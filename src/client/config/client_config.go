package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ida-wong/leaf/config"

	"client/const/log"
)

var Client struct {
	LogLevel string
	LogPath  string
	WSAddr   string
	WSPort   int
	CertFile string
	KeyFile  string
	TCPAddr  string
	TCPPort  int
	//port := strings.SplitAfter(agent.RemoteAddr().String(), ":")
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
	configPath := filepath.Join(dirPath, "client.json")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Client)
	if err != nil {
		panic(err)
	}

	config.LogFlag = log_const.Flag
	config.LogLevel = Client.LogLevel
	config.LogPath = Client.LogPath
	config.ConsolePort = Client.ConsolePort
	config.ProfilePath = Client.ProfilePath
}
