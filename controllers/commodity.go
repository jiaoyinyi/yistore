package controllers

import (
	"github.com/astaxie/beego/orm"
	"yistore/codes"
	"yistore/conf"
	"yistore/models"
	"yistore/objects"
)

type CommodityController struct {
	DefaultController
}

func (this *CommodityController) AddCommodity() {

}

func (this *CommodityController) DeleteCommodity() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	_, code := this.checkAdminLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_CommodityDeleteObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	commodity := &models.CommodityBase{Id: recvObject.Id}
	code = this.deleteCommodity(commodity)
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *CommodityController) UpdateCommodity() {

}

func (this *CommodityController) UpdateCommodityStatus() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	_, code := this.checkAdminLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_CommodityStatusUpdateObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	commodity := &models.CommodityBase{Id: recvObject.Id, Status: recvObject.Status}
	code = this.updateCommodity(commodity, []string{"Status"})
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *CommodityController) GetCommodityBase() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	commodityId, err := this.GetInt32("commodity_id")
	if err == nil {
		commodity := &models.CommodityBase{Id: commodityId}
		code := this.getCommodityBaseById(commodity)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		var images []*models.Image
		(&ImageController{}).getImagesByCommodityId(commodity.Id, &images)
		commodity.Images = images
		sendObject := new(objects.S2C_CommodityBaseObject)
		objects.SetCommodityBaseObject(sendObject, commodity)
		jsonObject.Data = append(jsonObject.Data, sendObject)
		objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
		return
	}

	page, err := this.GetInt("page")
	if err != nil {
		page = 0
	}

	start, err := this.GetFloat("start")
	if err != nil {
		start = 0
	} else {
		if start < 0 {
			start = 0
		}
	}

	end, err := this.GetFloat("end")
	if err != nil {
		end = 0
	} else {
		if end < 0 {
			end = 0
		}
	}

	if start > end {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}

	brand := this.GetString("brand")
	var commoditys []*models.CommodityBase
	if brand != "" {
		if start != 0 && end != 0 {
			code := this.getCommodityBasesByBrandAndPageOrderByPrice(brand, page, float32(start), float32(end), &commoditys)
			if code != 0 {
				objects.SetJsonObjectMsg(jsonObject, code)
				return
			}
		} else {
			code := this.getCommodityBasesByBrandAndPage(brand, page, &commoditys)
			if code != 0 {
				objects.SetJsonObjectMsg(jsonObject, code)
				return
			}
		}
	} else {
		if start != 0 && end != 0 {
			code := this.getCommodityBasesByPageOrderByPrice(page, float32(start), float32(end), &commoditys)
			if code != 0 {
				objects.SetJsonObjectMsg(jsonObject, code)
				return
			}
		} else {
			code := this.getCommodityBasesByPage(page, &commoditys)
			if code != 0 {
				objects.SetJsonObjectMsg(jsonObject, code)
				return
			}
		}
	}

	for _, commodity := range commoditys {
		var images []*models.Image
		(&ImageController{}).getImagesByCommodityId(commodity.Id, &images)
		commodity.Images = images
		sendObject := new(objects.S2C_CommodityBaseObject)
		objects.SetCommodityBaseObject(sendObject, commodity)
		jsonObject.Data = append(jsonObject.Data, sendObject)
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
	return
}

//GET 带id就返回该id的商品信息，否则返回全部商品信息
func (this *CommodityController) GetCommodity() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	commodityId, err := this.GetInt32("commodity_id")
	if err == nil {
		commodity := &models.CommodityBase{Id: commodityId}
		code := this.getCommodityById(commodity)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		var images []*models.Image
		(&ImageController{}).getImagesByCommodityId(commodity.Id, &images)
		commodity.Images = images
		sendObject := new(objects.S2C_CommodityObject)
		objects.SetCommodityObject(sendObject, commodity)
		jsonObject.Data = append(jsonObject.Data, sendObject)
		objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
		return
	}

	page, err := this.GetInt("page")
	if err != nil {
		page = 0
	}

	start, err := this.GetFloat("start")
	if err != nil {
		start = 0
	} else {
		if start < 0 {
			start = 0
		}
	}

	end, err := this.GetFloat("end")
	if err != nil {
		end = 0
	} else {
		if end < 0 {
			end = 0
		}
	}

	if start > end {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}

	brand := this.GetString("brand")
	var commoditys []*models.CommodityBase
	if brand != "" {
		if start != 0 && end != 0 {
			code := this.getCommoditysByBrandAndPageOrderByPrice(brand, page, float32(start), float32(end), &commoditys)
			if code != 0 {
				objects.SetJsonObjectMsg(jsonObject, code)
				return
			}
		} else {
			code := this.getCommoditysByBrandAndPage(brand, page, &commoditys)
			if code != 0 {
				objects.SetJsonObjectMsg(jsonObject, code)
				return
			}
		}
	} else {
		if start != 0 && end != 0 {
			code := this.getCommoditysByPageOrderByPrice(page, float32(start), float32(end), &commoditys)
			if code != 0 {
				objects.SetJsonObjectMsg(jsonObject, code)
				return
			}
		} else {
			code := this.getCommoditysByPage(page, &commoditys)
			if code != 0 {
				objects.SetJsonObjectMsg(jsonObject, code)
				return
			}
		}
	}

	for _, commodity := range commoditys {
		var images []*models.Image
		(&ImageController{}).getImagesByCommodityId(commodity.Id, &images)
		commodity.Images = images
		sendObject := new(objects.S2C_CommodityObject)
		objects.SetCommodityObject(sendObject, commodity)
		jsonObject.Data = append(jsonObject.Data, sendObject)
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
	return
}

func (this *CommodityController) getCommodityBaseById(commodity *models.CommodityBase) (code int32) {
	o := orm.NewOrm()
	if err := o.QueryTable(new(models.CommodityBase)).Filter("Id", commodity.Id).One(commodity); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommodityById(commodity *models.CommodityBase) (code int32) {
	o := orm.NewOrm()
	if err := o.QueryTable(new(models.CommodityBase)).Filter("Id", commodity.Id).RelatedSel().One(commodity); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommodityBasesByPage(page int, commoditys *[]*models.CommodityBase) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.CommodityBase)).OrderBy("-Id").Limit(pageCount, offset).All(commoditys); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommoditysByPage(page int, commoditys *[]*models.CommodityBase) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.CommodityBase)).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(commoditys); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommodityBasesByBrandAndPage(brand string, page int, commoditys *[]*models.CommodityBase) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.CommodityBase)).Filter("CommodityMsg__Brand__icontains", brand).OrderBy("-Id").Limit(pageCount, offset).All(commoditys); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommoditysByBrandAndPage(brand string, page int, commoditys *[]*models.CommodityBase) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.CommodityBase)).Filter("CommodityMsg__Brand__icontains", brand).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(commoditys); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommodityBasesByPageOrderByPrice(page int, start float32, end float32, commoditys *[]*models.CommodityBase) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount

	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.CommodityBase)).Filter("Price__gte", start).Filter("Price__lte", end).OrderBy("-Id").Limit(pageCount, offset).All(commoditys); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommoditysByPageOrderByPrice(page int, start float32, end float32, commoditys *[]*models.CommodityBase) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount

	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.CommodityBase)).Filter("Price__gte", start).Filter("Price__lte", end).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(commoditys); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommodityBasesByBrandAndPageOrderByPrice(brand string, page int, start float32, end float32, commoditys *[]*models.CommodityBase) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount

	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.CommodityBase)).Filter("Brand__contains", brand).Filter("Price__gte", start).Filter("Price__lte", end).OrderBy("-Id").Limit(pageCount, offset).All(commoditys); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) getCommoditysByBrandAndPageOrderByPrice(brand string, page int, start float32, end float32, commoditys *[]*models.CommodityBase) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount

	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.CommodityBase)).Filter("Brand__contains", brand).Filter("Price__gte", start).Filter("Price__lte", end).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(commoditys); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *AddressController) insertCommodity(commodity *models.CommodityBase) (code int32) {
	o := orm.NewOrm()
	if _, _, err := o.ReadOrCreate(commodity, "Id"); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommodityController) deleteCommodity(commodity *models.CommodityBase) (code int32) {
	code = this.getCommodityBaseById(commodity)
	if code != 0 {
		return
	}
	commodity.Status = false
	commodity.IsDeleted = true

	code = this.updateCommodity(commodity, []string{"Status", "IsDeleted"})
	return
}

func (this *CommodityController) updateCommodity(commodity *models.CommodityBase, cols []string) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Update(commodity, cols...); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}
