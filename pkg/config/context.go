package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Context repersents the run context
type Context struct {
	configuration *Config
	dataDir       string
}

// New returns a context
func New() (*Context, error) {
	dd, err := getDataDir()
	if err != nil {
		return nil, err
	}
	cf, err := ReadConfig()
	c := &Context{
		configuration: cf,
		dataDir:       dd,
	}
	return c, nil
}

// DataDir gets the data dir
func (c *Context) DataDir() string {
	return c.dataDir
}

// GetStandUpFileDir returns path to use for standup record file based on day, month and year
func (c *Context) GetStandUpFileDir() (string, error) {
	sfd := filepath.Join(c.dataDir, "standups")
	_, err := os.Stat(sfd)
	if err != nil {
		if os.IsNotExist(err) {
			err1 := os.Mkdir(sfd, os.ModePerm)
			if err1 != nil {
				return "", fmt.Errorf("failed to create standup file dir %w", err1)
			}
		} else {
			return "", fmt.Errorf("failed to stat standup file dir %w", err)
		}
	}
	return sfd, nil
}

func (c *Context) GetStartOfWeekDay() (wkday time.Weekday, err error) {
	var wkdaymap map[string]time.Weekday
	wkdaymap = make(map[string]time.Weekday)
	wkdaymap[time.Monday.String()] = time.Monday
	wkdaymap[time.Tuesday.String()] = time.Tuesday
	wkdaymap[time.Wednesday.String()] = time.Wednesday
	wkdaymap[time.Thursday.String()] = time.Thursday
	wkdaymap[time.Friday.String()] = time.Friday
	wkdaymap[time.Saturday.String()] = time.Saturday
	wkdaymap[time.Sunday.String()] = time.Sunday
	wkday, ok := wkdaymap[c.configuration.StartOfWeekDay]
	if !ok {
		return wkday, fmt.Errorf("invalid workday string %s", c.configuration.StartOfWeekDay)
	}
	return wkday, nil
}

func (c *Context) SetStartOfWeekDay(val string) error {
	c.configuration.StartOfWeekDay = val
	return c.configuration.WriteConfig()
}

func (c *Context) GetHolidays() []string {
	return c.configuration.Holidays
}

func (c *Context) SetHolidays(val []string) error {
	c.configuration.Holidays = val
	return c.configuration.WriteConfig()
}

func (c *Context) GetSections() []*ConfigSection {
	var sn []*ConfigSection
	sn = c.configuration.Sections
	return sn
}

func (c *Context) GetDefaultSection() string {
	return c.configuration.DefaultSection
}

func (c *Context) SetDefaultSection(val string) error {
	c.configuration.DefaultSection = val
	return c.configuration.WriteConfig()
}

func (c *Context) GetName() string {
	n := c.configuration.Name
	if n == "" {
		n = "UNKNOWN - PLEASE SET"
	}
	return n
}

func (c *Context) SetName(name string) error {
	c.configuration.Name = name
	return c.configuration.WriteConfig()
}

// SectionExists checks if specified section exists
func (c *Context) SectionExists(sectionName string) *ConfigSection {
	for _, s := range c.configuration.Sections {
		if s.Name == sectionName {
			return s
		}
	}
	return nil
}

func (c *Context) ExistsOrCreate(name string) error {
	if c.SectionExists(name) == nil {
		c.configuration.Sections = append(c.configuration.Sections, &ConfigSection{Name: name})
	}
	return c.configuration.WriteConfig()
}

func (c *Context) UpdateSectionName(providedName, name string) error {
	err := c.ExistsOrCreate(providedName)
	if err != nil {
		return err
	}
	c.SectionExists(providedName).Name = name
	return c.configuration.WriteConfig()
}

func (c *Context) UpdateSectionShortName(providedName, shortName string) error {
	err := c.ExistsOrCreate(providedName)
	if err != nil {
		return err
	}
	c.SectionExists(providedName).Short = shortName
	return c.configuration.WriteConfig()
}

func (c *Context) UpdateSectionDescription(providedName string, description string) error {
	err := c.ExistsOrCreate(providedName)
	if err != nil {
		return err
	}
	c.SectionExists(providedName).Description = description
	return c.configuration.WriteConfig()
}

func (c *Context) GetSectionDescription(name string) string {
	s := c.SectionExists(name)
	if s != nil {
		return s.Description
	}
	return ""
}

func (c *Context) GetSectionNameByShortName(shortName string) string {
	for _, s := range c.configuration.Sections {
		if s.Short == shortName {
			return s.Name
		}
	}
	return ""
}

func (c *Context) DeleteSection(name string) error {
	for i, s := range c.configuration.Sections {
		if s.Name == name {
			c.configuration.Sections = append(c.configuration.Sections[:i], c.configuration.Sections[i+1:]...)
			break
		}
	}
	return c.configuration.WriteConfig()
}

func (c *Context) GetSectionsPerRow() int {
	return c.configuration.SectionsPerRow
}

func (c *Context) SetSectionsPerRow(val int) error {
	c.configuration.SectionsPerRow = val
	return c.configuration.WriteConfig()
}
