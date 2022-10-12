package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

var config Config

const (
	configName = "config.yaml"
)

type Config struct {
	ServiceName string `json:"serviceName"`
	ServicePath string `json:"servicePath"`
	ModulePath  string `json:"modulePath"`
}

func (c Config) String() string {
	return fmt.Sprintf("Config{ serviceName: %s, servicePath: %s, modulePath: %s }",
		c.ServiceName, c.ServicePath, c.ModulePath)
}

func init() {
	// set defaults
	config = Config{
		ServiceName: "blanksvc",
		ServicePath: "./",
		ModulePath:  "github.com/company/blanksvc",
	}

	// init config from config.yaml if exists
	if !isFileExists(configName) {
		return
	}
	log.Println("Reading configurations from config.yaml file...")
	yamlFile, err := ioutil.ReadFile(configName)
	if err != nil {
		log.Fatalf("failed to read config file:%v", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("failed to unmarshal yaml: %v", err)
	}
	log.Println("Config was loaded successfully")
}

func isFileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
