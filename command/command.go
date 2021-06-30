package command

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"user_center/pkg/color"
	"user_center/pkg/tool"
)

// Command is the unit of execution
type Command struct {
	// Run runs the command.
	// The args are the arguments after the command name.
	Run func(cmd *Command, args []string) int

	// PreRun performs an operation before running the command
	PreRun func(cmd *Command, args []string)

	// UsageLine is the one-line Usage message.
	// The first word in the line is taken to be the command name.
	UsageLine string

	// Short is the short description shown in the 'go help' output.
	Short string

	// Long is the long message shown in the 'go help <this-command>' output.
	Long string

	// Flag is a set of flags specific to this command.
	Flag flag.FlagSet

	// CustomFlags indicates that the command will do its own
	// flag parsing.
	CustomFlags bool

	// output out writer if set in SetOutput(w)
	output *io.Writer
}

type Args []string

// available commands in uims system
type CMDManager struct {
	Commands []*Command
}

func (cmd *CMDManager) Register(c *Command) {
	cmd.Commands = append(cmd.Commands, c)
}

func (cmd *CMDManager) Call(call *Command, args Args) {
	if cmd.Exist(call) {
		call.Run(call, args)
	}
}

func (cmd *CMDManager) Exist(call *Command) bool {
	for _, availableCommand := range cmd.Commands {
		if availableCommand.UsageLine == call.UsageLine {
			return true
		}
	}
	return false
}

func (cmd *CMDManager) ExistName(name string) bool {
	for _, availableCommand := range cmd.Commands {
		if availableCommand.Name() == name {
			return true
		}
	}
	return false
}

func (cmd *CMDManager) Get(name string) (*Command, bool) {
	for _, availableCommand := range cmd.Commands {
		if availableCommand.Name() == name {
			return availableCommand, true
		}
	}
	return nil, false
}

var CMD CMDManager

var cmdUsage = `Use {{printf "uims help %s" .Name | bold}} for more information.{{endline}}`

// Name returns the command's name: the first word in the Usage line.
func (c *Command) Name() string {
	name := c.UsageLine
	i := strings.Index(name, " ")
	if i >= 0 {
		name = name[:i]
	}
	return name
}

// SetOutput sets the destination for Usage and error messages.
// If output is nil, os.Stderr is used.
func (c *Command) SetOutput(output io.Writer) {
	c.output = &output
}

// Out returns the out writer of the current command.
// If cmd.output is nil, os.Stderr is used.
func (c *Command) Out() io.Writer {
	if c.output != nil {
		return *c.output
	}
	return color.NewColorWriter(os.Stderr)
}

// Usage puts out the Usage for the command.
func (c *Command) Usage() {
	tool.TmplTextParseAndOutput(cmdUsage, c)
	os.Exit(2)
}

// Runnable reports whether the command can be run; otherwise
// it is a documentation pseudo-command such as import path.
func (c *Command) Runnable() bool {
	return c.Run != nil
}

// Options return all options of command.
func (c *Command) Options() map[string]string {
	options := make(map[string]string)
	c.Flag.VisitAll(func(f *flag.Flag) {
		defaultVal := f.DefValue
		if len(defaultVal) > 0 {
			options[f.Name+"="+defaultVal] = f.Usage
		} else {
			options[f.Name] = f.Usage
		}
	})
	return options
}

func (c *Command) formatMessage(message string) string {
	return fmt.Sprintf("[%s] %s \n", time.Now().Format("2006/01/02 15:04:05"), message)
}

func (c *Command) Info(message string) {
	_, _ = c.Out().Write([]byte(color.Green(c.formatMessage(message))))
}

func (c *Command) Warn(message string) {
	_, _ = c.Out().Write([]byte(color.Cyan(c.formatMessage(message))))
}

func (c *Command) Error(message string) {
	_, _ = c.Out().Write([]byte(color.Red(c.formatMessage(message))))
}

func (c *Command) Errorf(message string, a interface{}) {
	_, _ = c.Out().Write([]byte(color.Red(c.formatMessage(fmt.Sprintf(message, a)))))
}
