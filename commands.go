package main

import (
	"errors"

	"github.com/aott33/gator/internal/config"
	"github.com/aott33/gator/internal/database"
)

type state struct {
	cfg		*config.Config
	db		*database.Queries
}

type command struct {
	name	string
	args	[]string	
}


type commands struct {
	m		map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.m[cmd.name]

	if !ok {
		return errors.New("Could not find command")
	}

	err := value(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.m == nil {
		c.m = make(map[string]func(*state, command) error)
	}

	c.m[name] = f	
}
