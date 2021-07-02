package make_migration

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"text/template"
	"time"
	"user_center/app"
	"user_center/command"
	"user_center/pkg/tool"
)

var CMDmigrator = &command.Command{
	UsageLine: "make:migration [name]",
	Short:     "创建迁移文件",
	Long:      usageDoc(),
	Run:       RunMigration,
}

func init() {
	command.CMD.Register(CMDmigrator)
}

func RunMigration(cmd *command.Command, args []string) int {
	name := args[0]
	fileName := time.Now().Format("20060102_150405_") + tool.Camel2Case(name) + ".go"
	p := path.Join(app.DatabasePath, "MigrationFile", fileName)

	if tool.PathExist(p) {
		cmd.Error(fmt.Sprintf("file %s aleardy exist", p))
		return 1
	}
	file, err := os.Create(p)
	defer func() { _ = file.Close() }()
	if err != nil {
		cmd.Errorf("创建文件失败: %s", err.Error())
		return 1
	}
	temp, err := Template(name, fileName)
	if err != nil {
		cmd.Errorf("生成模板文件失败: %s", err.Error())
		return 1
	}
	_, err = file.WriteString(temp)
	if err != nil {
		cmd.Errorf("向文件中写数据失败: %s", err.Error())
		return 1
	}
	return 0
}

func usageDoc() string {
	return fmt.Sprintf(`
创建迁移文件到文件夹 database/MigrationFile
make:migration [name]
`)
}

func Template(structName string, fileName string) (string, error) {
	t := template.New("t")
	t, err := t.Parse(`
package MigrationFile

import (
	"fmt"
	"user_center/app/Model"
	"user_center/pkg/db"
)

type {{.struct_name}} struct {
}

func ({{.struct_name}}) Key() string {
	return "{{.file_name}}"
}

func ({{.struct_name}}) Up() (err error) {
	return
}

func ({{.struct_name}}) Down() (err error) {
	return
}
`)
	if err != nil {
		return "", err
	}
	out := bytes.Buffer{}
	_ = t.Execute(&out, map[string]string{
		"struct_name": structName,
		"file_name":   fileName,
	})
	return out.String(), nil
}
