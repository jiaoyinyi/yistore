package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"os"
	"strings"
	"yistore/codes"
	"yistore/models"
	"yistore/objects"
)

var Image_Suffixs = map[string]string{"jpg": "jpg", "jpeg": "jpeg", "png": "png"}

type ImageController struct {
	DefaultController
}

func (this *ImageController) AddImage() {
	jsonObject := objects.New()
	defer this.sendJsonObject(jsonObject)

	//判断是否登录
	_, code := this.checkAdminLogined()
	if code != 0 {
		objects.SetJsonObjectMsg(jsonObject, code)
		return
	}

	f, h, err := this.GetFile("upload_file")
	if err != nil {
		logs.Debug("上传文件出错")
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}
	defer f.Close()

	strs := strings.Split(h.Filename, ".")
	suffix := strs[len(strs)-1]
	if _, ok := Image_Suffixs[suffix]; !ok {
		objects.SetJsonObjectMsg(jsonObject, codes.Code_DataFormatError)
		return
	}

	this.SaveToFile("upload_file", "static/img/"+h.Filename)
	objects.SetJsonObjectMsg(jsonObject, codes.Code_Right)
}

func (this *ImageController) getImagesByCommodityId(commodityId int32, images *[]*models.Image) (code int32) {
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(models.Image)).Filter("CommodityBase__Id", commodityId).All(images); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *ImageController) getImageById(image *models.Image) (code int32) {
	o := orm.NewOrm()
	if err := o.QueryTable(new(models.Image)).Filter("Id", image.Id).One(image); err != nil {
		code = 900
		return
	}
	code = codes.Code_Right
	return
}

func (this *ImageController) saveFile()  {

}

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
