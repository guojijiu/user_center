package user_application

import (
	"fmt"
	"time"
	"user_center/app"
	"user_center/app/Http/Controllers/API/Admin/Context/User/ExportUser"
	"user_center/app/Repository"
	"user_center/pkg/file"
	"user_center/pkg/tool"
)

func ExportUserData(req *ExportUser.Req) (string, error) {

	userData, err := Repository.UserRepository{}.GetDataByExport(req)
	if err != nil {
		return "", err
	}

	result := ExportUser.Result(userData)
	filePath := app.GetStoragePath(fmt.Sprintf("app/public/file/tmp/user/export/%s/%s/aaa.csv", time.Now().Format("20060102"), tool.GenerateRandStrWithMath(16)))
	if err = file.ExportToCSV(filePath, result); err != nil {
		return "", err
	}
	return filePath, nil
}
