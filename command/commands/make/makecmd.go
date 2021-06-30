package make

import (
	"user_center/command"
)

var CMDmigrator = &command.Command{
	UsageLine: "make [command]",
	Short:     "make 命令",
	Long:      usageDoc(),
	Run:       RunMigration,
}

func init() {
	command.CMD.Register(CMDmigrator)
}

func RunMigration(cmd *command.Command, args []string) int {
	usageDoc()
	return 0
}

func usageDoc() string {
	return `make:migration [name]`
}
