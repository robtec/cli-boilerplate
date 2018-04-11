package config

import (
	"bufio"
	"os"
	"path/filepath"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

type Config struct {
	mux sync.Mutex
}

func NewConfig() *Config {
	return &Config{}
}

var filename = ".cli"

func (c *Config) Load() (interface{}, error) {
	c.mux.Lock()
	defer c.mux.Unlock()

	var data interface{}

	filepath := getFilePath()

	_, err := toml.DecodeFile(filepath, &data)

	return data, err
}

func (c *Config) Save(data interface{}) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	filepath := getFilePath()
	createFile(filepath)

	var file, err = os.OpenFile(filepath, os.O_RDWR, 0644)

	w := bufio.NewWriter(file)
	err = toml.NewEncoder(w).Encode(data)

	return err
}

func getFilePath() string {

	home, _ := homedir.Dir()
	filepath := filepath.Join(home, filename)

	return filepath
}

func createFile(path string) error {

	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		file, _ := os.Create(path)
		defer file.Close()
	}
	return err
}
