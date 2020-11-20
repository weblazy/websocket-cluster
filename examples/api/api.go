package api

import (
	"websocket-cluster/examples/domain"

	"github.com/labstack/echo/v4"
)

// @desc 登录
// @auth liuguoqiang 2020-11-20
// @param
// @return
func Login(c echo.Context) error {
	//参数验证绑定
	req, response, err := ParseJson(c)
	if err != nil {
		return response.RetError(err, -1)
	}
	username := req.Get("username").String()
	password := req.Get("password").String()
	resp, err := domain.Login(username, password)
	if err != nil {
		return response.RetError(err, -1)
	}
	return response.RetCustomize(0, resp, "")
}

// @desc 聊天初始化
// @auth liuguoqiang 2020-11-20
// @param
// @return
func ChatInit(c echo.Context) error {
	id, _, response, err := ParseParams(c)
	if err != nil {
		return response.RetError(err, -1)
	}
	resp, err := domain.ChatInit(id)
	if err != nil {
		return response.RetError(err, -1)
	}
	return response.RetCustomize(0, resp, "")
}

// @desc 获取群成员
// @auth liuguoqiang 2020-11-20
// @param
// @return
func GetGroupMembers(c echo.Context) error {
	id, _, response, err := ParseParams(c)
	if err != nil {
		return response.RetError(err, -1)
	}
	resp, err := domain.GetGroupMembers(id)
	if err != nil {
		return response.RetError(err, -1)
	}
	return response.RetCustomize(0, resp, "")
}
