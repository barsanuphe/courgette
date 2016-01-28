package courgette

import "github.com/spf13/viper"

// Config holds all input necessary to deal with a Collection
type Config struct {
	Root       string
	Incoming   string
	Cameras    map[string]string
	Lenses     map[string]string
	NameFormat string
	GroupBy    string // month, year, other??
}

// String displays Config information.
func (c *Config) String() (desc string) {
	return c.Root
}

// Load the configuration.
func (c *Config) Load(configPath string) (err error) {
	conf := viper.New()
	conf.SetConfigName(configPath)
	conf.SetConfigType("yaml")
	// TODO xdg
	conf.AddConfigPath("$HOME/.config/courgette")
	conf.AddConfigPath(".")
	err = conf.ReadInConfig()
	if err != nil {
		c.Root = conf.GetString("config.directory")
		c.Incoming = conf.GetString("config.incoming")
		c.Lenses = conf.GetStringMapString("lenses")
		c.Cameras = conf.GetStringMapString("cameras")
	}
	return
}

// Check the configuration.
func (c *Config) Check() (err error) {
	return
}
