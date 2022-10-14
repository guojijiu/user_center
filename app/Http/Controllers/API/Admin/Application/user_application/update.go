package user_application

import (
	"errors"
	"user_center/app/Http/Controllers/API/Admin/Context/User/UpdateUser"
	"user_center/app/Model"
	"user_center/app/Repository"
	"user_center/pkg/tool"
)

func Update(req *UpdateUser.Req) error {

	detail, err := Repository.UserRepository{}.Detail(req.ID)

	if err != nil {
		return err
	}
	if detail.ID == 0 {
		return errors.New("数据不存在或者已被删除。")
	}
	userInfoDetail, err := Repository.UserInfoRepository{}.FindByUserID(req.ID)
	if err != nil {
		return nil
	}
	var user Model.UserAuth
	var userInfo Model.UserInformation

	if req.Account != "" && req.Account != detail.Account {
		user.Account = req.Account
	}
	if req.Phone != "" && req.Phone != detail.Phone {
		user.Phone = req.Phone
	}
	if req.Email != "" && req.Email != detail.Email {
		user.Email = req.Email
	}
	if req.Nickname != "" && req.Nickname != userInfoDetail.Nickname {
		userInfo.Nickname = req.Nickname
	}
	if req.HeaderImgPath != "" && req.HeaderImgPath != userInfoDetail.HeaderImgPath {
		userInfo.HeaderImgPath = req.HeaderImgPath
	}
	if req.Sex != "" && req.Sex != userInfoDetail.Sex {
		userInfo.Sex = req.Sex
	}
	if req.Birthday != "" {
		birthday := tool.TimeStrToDatetime(req.Birthday)
		userInfo.Birthday = &birthday
	}
	if req.Address != "" && req.Address != userInfoDetail.Address {
		userInfo.Address = req.Address
	}
	if req.Organization != "" && req.Organization != userInfoDetail.Organization {
		userInfo.Organization = req.Organization
	}
	if req.PersonalProfile != "" && req.PersonalProfile != userInfoDetail.PersonalProfile {
		userInfo.PersonalProfile = req.PersonalProfile
	}

	return Repository.UserRepository{}.Update(&user, &userInfo, req.ID)
}
