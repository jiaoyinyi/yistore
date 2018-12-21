package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"regexp"
	"strings"
	"yistore/codes"
	"yistore/models"
	"yistore/objects"
)

type UserController struct {
	DefaultController
}

//POST
func (this *UserController) RegisterPost() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//获取前端发过来的json数据
	recvObject := new(objects.C2S_UserRegisterObject)
	code := this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	//检测用户名、两次密码、手机号码是否符合要求
	code = checkUserName(recvObject.UserName)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = checkSamePassword(recvObject.Password, recvObject.Repassword)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = checkPassword(recvObject.Password)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = checkTelphone(recvObject.Telphone)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = this.checkUserNameUsed(recvObject.UserName)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = this.checkTelphoneUsed(recvObject.Telphone)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	user := &models.User{UserName: recvObject.UserName, Password: recvObject.Password, Telphone: recvObject.Telphone, Sex: "男", Description: "简介", Permission: 0}
	code = this.insertUser(user)
	objects.SetJsonObjectMsg(jsonObject, code)
}

//GET
func (this *UserController) LoginGet() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	_, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)

	username, ok := this.GetSession("username").(string)
	if !ok {
		objects.SetJsonObjectMsg(jsonObject, 900)
		return
	}
	sendObject := &objects.S2C_UserNameObject{UserName: username}
	jsonObject.Data = append(jsonObject.Data, sendObject)
	objects.SetJsonObjectMsg(jsonObject, code)
}

//POST
func (this *UserController) LoginPost() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	recvObject := new(objects.C2S_UserLoginObject)
	code := this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = checkUserName(recvObject.UserName)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = checkPassword(recvObject.Password)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	user := &models.User{UserName: recvObject.UserName, Password: recvObject.Password}
	code = this.getUserByUserNameAndPassword(user)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_UserLoginFailError)
		return
	}

	this.SetSession("id", user.Id)
	this.SetSession("username", user.UserName)
	this.SetSession("password", user.Password)
	this.SetSession("permission", user.Permission)
	logs.Info("设置Session成功")

	sendObject := &objects.S2C_UserNameObject{UserName: user.UserName}
	jsonObject.Data = append(jsonObject.Data, sendObject)
	objects.SetJsonObjectMsg(jsonObject, code)
}

//GET
func (this *UserController) Logout() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	this.DelSession("id")
	this.DelSession("username")
	this.DelSession("password")
	this.DelSession("permission")
	logs.Info("删除Session成功")

	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

func (this *UserController) UserGetMsg() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	user := &models.User{Id: id}
	code = this.getUserById(user)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	sendObject := &objects.S2C_UserObject{UserName: user.UserName, Password: user.Password, Telphone: user.Telphone, Sex: user.Sex, Description: user.Description}
	jsonObject.Data = append(jsonObject.Data, sendObject)
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

func (this *UserController) UserUpdateMsg() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_UserUpdateObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if recvObject.UserName == "" && recvObject.Password == "" && recvObject.Telphone == "" && recvObject.Description == "" {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}

	user := &models.User{Id: id}

	code = this.getUserById(user)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if recvObject.UserName != "" {
		code = checkUserName(recvObject.UserName)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		user.UserName = recvObject.UserName
	}

	if recvObject.Password != "" {
		code = checkPassword(recvObject.Password)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		user.Password = recvObject.Password
	}

	if recvObject.Telphone != "" {
		code = checkTelphone(recvObject.Telphone)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		user.Telphone = recvObject.Telphone
	}

	if recvObject.Sex != "" {
		code = checkSex(recvObject.Sex)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		user.Sex = recvObject.Sex
	}

	if recvObject.Description != "" {
		user.Description = recvObject.Description
	}

	code = this.updateUser(user)
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *UserController) CheckUserName() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	recvObject := new(objects.C2S_CheckUserNameObject)
	code := this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = this.checkUserNameUsed(recvObject.UserName)
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *UserController) CheckTelphone() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	recvObject := new(objects.C2S_CheckTelphoneObject)
	code := this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	code = this.checkTelphoneUsed(recvObject.Telphone)
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *UserController) checkUserNameUsed(userName string) (code int32) {
	user := &models.User{UserName: userName}

	o := orm.NewOrm()
	if err := o.Read(user, "UserName"); err == orm.ErrNoRows {
		code = codes.Code_Right
		return
	} else if err != nil {
		code = 900
		return
	}
	code = codes.Code_UsernameUsedError
	return
}

func (this *UserController) checkTelphoneUsed(telphone string) (code int32) {
	o := orm.NewOrm()
	user := &models.User{Telphone: telphone}
	if err := o.Read(user, "Telphone"); err == orm.ErrNoRows {
		code = codes.Code_Right
		return
	} else if err != nil {
		code = 900
		return
	}
	code = codes.Code_TelphoneUsedError
	return
}

//插入用户信息到数据库
func (this *UserController) insertUser(user *models.User) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Insert(user); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

//删除用户
func (this *UserController) deleteUser(user *models.User) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Delete(user, "Id"); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

//修改用户信息
func (this *UserController) updateUser(uesr *models.User) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Update(uesr); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

//通过用户ID获取用户信息
func (this *UserController) getUserById(user *models.User) (code int32) {
	o := orm.NewOrm()
	if err := o.Read(user, "Id"); err != nil {
		code = codes.Code_UserNoExistError
		return
	}
	code = codes.Code_Right
	return
}

func (this *UserController) getUserByUserNameAndPassword(user *models.User) (code int32) {
	o := orm.NewOrm()
	if err := o.Read(user, "UserName", "Password"); err != nil {
		code = codes.Code_UserNoExistError
		return
	}
	code = codes.Code_Right
	return
}

//第一页是0
func (this *UserController) getUsersByPage(page int, users *[]*models.User) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := 20
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.User)).OrderBy("-Id").Limit(pageCount, offset).All(users); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

//获取所有用户
func (this *UserController) getAllUsers(users *[]*models.User) (code int32) {
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.User)).All(users); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

//////////////////////////////////////////////////////////////////////////////
//验证用户名
func checkUserName(userName string) (code int32) {
	if userName == "" {
		code = codes.Code_UsernameEmptyError
		return
	} else {
		isOk, _ := regexp.MatchString("^[a-zA-Z0-9]{0,14}$", userName)
		if !isOk {
			code = codes.Code_UsernameFormatError
			return
		}
		code = codes.Code_Right
		return
	}
}

//验证密码
func checkPassword(password string) (code int32) {
	if password == "" {
		code = codes.Code_PasswordEmptyError
		return
	} else {
		isOk, _ := regexp.MatchString("^[a-zA-Z0-9_]{8,16}$", password)
		if !isOk {
			code = codes.Code_PasswordFormatError
			return
		}
		code = codes.Code_Right
		return
	}
}

func checkSamePassword(password string, repassword string) (code int32) {
	if password == repassword {
		code = codes.Code_Right
		return
	}
	code = codes.Code_PasswordNoSameError
	return
}

//验证手机号码
func checkTelphone(telphone string) (code int32) {
	if telphone == "" {
		code = codes.Code_TelphoneEmptyError
		return
	} else {
		isOk, _ := regexp.MatchString("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$", telphone)
		if !isOk {
			code = codes.Code_TelphoneFormatError
			return
		}
		code = codes.Code_Right
		return
	}
}

func checkSex(sex string) (code int32) {
	if sex == "" {
		code = codes.Code_SexEmptyError
		return
	}

	if strings.Compare(sex, "男") == 0 || strings.Compare(sex, "女") == 0 {
		code = codes.Code_SexFormatError
		return
	}
	code = codes.Code_Right
	return
}
