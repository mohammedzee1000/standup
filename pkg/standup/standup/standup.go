package standup

import (
	"time"

	"github.com/mohammedzee1000/standup/pkg/standup/task"
)

type Section []*task.Task

type StandUp struct {
	Date      time.Time          `json:"Date"`
	Sections  map[string]Section `json:"Sections"`
	IsHoliday bool               `json:"IsHoliday"`
}

func NewEmptyStandUp() *StandUp {
	return &StandUp{
		Sections: make(map[string]Section),
	}
}

func NewStandUpOnDate(t time.Time) *StandUp {
	s := NewEmptyStandUp()
	s.Date = t
	return s
}

func (s *StandUp) GetSectionNames() (sn []string) {
	for k, _ := range s.Sections {
		sn = append(sn, k)
	}
	return sn
}
