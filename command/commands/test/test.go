package test

import (
	"user_center/command"
)

var CMDtest = &command.Command{
	UsageLine: "test",
	Short:     "test 测试命令",
	Long:      testDoc(),
	Run:       RunTest,
}

func init() {
	command.CMD.Register(CMDtest)
}

func RunTest(cmd *command.Command, args []string) int {
	testDoc()
	return 0
}

func testDoc() string {
	return `测试`
}
