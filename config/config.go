package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
)

type Log struct {
	Type     string `json:"type" `
	FilePath string `json:"filePath"`
}
type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Redis struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Password string `json:"password"`
}

type Config struct {
	Log   Log   `json:"log"`
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`
}

var Conf *Config

func init()  {
	if Conf != nil {
		return
	}
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return
	}
	f := filepath.Dir(file) + "/config.json"

	b := isFile(f)
	if !b {
		return
	}

	bytes, _ := ioutil.ReadFile(f)
	err := json.Unmarshal(bytes, &Conf)
	if err != nil {
		return
	}

	_, err = json.Marshal(&Conf)
	return
}

func isFile(f string) bool {
	fi, e := os.Stat(f)
	if e != nil {
		return false
	}
	return !fi.IsDir()
}
