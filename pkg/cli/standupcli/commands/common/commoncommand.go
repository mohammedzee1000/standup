package common

import (
	"github.com/mohammedzee1000/standup/pkg/config"
)

type CommonOptions struct {
	Context *config.Context
}

func NewCommonOptions() *CommonOptions {
	return &CommonOptions{}
}

func (cc *CommonOptions) InitContext() (err error) {
	c, err := config.New()
	if err != nil {
		return err
	}
	cc.Context = c
	return nil
}
