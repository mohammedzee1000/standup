package config

import (
	"fmt"
	"time"

	"github.com/mohammedzee1000/standup/pkg/system"
	"github.com/pkg/errors"
)

// Config repersents the app config
type Config struct {
	Name           string            `json:"Name"`
	SectionNames   map[string]string `json:"SectionNames,omitempty"`
	DefaultSection string            `json:"DefaultSection"`
	StartOfWeekDay string            `json:"StartOfWeekDay"`
	Holidays       []string          `json:"Holidays"`
}

// new creates a new Config struct
func new() *Config {
	return &Config{}
}

// WriteConfig writes config to a file
func (c *Config) WriteConfig() error {
	cfp, err := getConfigFilePath()
	if err != nil {
		return errors.Wrap(err, "unable to get config file path")
	}
	return system.WriteYamlFile(cfp, &c)
}

// ReadConfig reads configuration from file
func ReadConfig() (*Config, error) {
	var c *Config
	c = new()
	cfp, err := getConfigFilePath()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get config file path")
	}
	e, err := system.ReadYamlFile(cfp, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to read yaml file %w", err)
	}
	if !e {
		c.Name = "John Doe"
		c.DefaultSection = "Worked On"
		c.SectionNames = make(map[string]string)
		c.SectionNames["Worked On"] = "What tasks were worked on for the day"
		c.SectionNames["Blockers"] = "Blockers for completing tasks that are affecting completion"
		c.SectionNames["At Risk"] = "Possible non-completion due to various reasons"
		c.SectionNames["PR Reviews"] = "All pull request reviews"
		c.StartOfWeekDay = time.Monday.String()
		c.Holidays = []string{time.Saturday.String(), time.Sunday.String()}
		err = c.WriteConfig()
		if err != nil {
			return c, err
		}
	}
	return c, nil
}
