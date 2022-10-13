package client_application

import (
	"user_center/app/Http/Controllers/API/Admin/Context/Client/ListClient"
	"user_center/app/Repository"
)

func List(req *ListClient.Req) ([]ListClient.Resp, int, error) {

	list, total, err := Repository.ClientRepository{}.List(req)
	res := ListClient.Item(list)
	return res, total, err
}
