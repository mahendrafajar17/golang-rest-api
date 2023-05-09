package configurations

import (
	"io/ioutil"

	"example.com/restapi/configurations/app"
	"example.com/restapi/configurations/artemis"
	"example.com/restapi/configurations/mysql"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DB      mysql.Config   `yaml:"mysql"`
	App     app.Config     `yaml:"app"`
	Artemis artemis.Config `yaml:"artemis"`
}

func LoadConfig(configpath string) *Config {
	yfile, err1 := ioutil.ReadFile(configpath)
	if err1 != nil {
		panic(err1)
	}
	config := Config{}
	err := yaml.Unmarshal(yfile, &config)
	if err != nil {
		panic(err)
	}
	return &config
}
