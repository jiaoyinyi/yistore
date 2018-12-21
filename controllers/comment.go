package controllers

import (
	"github.com/astaxie/beego/orm"
	"time"
	"yistore/codes"
	"yistore/conf"
	"yistore/models"
	"yistore/objects"
)

type CommentController struct {
	DefaultController
}

//商品id获取，页数获取 不给商品id就返回该用户的评论
func (this *CommentController) UserGetComment() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	page, err := this.GetInt("page")
	if err != nil {
		page = 0
	}

	var comments []*models.Comment
	if commodityId, err := this.GetInt32("commodity_id"); err != nil {
		//判断是否登录
		id, code := this.checkUserLogined()
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		//返回该用户的评论
		code = this.getCommentsByUserId(id, page, &comments)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
	} else {
		//返回该商品的评论
		code := this.getCommentsByCommodityId(commodityId, page, &comments)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
	}
	//处理sendObject
	for _, comment := range comments {
		sendObject := &objects.S2C_CommentObject{UserName: comment.User.UserName, Id: comment.Id, Content: comment.Content, CommodityId: comment.CommodityBase.Id}
		jsonObject.Data = append(jsonObject.Data, sendObject)
	}
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

//商品id获取，按页数获取，按用户id获取
func (this *CommentController) AdminGetComment() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	_, code := this.checkAdminLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	page, err := this.GetInt("page")
	if err != nil {
		page = 0
	}

	var comments []*models.Comment
	if commodityId, err := this.GetInt32("commodity_id"); err == nil {
		code = this.getCommentsByCommodityId(commodityId, page, &comments)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		goto DEAL
	}
	if userId, err := this.GetInt32("user_id"); err == nil {
		code = this.getCommentsByUserId(userId, page, &comments)
		if code != 0 {
			objects.SetJsonObjectMsg(jsonObject, code)
			return
		}
		goto DEAL
	}
	code = this.getCommentsByPage(page, &comments)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}
	goto DEAL

DEAL:
	{
		//处理sendObject
		for _, comment := range comments {
			sendObject := &objects.S2C_CommentObject{UserName: comment.User.UserName, Id: comment.Id, Content: comment.Content, CommodityId: comment.CommodityBase.Id}
			jsonObject.Data = append(jsonObject.Data, sendObject)
		}
		objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
	}
}

func (this *CommentController) AddComment() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_CommentAddObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}
	comment := &models.Comment{User: &models.User{Id: id}, CommodityBase: &models.CommodityBase{Id: recvObject.CommodityId}, Content: recvObject.Content, CreateTime: time.Now()}
	code = this.insertComment(comment)
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *CommentController) UserDeleteComment() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	id, code := this.checkUserLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_CommentUserDeleteObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	comment := &models.Comment{User: &models.User{Id: id}, Id: recvObject.Id}
	code = this.deleteComment(comment)
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *CommentController) AdminDeleteComment() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	_, code := this.checkAdminLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	recvObject := new(objects.C2S_CommentAdminDeleteObject)
	code = this.getJsonData(recvObject)
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	comment := &models.Comment{User: &models.User{Id: recvObject.UserId}, Id: recvObject.Id}
	code = this.deleteComment(comment)
	objects.SetJsonObjectMsg(jsonObject, code)
}

func (this *CommentController) getCommentsByPage(page int, comments *[]*models.Comment) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.Comment)).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(comments); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommentController) getCommentsByUserId(userId int32, page int, comments *[]*models.Comment) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.Comment)).Filter("User__Id", userId).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(comments); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommentController) getCommentsById(comment *models.Comment) (code int32) {
	o := orm.NewOrm()
	err := o.QueryTable(new(models.Comment)).Filter("User__Id", comment.User.Id).Filter("Id", comment.Id).RelatedSel().One(comment)
	if err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommentController) getCommentsByCommodityId(commodityId int32, page int, comments *[]*models.Comment) (code int32) {
	if page < 0 {
		page = 0
	}
	pageCount := conf.EachPageCount
	offset := page * pageCount
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.Comment)).Filter("CommodityBase__Id", commodityId).OrderBy("-Id").Limit(pageCount, offset).RelatedSel().All(comments); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommentController) insertComment(comment *models.Comment) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Insert(comment); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *CommentController) deleteComment(comment *models.Comment) (code int32) {
	o := orm.NewOrm()
	if _, err := o.Delete(comment); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}
