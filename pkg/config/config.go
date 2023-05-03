package config

import (
	"fmt"
	"time"

	"github.com/mohammedzee1000/standup/pkg/system"
	"github.com/pkg/errors"
)

const (
	ViewInPanels = iota + 1
	ViewSimple
)

const ViewUnknown = 999999

type ConfigSection struct {
	Name        string `json:"Name"`
	Short       string `json:"Short"`
	Description string `json:"Description"`
}

type ConfigDuration struct {
	Days   int `json:"Days"`
	Weeks  int `json:"Weeks"`
	Months int `json:"Months"`
}

type StandupView struct {
	ConfigView uint `json:"ConfigView"`
	ReportView uint `json:"ReportView"`
}

// Config represents the app config
type Config struct {
	Name           string           `json:"Name"`
	Sections       []*ConfigSection `json:"Sections,omitempty"`
	DefaultSection string           `json:"DefaultSection"`
	StartOfWeekDay string           `json:"StartOfWeekDay"`
	Holidays       []string         `json:"Holidays"`
	SectionsPerRow int              `json:"SectionsPerRow"`
	KeepOldUntil   *ConfigDuration  `json:"KeepOldUntil"`
	StandupView    *StandupView     `json:"StandupView"`
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
		c.Sections = append(c.Sections, &ConfigSection{
			Name:        "Worked On",
			Short:       "wo",
			Description: "Tasks worked on for the day",
		}, &ConfigSection{
			Name:        "Blockers",
			Short:       "bl",
			Description: "Blockers affect completion of tasks",
		}, &ConfigSection{
			Name:        "At Risk",
			Short:       "ar",
			Description: "May not complete due to some issue",
		}, &ConfigSection{
			Name:        "PR Reviews",
			Short:       "prr",
			Description: "Reviews of pull requests",
		}, &ConfigSection{
			Name:        "Meetings",
			Short:       "mt",
			Description: "Meetings attended during the day",
		})
		c.StartOfWeekDay = time.Monday.String()
		c.Holidays = []string{time.Saturday.String(), time.Sunday.String()}
		c.SectionsPerRow = 1
		c.KeepOldUntil = &ConfigDuration{
			Days:   0,
			Weeks:  0,
			Months: 1,
		}
		c.StandupView = &StandupView{
			ConfigView: ViewInPanels,
			ReportView: ViewInPanels,
		}

		err = c.WriteConfig()
		if err != nil {
			return c, err
		}
	}
	return c, nil
}
