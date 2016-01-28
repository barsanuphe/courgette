package courgette

// Config holds all input necessary to deal with a Collection
type Config struct {
	Root       string
	Incoming   string
	Cameras    map[string]string
	Lenses     map[string]string
	NameFormat string
	GroupBy    string // month, year, other??
}

// Load the configuration
func (c *Config) Load() (err error) {
	return
}

// Check the configuration
func (c *Config) Check() (err error) {
	return
}
