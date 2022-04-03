package config

var Data *ConfigModel

// Here we declare custom structures which we will be using inside the app!
type ConfigModel struct {
	// Add here Your structure as you wish!
	HelloWorld string `yaml:"hello_wolrd" mapstructure:"hello_wolrd" default:"hey"`
	Aaaaa      string
}
