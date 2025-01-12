package commands

import (
	"fmt"
	"strings"

	"github.com/ithinkiborkedit/GUSH.git/pkg/game"
	"github.com/ithinkiborkedit/GUSH.git/pkg/models"
)

type CommandFunc func(args []string, usr *models.User, world *game.World) string

type Dispatcher struct {
	commands map[string]CommandFunc
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		commands: make(map[string]CommandFunc),
	}

}

func (d *Dispatcher) Register(cmd string, fn CommandFunc) {
	d.commands[strings.ToLower(cmd)] = fn
}

func (d *Dispatcher) Dispatch(input string, usr *models.User, world *game.World) string {
	tokens := strings.Fields(input)
	if len(tokens) == 0 {
		return "No command entered"
	}

	cmd := strings.ToLower(tokens[0])
	args := tokens[1:]

	if fn, exists := d.commands[cmd]; exists {
		return fn(args, usr, world)
	}

	fmt.Println(input)

	return "unknown command."
}
