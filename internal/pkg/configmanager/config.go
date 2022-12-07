package configmanager

import (
	"fmt"
)

// Config is a general config in DevStream.
type Config struct {
	Config CoreConfig     `yaml:"config"`
	Vars   map[string]any `yaml:"vars"`
	Tools  Tools          `yaml:"tools"`
}

func (c *Config) renderInstanceIDtoOptions() {
	for _, t := range c.Tools {
		if t.Options == nil {
			t.Options = make(RawOptions)
		}
		t.Options["instanceID"] = t.InstanceID
	}
}

func (c *Config) validate() error {
	if c.Config.State == nil {
		return fmt.Errorf("config.state is not defined")
	}

	if err := c.Tools.validateAll(); err != nil {
		return err
	}
	return nil
}
