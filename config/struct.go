package config

type Class struct {
	Command    string `yaml:"command"`
	Protocol   string `yaml:"protocol"`
	ScriptName string `yaml:"scriptName"`
	AppName    string
	AppVersion string
}
