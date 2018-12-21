package controllers

import (
	"github.com/astaxie/beego/orm"
	"yistore/codes"
	"yistore/conf"
	"yistore/models"
	"yistore/objects"
)

type AddressController struct {
	DefaultController
}

//POST
func (this *AddressController) AddAddress() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_AddressAddOject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if recvObject.RecvName == "" || recvObject.RecvTelphone == "" || recvObject.Address == "" {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}

	address := &models.Address{User: &models.User{Id: id}, RecvName: recvObject.RecvName, RecvTelphone: recvObject.RecvTelphone, Address: recvObject.Address}
	code = this.insertAddress(address)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_AddAddressError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//GET
func (this *AddressController) GetAddress() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	addressId, err := this.GetInt32("address_id")
	if err != nil {
		page, err := this.GetInt("page")
		if err != nil {
			page = 0
		}
		var addresses []*models.Address
		code := this.getAddressesByPage(id, page, &addresses)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, codes.Code_GetAddressError)
			return
		}
		for _, address := range addresses {
			sendObject := &objects.S2C_AddressObject{Id: address.Id, RecvName: address.RecvName, RecvTelphone: address.RecvTelphone, Address: address.Address}
			jsonObject.Data = append(jsonObject.Data, sendObject)
		}
	} else {
		address := &models.Address{User: &models.User{Id: id}, Id: addressId}
		code := this.getAddressById(address)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, codes.Code_GetAddressError)
			return
		}
		sendObject := &objects.S2C_AddressObject{Id: address.Id, RecvName: address.RecvName, RecvTelphone: address.RecvTelphone, Address: address.Address}
		jsonObject.Data = append(jsonObject.Data, sendObject)
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *AddressController) DeleteAddress() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_AddressDeleteObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	address := &models.Address{User: &models.User{Id: id}, Id: recvObject.Id}
	code = this.deleteAddress(address)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DeleteAddressError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *AddressController) UpdateAddress() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_AddressUpdateOject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	address := &models.Address{User: &models.User{Id: id}, Id: recvObject.Id}

	code = this.getAddressById(address)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	cols := make([]string, 0)

	if recvObject.Address != "" {
		address.Address = recvObject.Address
		cols = append(cols, "Address")
	}
	if recvObject.RecvName != "" {
		address.RecvName = recvObject.RecvName
		cols = append(cols, "RecvName")
	}
	if recvObject.RecvTelphone != "" {
		address.RecvTelphone = recvObject.RecvTelphone
		cols = append(cols, "RecvTelphone")
	}

	code = this.updateAddress(address, cols)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_UpdateAddressError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *AddressController) insertAddress(address *models.Address) (code int32) {
	o := orm.NewOrm()
	if created, num, err := o.ReadOrCreate(address, "User", "RecvName", "RecvTelphone", "Address"); err != nil {
		code = 900
		return
	} else {
		if created {
			code = codes.Code_Right
			return
		}
		if num == 1 {
			if address.IsDeleted {
				address.IsDeleted = false
				code = this.updateAddress(address, []string{"IsDeleted"})
				return
			}
		}
		code = 900
		return
	}
}

func (this *AddressController) getAddressesByPage(userId int32, page int, addresses *[]*models.Address) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.Address)).Filter("User__Id", userId).Filter("IsDeleted", false).OrderBy("-Id").Limit(pageCount, offset).All(addresses)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *AddressController) getAddressById(address *models.Address) (code int32) {
	o := orm.NewOrm()
	err := o.QueryTable(new(models.Address)).Filter("User__Id", address.User.Id).Filter("Id", address.Id).Filter("IsDeleted", false).One(address)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *AddressController) updateAddress(address *models.Address, cols []string) (code int32) {
	o := orm.NewOrm()
	_, err := o.Update(address, cols...)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

//要假删
func (this *AddressController) deleteAddress(address *models.Address) (code int32) {
	code = this.getAddressById(address)
	if code != 0 {
		return
	}
	if address.IsUsed {
		address.IsDeleted = true
		code = this.updateAddress(address, []string{"IsDeleted"})
		return
	} else {
		o := orm.NewOrm()
		_, err := o.Delete(address, "User", "Id")
		if err != nil {
			code = 900
			return
		}
	}
	code = codes.Code_Right
	return
}
