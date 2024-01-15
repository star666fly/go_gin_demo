package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type conf struct {
	Url      string
	Port     string
	Username string
	DbName   string
	Password string
}

func (c *conf) getConf() *conf {
	ymlFile, err := os.ReadFile("config/db.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(ymlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}
