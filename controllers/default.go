package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"yistore/codes"
	"yistore/objects"
)

type DefaultController struct {
	beego.Controller
}

//检查是否是用户登录
func (this *DefaultController) checkUserLogined() (id int32, code int32) {
	id, ok := this.GetSession("id").(int32)
	if !ok {
		code = codes.Code_NoLoginError
		return
	}
	code = codes.Code_Right
	return
}

//检查是否是管理员登录
func (this *DefaultController) checkAdminLogined() (id int32, code int32) {
	id, ok := this.GetSession("id").(int32)
	if !ok {
		code = codes.Code_NoLoginError
		return
	} else {
		if permission, ok := this.GetSession("permission").(int32); !ok {
			code = codes.Code_NoLoginError
			return
		} else {
			if permission == 1 {
				code = codes.Code_Right
				return
			}
			code = codes.Code_PermissionError
			return
		}
	}
}

//获取数据
func (this *DefaultController) getJsonData(object interface{}) (code int32) {
	err := json.Unmarshal(this.Ctx.Input.RequestBody, object)
	if err == nil {
		return codes.Code_Right
	}
	return codes.Code_DataFormatError
}

//设置json数据返回给前端
func (this *DefaultController) sendJsonObject(object *objects.JsonObject) {
	this.Data["json"] = object
	this.ServeJSON()
}
