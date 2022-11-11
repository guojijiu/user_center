package department_application

import (
	"errors"
	"fmt"
	"user_center/app/Http/Controllers/API/Admin/Context/Department/UpdateDepartment"
	"user_center/app/Model"
	"user_center/app/Repository"
)

func Update(req *UpdateDepartment.Req) error {

	detail, err := Repository.DepartmentRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}

	var organize Model.Department

	if req.Name != "" && req.Name != detail.Name {
		organize.Name = req.Name
	}
	if req.Mark != "" && req.Mark != detail.Mark {
		organize.Mark = req.Mark
	}
	if req.Remark != "" && req.Remark != detail.Remark {
		detailByMark, err := Repository.DepartmentRepository{}.FindByMark(req.Mark)
		if err != nil {
			return err
		}
		if detailByMark.ID != 0 {
			return errors.New(fmt.Sprintf("唯一标识符：%s，已存在。", req.Mark))
		}

		organize.Remark = req.Remark
	}

	return Repository.DepartmentRepository{}.Update(&organize, req.ID)
}
