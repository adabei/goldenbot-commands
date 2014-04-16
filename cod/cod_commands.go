package cod

import (
	"errors"
	"github.com/schwarz/goldenbot/events"
	"github.com/schwarz/goldenbot/events/cod"
	"github.com/schwarz/goldenbot/rcon"
)

type Commands struct {
	commands map[string]func()
	events   chan interface{}
	requests chan rcon.RCONQuery
}

func NewCommands(requests chan rcon.RCONQuery, ea events.Aggregator) *Commands {
	c := new(Commands)
	c.requests = requests
	c.events = ea.Subscribe(c)
	return c
}

func (c *Commands) Setup() error {
	return nil
}

func (c *Commands) Start() {
	for {
		ev := <-c.events
		if ev, ok := in.(cod.Say); ok {
			if cmd, ok := c.commands[ev.Message]; ok {
				cmd()
			}
		}
	}
}

func (c *Commands) Register(cmd string, fn func()) error {
	if _, dup := c.commands[cmd]; dup {
		return errors.New("Command %v already defined", cmd)
	}

	c.commands[cmd] = fn
	return nil
}
