package main

import (
	"flag"
	"log"
	"os"
	"user_center/boot"
	"user_center/command"
	"user_center/command/cmd_template"
	"user_center/pkg/tool"
)

func main() {
	flag.Usage = cmd_template.Usage
	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		cmd_template.Usage()
		os.Exit(2)
		return
	}

	if args[0] == "help" {
		cmd_template.Help(args[1:])
		return
	}

	if c, ok := command.CMD.Get(args[0]); ok && c.Run != nil {
		defer boot.Destroy()
		if c.Name() == "server" {
			boot.Boot()
		} else {
			boot.SetInConsole()
			boot.Boot()
		}
		c.Flag.Usage = func() { c.Usage() }
		if c.CustomFlags {
			args = args[1:]
		} else {
			err := c.Flag.Parse(args[1:])
			if err != nil {
				log.Fatal(err)
			}
			args = c.Flag.Args()
		}

		if c.PreRun != nil {
			c.PreRun(c, args)
		}

		os.Exit(c.Run(c, args))

		return
	}
	tool.PrintErrorTmplAndExit("Unknown command", cmd_template.ErrorTemplate)
}
