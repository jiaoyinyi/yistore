package controllers

import (
	"github.com/astaxie/beego/orm"
	"yistore/codes"
	"yistore/conf"
	"yistore/models"
	"yistore/objects"
)

type ShopCartController struct {
	DefaultController
}

func (this *ShopCartController) ShopCartGet() {
	//从数据库中查询购物车
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if shopCartId, err := this.GetInt32("shopcart_id"); err != nil {
		var shopcarts []*models.ShopCart
		page, err := this.GetInt("page")
		if err != nil {
			page = 0
		}

		code = this.getShopCartsByPage(id, page, &shopcarts)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, codes.Code_GetShopcartError)
			return

		}
		for _, shopcart := range shopcarts {
			var images []*models.Image
			var paths []string
			(&ImageController{}).getImagesByCommodityId(shopcart.CommodityBase.Id, &images)
			objects.SetImage(&paths, images)
			commodity := &objects.S2C_CommodityBaseObject{Id: shopcart.CommodityBase.Id, Name: shopcart.CommodityBase.Name, Status: shopcart.CommodityBase.Status, Price: shopcart.CommodityBase.Price, Images: paths}
			sendObject := &objects.S2C_ShopCartObject{Id: shopcart.Id, Count: shopcart.Count, Commodity: commodity}
			jsonObject.Data = append(jsonObject.Data, sendObject)
		}
	} else {
		shopcart := &models.ShopCart{User: &models.User{Id: id}, Id: shopCartId}
		code = this.getShopCartById(shopcart)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, codes.Code_GetShopcartError)
			return
		}
		var images []*models.Image
		var paths []string
		(&ImageController{}).getImagesByCommodityId(shopcart.CommodityBase.Id, &images)
		objects.SetImage(&paths, images)
		commodity := &objects.S2C_CommodityBaseObject{Id: shopcart.CommodityBase.Id, Name: shopcart.CommodityBase.Name, Status: shopcart.CommodityBase.Status, Price: shopcart.CommodityBase.Price, Images: paths}
		sendObject := &objects.S2C_ShopCartObject{Id: shopcart.Id, Count: shopcart.Count, Commodity: commodity}
		jsonObject.Data = append(jsonObject.Data, sendObject)
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *ShopCartController) ShopCartAdd() {
	//添加物品到购物车数据库
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_ShopCartAddObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}
	if recvObject.Count <= 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}

	shopcart := &models.ShopCart{User: &models.User{Id: id}, CommodityBase: &models.CommodityBase{Id: recvObject.CommodityId}, Count: recvObject.Count}
	code = this.insertShopCart(shopcart)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_AddShopcartError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *ShopCartController) ShopCartUpdate() {
	//更新物品到购物车数据库
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_ShopCartUpdateObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}
	if recvObject.Count <= 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}

	shopcart := &models.ShopCart{User: &models.User{Id: id}, Id: recvObject.Id, Count: recvObject.Count}
	code = this.updateShopCart(shopcart, []string{"Count"})
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_UpdateShopcartError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *ShopCartController) ShopCartDelete() {
	//从购物车数据库中删除物品
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_ShopCartDeleteObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	shopcart := &models.ShopCart{User: &models.User{Id: id}, Id: recvObject.Id}
	code = this.deleteShopCart(shopcart)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DeleteShopcartError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//func (this *ShopCartController) getAllShopCarts(userId int32, shopcarts *[]*models.ShopCart) (code int32) {
//	o := orm.NewOrm()
//	_, err := o.QueryTable(new(models.ShopCart)).Filter("User__Id", userId).RelatedSel().OrderBy("-Id").All(shopcarts)
//	if err != nil {
//		code = 900
//		return
//	}
//	code = codes.Code_Right
//	return
//}

func (this *ShopCartController) getShopCartById(shopcart *models.ShopCart) (code int32) {
	o := orm.NewOrm()

	err := o.QueryTable(new(models.ShopCart)).Filter("User__Id", shopcart.User.Id).Filter("Id", shopcart.Id).RelatedSel().One(shopcart)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *ShopCartController) getShopCartsByPage(userId int32, page int, shopcarts *[]*models.ShopCart) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.ShopCart)).Filter("User__Id", userId).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(shopcarts); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *ShopCartController) insertShopCart(shopcart *models.ShopCart) (code int32) {
	count := shopcart.Count
	o := orm.NewOrm()
	if created, _, err := o.ReadOrCreate(shopcart, "User", "CommodityBase"); err != nil {
		code = 900
		return
	} else {
		if !created {
			shopcart.Count += count
			code = this.updateShopCart(shopcart, []string{"Count"})
			if code != 0 {
				return
			}
		}
	}
	code = codes.Code_Right
	return
}

func (this *ShopCartController) deleteShopCart(shopcart *models.ShopCart) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Delete(shopcart); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *ShopCartController) updateShopCart(shopcart *models.ShopCart, cols []string) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Update(shopcart, cols...); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}
