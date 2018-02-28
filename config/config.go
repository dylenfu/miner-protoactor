package config

import (
	"github.com/naoina/toml"
	"os"
)

type GlobalConfig struct {
	Title string `required:"true"`
	Owner struct {
		Name string
	}
	Mysql MysqlOptions
}

func LoadConfig(file string) *GlobalConfig {
	if "" == file {
		dir, _ := os.Getwd()
		file = dir + "/config/relay.toml"
	}

	io, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer io.Close()

	c := &GlobalConfig{}
	if err := toml.NewDecoder(io).Decode(c); err != nil {
		panic(err)
	}

	return c
}

type MysqlOptions struct {
	Hostname    string
	Port        string
	User        string
	Password    string
	DbName      string
	TablePrefix string
	Debug       bool
}
