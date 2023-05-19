package cmd_template

import (
	"user_center/command"
	_ "user_center/command/commands/app_key_generator"
	_ "user_center/command/commands/app_rsakey_generator"
	_ "user_center/command/commands/make"
	_ "user_center/command/commands/make_migration"
	_ "user_center/command/commands/migrator"
	_ "user_center/command/commands/server"
	_ "user_center/command/commands/test"
	"user_center/pkg/tool"
)

var usageTemplate = `user_center is a user_repository and authority management system based on THE GIN framework.

{{"用法(注意：如果user_center可执行文件不在系统PATH中，请使用<user_center所在的目录>/user_center)" | headline}}
    {{"user_center command [arguments]" | bold}}

{{"可使用的命令" | headline}}
{{range .}}{{if .Runnable}}
    {{.Name | printf "%-75s" | greenbold }} {{.Short | green }}{{ end }}{{ end }}

Use {{"user_center help [command]" | bold}} for more information about a command.

{{"额外的帮助" | headline}}
{{range .}}{{if not .Runnable}}
    {{.Name | printf "%-11s"}} {{.Short | green }}{{end}}{{end}}

Use {{"user_center help [topic]" | bold}} for more information about that topic.
`

var helpTemplate = `{{"用法" | headline}}
  {{.UsageLine | printf "user_center %s" | bold}}
{{if .Options}}{{endline}}{{"OPTIONS" | headline}}{{range $k,$v := .Options}}
  {{$k | printf "-%s" | bold}}
      {{$v}}
  {{end}}{{end}}
{{"DESCRIPTION" | headline}}
  {{tmpltostr .Long . | trim}}
`

var ErrorTemplate = `user_center: %s.
Use {{"user_center help" | bold}} for more information.
`

func Usage() {
	tool.TmplTextParseAndOutput(usageTemplate, command.CMD.Commands)
}

func Help(args []string) {
	if len(args) == 0 {
		Usage()
	}
	if len(args) != 1 {
		tool.PrintErrorTmplAndExit("Too many arguments", ErrorTemplate)
	}

	arg := args[0]

	if cmd, ok := command.CMD.Get(arg); ok {
		tool.TmplTextParseAndOutput(helpTemplate, cmd)
	} else {
		tool.PrintErrorTmplAndExit("Unknown help topic", ErrorTemplate)
	}
}
