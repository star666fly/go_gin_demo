package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type Conf struct {
	Url      string
	Port     string
	Username string
	DbName   string
	Password string
	Redis    struct {
		Host     string
		Password string
		DB       int
	}
}

func (c *Conf) GetConf() *Conf {
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
