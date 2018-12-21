package controllers

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yistore/codes"
	"yistore/conf"
	"yistore/models"
	"yistore/objects"
)

type OrderController struct {
	DefaultController
}

/*
	订单状态status:
	0:用户下单
	1:管理员发货确认
	2:用户收货确认订单完成
*/

func (this *OrderController) UserGetOrder() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	orderId, err := this.GetInt32("order_id");
	if err == nil {
		order := &models.Order{Id: orderId}
		code = this.getOrderById(order)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		var images []*models.Image
		var paths []string
		(&ImageController{}).getImagesByCommodityId(order.CommodityBase.Id, &images)
		objects.SetImage(&paths, images)
		commodityBaseObject := &objects.S2C_CommodityBaseObject{Id: order.CommodityBase.Id, Name: order.CommodityBase.Name, Price: order.CommodityBase.Price, Status: order.CommodityBase.Status, Images: paths}
		addressObject := &objects.S2C_AddressObject{Id: order.RecvAddress.Id, RecvName: order.RecvAddress.RecvName, RecvTelphone: order.RecvAddress.RecvTelphone, Address: order.RecvAddress.Address}
		sendObject := &objects.S2C_OrderObject{Id: order.Id, Status: order.Status, CreateTime: order.CreateTime, FinishTime: order.FinishTime, Count: order.Count, UnitPrice: order.UnitPrice, Commodity: commodityBaseObject, RecvAddress: addressObject}
		jsonObject.Data = append(jsonObject.Data, sendObject)
		objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
		return
	}

	var orders []*models.Order
	page, err := this.GetInt("page")
	if err != nil {
		page = 0
	}
	status, err := this.GetInt32("status")
	if err != nil {
		code = this.getOrdersByUserId(id, page, &orders)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
	} else {
		if status < 0 || status > 2 {
			status = 2
		}
		code = this.getOrdersByUserIdAndStatus(id, status, page, &orders)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
	}
	for _, order := range orders {
		var images []*models.Image
		var paths []string
		(&ImageController{}).getImagesByCommodityId(order.CommodityBase.Id, &images)
		objects.SetImage(&paths, images)
		commodityBaseObject := &objects.S2C_CommodityBaseObject{Id: order.CommodityBase.Id, Name: order.CommodityBase.Name, Price: order.CommodityBase.Price, Status: order.CommodityBase.Status, Images: paths}
		addressObject := &objects.S2C_AddressObject{Id: order.RecvAddress.Id, RecvName: order.RecvAddress.RecvName, RecvTelphone: order.RecvAddress.RecvTelphone, Address: order.RecvAddress.Address}
		sendObject := &objects.S2C_OrderObject{Id: order.Id, Status: order.Status, CreateTime: order.CreateTime, FinishTime: order.FinishTime, Count: order.Count, UnitPrice: order.UnitPrice, Commodity: commodityBaseObject, RecvAddress: addressObject}
		jsonObject.Data = append(jsonObject.Data, sendObject)
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *OrderController) DeleteOrder() {
	//当订单状态status为2的时候，用户才能删除订单
	//用户只能删除完成后的订单
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_OrderDeleteObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	order := &models.Order{User: &models.User{Id: id}, Id: recvObject.Id}
	code = this.deleteOrder(order)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *OrderController) UserComfirmOrder() {
	//只能修改订单状态
	//当订单状态status为1的时候，用户才能修改状态
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_OrderUserUpdateObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	order := &models.Order{User: &models.User{Id: id}, Id: recvObject.Id}
	code = this.getOrderById(order)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if order.Status != 1 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_UserComfirmOrderError)
		return
	}
	order.Status = 2
	order.FinishTime = time.Now()
	code = this.updateOrder(order, []string{"Status", "FinishTime"})
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_UserComfirmOrderError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

func (this *OrderController) AdminComfirmOrder() {
	//只能修改订单状态
	//当订单状态status为0的时候，管理员才能修改状态
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	_, code := this.checkAdminLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_OrderAdminUpdateObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	order := &models.Order{User: &models.User{Id: recvObject.UserId}, Id: recvObject.Id}
	code = this.getOrderById(order)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if order.Status != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_AdminComfirmOrderError)
	}
	order.Status = 1
	code = this.updateOrder(order, []string{"Status"})
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_AdminComfirmOrderError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

func (this *OrderController) AdminCancelOrder() {
	//只能修改订单状态
	//当订单状态status为0的时候，管理员才能修改状态
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	_, code := this.checkAdminLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_OrderAdminUpdateObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	order := &models.Order{User: &models.User{Id: recvObject.UserId}, Id: recvObject.Id}
	code = this.getOrderById(order)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if order.Status != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_AdminCancelOrderError)
	}
	code = this.deleteOrder(order)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_AdminCancelOrderError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *OrderController) AddOrder() {
	//添加订单
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_OrderAddObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}
	if recvObject.AddressId <= 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}

	if recvObject.ShopCartId > 0 && (recvObject.CommodityId <= 0 || recvObject.Count <= 0) {
		//购物车下单
		code = this.addOrderByShopCart(id, recvObject.ShopCartId, recvObject.AddressId)
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	} else if recvObject.ShopCartId <= 0 && (recvObject.CommodityId > 0 || recvObject.Count > 0) {
		//直接下单
		code = this.addOrderByCommodityAndCount(id, recvObject.CommodityId, recvObject.Count, recvObject.AddressId)
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	} else {
		//错误
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}
}

func (this *OrderController) getOrdersByUserId(userId int32, page int, orders *[]*models.Order) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.Order)).Filter("User__Id", userId).Filter("IsDeleted", false).OrderBy("-CreateTime").Limit(pageCount, offset).RelatedSel().All(orders)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *OrderController) getOrdersByUserIdAndStatus(userId int32, status int32, page int, orders *[]*models.Order) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.Order)).Filter("User__Id", userId).Filter("Status", status).Filter("IsDeleted", false).OrderBy("-CreateTime").Limit(pageCount, offset).RelatedSel().All(orders)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *OrderController) getOrderById(order *models.Order) (code int32) {
	o := orm.NewOrm()
	err := o.QueryTable(new(models.Order)).Filter("User__Id", order.User.Id).Filter("Id", order.Id).Filter("IsDeleted", false).RelatedSel().One(order)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *OrderController) getOrdersByCommodityId(commodityId int32, page int, orders *[]*models.Order) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.Order)).Filter("CommodityBase__Id", commodityId).Filter("IsDeleted", false).OrderBy("-CreateTime").Limit(pageCount, offset).RelatedSel().All(orders)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *OrderController) deleteOrder(order *models.Order) (code int32) {
	code = this.getOrderById(order)
	if code != 0 {
		return
	}
	if order.Status == 2 {
		order.IsDeleted = true
		code = this.updateOrder(order, []string{"IsDeleted"})
		return
	}
	code = 900
	return
}

func (this *OrderController) updateOrder(order *models.Order, cols []string) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Update(order, cols...); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *OrderController) addOrderByShopCart(userId int32, shopcartId int32, addressId int32) (code int32) {
	shopcart := &models.ShopCart{User: &models.User{Id: userId}, Id: shopcartId}
	code = (&ShopCartController{}).getShopCartById(shopcart)
	if code != 0 {
		return
	}
	//商品下架
	if !shopcart.CommodityBase.Status {
		return
	}

	address := &models.Address{User: &models.User{Id: userId}, Id: addressId}
	code = (&AddressController{}).getAddressById(address)
	if code != 0 {
		return
	}

	order := &models.Order{User: shopcart.User, CommodityBase: shopcart.CommodityBase, Count: shopcart.Count, UnitPrice: shopcart.CommodityBase.Price, RecvAddress: address, CreateTime: time.Now()}
	o := orm.NewOrm()
	if _, err := o.Insert(order); err != nil {
		address.IsUsed = true
		(&AddressController{}).updateAddress(address, []string{"IsUsed"})
		shopcart.CommodityBase.IsUsed = true
		(&CommodityController{}).updateCommodity(shopcart.CommodityBase, []string{"IsUsed"})
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *OrderController) addOrderByCommodityAndCount(userId int32, commodityId int32, count int32, addressId int32) (code int32) {
	commodity := &models.CommodityBase{Id: commodityId}
	code = (&CommodityController{}).getCommodityBaseById(commodity)
	if code != 0 {
		return
	}

	address := &models.Address{User: &models.User{Id: userId}, Id: addressId}
	code = (&AddressController{}).getAddressById(address)
	if code != 0 {
		return
	}

	order := &models.Order{User: address.User, CommodityBase: commodity, Count: count, UnitPrice: commodity.Price, RecvAddress: address, CreateTime: time.Now()}
	o := orm.NewOrm()
	if _, err := o.Insert(order); err != nil {
		address.IsUsed = true
		(&AddressController{}).updateAddress(address, []string{"IsUsed"})
		commodity.IsUsed = true
		(&CommodityController{}).updateCommodity(commodity, []string{"IsUsed"})
		code = 900
		return
	}
	code = codes.Code_Right
	return
}
