package config

import (
	"gopkg.in/yaml.v3"
	"log"
)

func New(rawConfig []byte, rawPkgInfo []byte) *Class {
	pkg := map[string]interface{}{}
	err := yaml.Unmarshal(rawPkgInfo, &pkg)
	if err != nil {
		log.Fatalln("Unable to parse pkg.info file")
	}
	cfg := Class{}
	err = yaml.Unmarshal(rawConfig, &cfg)
	if err != nil {
		log.Fatalln("Unable to parse the configuration file.")
	}
	cfg.AppName = pkg["name"].(string)
	cfg.AppVersion = pkg["version"].(string)

	return &cfg
}
