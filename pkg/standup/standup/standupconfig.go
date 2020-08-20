package standup

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/mohammedzee1000/standup/pkg/config"
	"github.com/mohammedzee1000/standup/pkg/system"
)

type StandUpConfig struct {
	configFile string
	standup    *StandUp
}

func NewEmptyStandUpConfig() *StandUpConfig {
	return &StandUpConfig{
		standup: NewEmptyStandUp(),
	}
}

func NewStandUpConfig(t time.Time) *StandUpConfig {
	s := NewEmptyStandUpConfig()
	s.configFile = fmt.Sprintf("%dd-%sm-%dy", t.Day(), t.Month().String(), t.Year())
	return s
}

func (sc *StandUpConfig) ConfigFileExists(c *config.Context) (exists bool, err error) {
	cfp, err := sc.getConfigFilePath(c)
	if err != nil {
		return false, err
	}
	_, err = os.Stat(cfp)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}

func (sc *StandUpConfig) getConfigFilePath(c *config.Context) (string, error) {
	sfd, err := c.GetStandUpFileDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(sfd, sc.configFile), nil

}

func (sc *StandUpConfig) FromFile(c *config.Context) error {
	cfp, err := sc.getConfigFilePath(c)
	if err != nil {
		return err
	}
	_, err = system.ReadYamlFile(cfp, sc.standup)
	return err
}

func (sc *StandUpConfig) ToFile(c *config.Context) error {
	cfp, err := sc.getConfigFilePath(c)
	if err != nil {
		return err
	}

	return system.WriteYamlFile(cfp, sc.standup)
}

func (sc *StandUpConfig) GetStandUp() *StandUp {
	return sc.standup
}
