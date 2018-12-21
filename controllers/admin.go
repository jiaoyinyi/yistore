package controllers

import (
	"github.com/astaxie/beego/logs"
	"yistore/codes"
	"yistore/models"
	"yistore/objects"
)

type AdminController struct {
	DefaultController
}

//POST
func (this *AdminController) LoginPost() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	recvObject := new(objects.C2S_UserLoginObject)
	code := this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	user := &models.User{UserName: recvObject.UserName, Password: recvObject.Password}

	code = (&UserController{}).getUserByUserNameAndPassword(user)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if user.Permission == 1 {
		this.SetSession("id", user.Id)
		this.SetSession("username", user.UserName)
		this.SetSession("password", user.Password)
		this.SetSession("permission", user.Permission)
		logs.Info("设置Session成功")

		sendObject := &objects.S2C_UserNameObject{UserName: user.UserName}
		jsonObject.Data = append(jsonObject.Data, sendObject)
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_AdminLoginFailError)
}

//GET
func (this *AdminController) Logout() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	this.DelSession("id")
	this.DelSession("username")
	this.DelSession("password")
	this.DelSession("permission")
	logs.Info("删除Session成功")

	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}
