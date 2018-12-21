package controllers

import (
	"github.com/astaxie/beego/orm"
	"yistore/codes"
	"yistore/conf"
	"yistore/models"
	"yistore/objects"
)

type CollectionController struct {
	DefaultController
}

//GET
func (this *CollectionController) GetCollection() {
	//获取收藏信息数据
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	if collectionId, err := this.GetInt32("collection_id"); err != nil {
		var collections []*models.Collection

		page, err := this.GetInt("page")
		if err != nil {
			page = 0
		}

		code = this.getCollectionByPage(id, page, &collections)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, codes.Code_GetCollectionError)
			return
		}

		for _, collection := range collections {
			var images []*models.Image
			var paths []string
			(&ImageController{}).getImagesByCommodityId(collection.CommodityBase.Id, &images)
			objects.SetImage(&paths, images)
			commodity := &objects.S2C_CommodityBaseObject{Id: collection.CommodityBase.Id, Name: collection.CommodityBase.Name, Price: collection.CommodityBase.Price, Status: collection.CommodityBase.Status, Images: paths}
			sendObject := &objects.S2C_CollectionObject{Id: collection.Id, Commodity: commodity}
			jsonObject.Data = append(jsonObject.Data, sendObject)
		}
	} else {
		collection := &models.Collection{User: &models.User{Id: id}, Id: collectionId}
		code = this.getCollectionsById(collection)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, codes.Code_GetCollectionError)
			return
		}
		var images []*models.Image
		var paths []string
		(&ImageController{}).getImagesByCommodityId(collection.CommodityBase.Id, &images)
		objects.SetImage(&paths, images)
		commodity := &objects.S2C_CommodityBaseObject{Id: collection.CommodityBase.Id, Name: collection.CommodityBase.Name, Price: collection.CommodityBase.Price, Status: collection.CommodityBase.Status, Images: paths}
		sendObject := &objects.S2C_CollectionObject{Id: collection.Id, Commodity: commodity}
		jsonObject.Data = append(jsonObject.Data, sendObject)
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//POST
func (this *CollectionController) AddCollection() {
	//处理添加收藏信息
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_CollectionAddObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	collection := &models.Collection{User: &models.User{Id: id}, CommodityBase: &models.CommodityBase{Id: recvObject.CommodityId}}
	code = this.insertCollection(collection)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_AddCollectionError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)

}

//POST
func (this *CollectionController) DeleteCollection() {
	//处理删除收藏信息
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	_, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_CollectionDeleteObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	collection := &models.Collection{Id: recvObject.Id}
	code = this.deleteCollection(collection)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DeleteCollectionError)
		return
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//func (this *CollectionController) getAllCollections(userId int32, collections *[]*models.Collection) (code int32) {
//	o := orm.NewOrm()
//	if _, err := o.QueryTable(new(models.Collection)).Filter("User__Id", userId).OrderBy("-Id").RelatedSel().All(collections); err != nil {
//		code = 900
//		return
//	}
//	code = codes.Code_Right
//	return
//}

func (this *CollectionController) getCollectionByPage(userId int32, page int, collections *[]*models.Collection) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.Collection)).Filter("User__Id", userId).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(collections); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CollectionController) getCollectionsById(collection *models.Collection) (code int32) {
	o := orm.NewOrm()
	if err := o.QueryTable(new(models.Collection)).Filter("User__Id", collection.User.Id).Filter("Id", collection.Id).RelatedSel().One(collection); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CollectionController) insertCollection(collection *models.Collection) (code int32) {
	o := orm.NewOrm()
	if _, _, err := o.ReadOrCreate(collection, "User", "CommodityBase"); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CollectionController) deleteCollection(collection *models.Collection) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Delete(collection); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CollectionController) updateCollection(collection *models.Collection, cols []string) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Update(collection, cols...); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}
