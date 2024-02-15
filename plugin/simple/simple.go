package simple

import (
	"flag"
	"fmt"
)

type simplePlugin struct {
	name  string
	value string
}

func NewSimplePlugin(name string) *simplePlugin {
	return &simplePlugin{
		name: name,
	}
}

func (s *simplePlugin) GetPrefix() string {
	return s.name
}

func (s *simplePlugin) Get() interface{} {
	return s
}

func (s *simplePlugin) Name() string {
	return s.name
}

func (s *simplePlugin) InitFlags() {
	flag.StringVar(&s.value, fmt.Sprintf("%s-value", s.name), "default value", "Simple plugin value")
}

func (s *simplePlugin) Configure() error {
	return nil
}

func (s *simplePlugin) Run() error {
	return nil
}

func (s *simplePlugin) Stop() <-chan bool {

	c := make(chan bool)
	go func() {
		c <- true
	}()

	return c
}

func (s *simplePlugin) GetValue() string {
	return s.value
}
