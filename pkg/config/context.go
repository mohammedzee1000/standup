package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

//Context repersents the run context
type Context struct {
	configuration *Config
	dataDir       string
}

//New returns a context
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

//DataDir gets the data dir
func (c *Context) DataDir() string {
	return c.dataDir
}

//GetStandUpFilePath returns path to use for standup record file based on day, month and year
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

//SectionExists checks if specified section exists
func (c *Context) SectionExists(sectioName string) bool {
	for s, _ := range c.configuration.SectionNames {
		if s == sectioName {
			return true
		}
	}
	return false
}

func (c *Context) GetSectionDescription(sectionName string) string {
	for k, v := range c.configuration.SectionNames {
		if k == sectionName {
			return v
		}
	}
	return ""
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

func (c *Context) GetHolidays() []string {
	return c.configuration.Holidays
}
