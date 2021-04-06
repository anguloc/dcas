package config

import (
	"errors"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Get(key string) (string, error) {
	if key == "logFilePath" {
		return "temp/logs/app.log", nil
	} else if key == "httpPort" {

	}
	switch key {
	case "logFilePath":
		return "temp/logs/app.log", nil
	case "httpPort":
		return "9505", nil
	case "tcpPort":
		return "9506", nil
	case "mysql":

	default:
		return "", errors.New("unknown key")
	}
}
