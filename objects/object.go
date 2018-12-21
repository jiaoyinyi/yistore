package objects

import (
	"fmt"
	"time"
	"yistore/codes"
	"yistore/conf"
	"yistore/models"
)

type JsonObject struct {
	Code int32         //状态码
	Msg  string        //消息
	Data []interface{} //数据
}

func New() (object *JsonObject) {
	object = new(JsonObject)
	object.Data = make([]interface{}, 0)
	return
}

//设置JsonObject引用的状态码和信息
func SetJsonObjectMsg(object *JsonObject, code int32) {
	object.Code = code
	object.Msg = codes.GetMsg(code)
}

//////////////////////////////////////////////////

type S2C_UserNameObject struct {
	UserName string `json:"username"`
}

type S2C_UserObject struct {
	UserName    string `json:"username"` //用户名
	Password    string
	Telphone    string //电话号码
	Sex         string
	Description string //简介
}

type S2C_CommodityBaseObject struct {
	Id     int32    `json:"commodity_id"`
	Name   string   `json:"name"`
	Price  float32  `json:"price"`  //价格
	Status bool     `json:"status"` //0表示下架 1表示在售
	Images []string `json:"images"`
}

type S2C_CommodityObject struct {
	Id         int32    `json:"commodity_id"`
	Name       string   `json:"name"`
	Price      float32  `json:"price"`  //价格
	Status     bool     `json:"status"` //0表示下架 1表示在售
	Images     []string `json:"images"`
	Brand      string   `json:"brand"`
	Color      string   `json:"color"`
	Cpu        string   `json:"cpu"`
	Date       string   `json:"date"`
	Gpu        string   `json:"gpu"`
	Bid        string   `json:"bid"`
	Os         string   `json:"os"`
	Position   string   `json:"position"`
	Ram        string   `json:"ram"`
	Rom        string   `json:"rom"`
	RomType    string   `json:"rom_type"`
	ScreenType string   `json:"screen_type"`
	Series     string   `json:"series"`
	Thickness  string   `json:"thickness"`
	Type       string   `json:"type"`
	Weight     string   `json:"weight"`
}

type S2C_CollectionObject struct {
	Id        int32                    `json:"collection_id"`
	Commodity *S2C_CommodityBaseObject `json:"commodity"`
}

type S2C_ShopCartObject struct {
	Id        int32                    `json:"shopcart_id"`
	Count     int32                    `json:"count"`
	Commodity *S2C_CommodityBaseObject `json:"commodity"`
}

type S2C_AddressObject struct {
	Id           int32  `json:"address_id"`
	RecvName     string `json:"recv_name"`
	RecvTelphone string `json:"recv_telphone"`
	Address      string `json:"address"`
}

type S2C_OrderObject struct {
	Id          int32                    `json:"order_id"`
	Status      int32                    `json:"status"`
	CreateTime  time.Time                `json:"create_time"`
	FinishTime  time.Time                `json:"finish_time"`
	Count       int32                    `json:"count"`
	UnitPrice   float32                  `json:"unit_price"`
	Commodity   *S2C_CommodityBaseObject `json:"commodity"`
	RecvAddress *S2C_AddressObject       `json:"recv_address"`
}

type S2C_CommentObject struct {
	Id          int32  `json:"comment_id"`
	UserName    string `json:"username"`
	Content     string `json:"content"`
	CommodityId int32  `json:"commodity_id"`
}

type S2C_ImageObject struct {
	Path string `json:"path"`
}

///////////////////////////////////////////////////

type C2S_UserRegisterObject struct {
	UserName   string `json:"username"`
	Password   string
	Repassword string
	Telphone   string
}

type C2S_UserLoginObject struct {
	UserName string `json:"username"`
	Password string
}

type C2S_UserUpdateObject struct {
	UserName    string `json:"username"`
	Password    string
	Telphone    string
	Description string
	Sex         string
}

type C2S_CheckUserNameObject struct {
	UserName string `json:"username"`
}

type C2S_CheckTelphoneObject struct {
	Telphone string `json:"telphone"`
}

type C2S_CollectionAddObject struct {
	CommodityId int32 `json:"commodity_id"`
}

type C2S_CollectionDeleteObject struct {
	Id int32 `json:"collection_id"`
}

type C2S_ShopCartAddObject struct {
	CommodityId int32 `json:"commodity_id"`
	Count       int32
}

type C2S_ShopCartUpdateObject struct {
	Id    int32 `json:"shopcart_id"`
	Count int32 `json:"count"`
}

type C2S_ShopCartDeleteObject struct {
	Id int32 `json:"shopcart_id"`
}

type C2S_AddressAddOject struct {
	RecvName     string `json:"recv_name"`
	RecvTelphone string `json:"recv_telphone"`
	Address      string
}

type C2S_AddressUpdateOject struct {
	Id           int32  `json:"address_id"`
	RecvName     string `json:"recv_name"`
	RecvTelphone string `json:"recv_telphone"`
	Address      string
}

type C2S_AddressDeleteObject struct {
	Id int32 `json:"address_id"`
}

type C2S_OrderAddObject struct {
	ShopCartId  int32 `json:"shopcart_id"`
	CommodityId int32 `json:"commodity_id"`
	Count       int32 `json:"count"`
	AddressId   int32 `json:"address_id"`
}

type C2S_OrderUserUpdateObject struct {
	Id int32 `json:"order_id"`
}

type C2S_OrderAdminUpdateObject struct {
	Id     int32 `json:"order_id"`
	UserId int32 `json:"user_id"`
}

type C2S_OrderDeleteObject struct {
	Id int32 `json:"order_id"`
}

type C2S_CommentAddObject struct {
	CommodityId int32  `json:"commodity_id"`
	Content     string `json:"content"`
}

type C2S_CommentUserDeleteObject struct {
	Id int32 `json:"comment_id"`
}

type C2S_CommentAdminDeleteObject struct {
	UserId int32 `json:"user_id"`
	Id     int32 `json:"comment_id"`
}

type C2S_CommodityDeleteObject struct {
	Id int32 `json:"commodity_id"`
}

type C2S_CommodityBaseUpdateObject struct {
	Id     int32    `json:"commodity_id"`
	Name   string   `json:"name"`
	Price  float32  `json:"price"`
	Images []string `json:"images"`
}

type C2S_CommodityAddObject struct {
	Id         int32    `json:"commodity_id"`
	Name       string   `json:"name"`
	Price      float32  `json:"price"`  //价格
	Status     bool     `json:"status"` //0表示下架 1表示在售
	Images     []string `json:"images"`
	Brand      string   `json:"brand"`
	Color      string   `json:"color"`
	Cpu        string   `json:"cpu"`
	Date       string   `json:"date"`
	Gpu        string   `json:"gpu"`
	Bid        string   `json:"bid"`
	Os         string   `json:"os"`
	Position   string   `json:"position"`
	Ram        string   `json:"ram"`
	Rom        string   `json:"rom"`
	RomType    string   `json:"rom_type"`
	ScreenType string   `json:"screen_type"`
	Series     string   `json:"series"`
	Thickness  string   `json:"thickness"`
	Type       string   `json:"type"`
	Weight     string   `json:"weight"`
}

type C2S_CommodityUpdateObject struct {
	Id         int32    `json:"commodity_id"`
	Name       string   `json:"name"`
	Price      float32  `json:"price"`  //价格
	Status     bool     `json:"status"` //0表示下架 1表示在售
	Images     []string `json:"images"`
	Brand      string   `json:"brand"`
	Color      string   `json:"color"`
	Cpu        string   `json:"cpu"`
	Date       string   `json:"date"`
	Gpu        string   `json:"gpu"`
	Bid        string   `json:"bid"`
	Os         string   `json:"os"`
	Position   string   `json:"position"`
	Ram        string   `json:"ram"`
	Rom        string   `json:"rom"`
	RomType    string   `json:"rom_type"`
	ScreenType string   `json:"screen_type"`
	Series     string   `json:"series"`
	Thickness  string   `json:"thickness"`
	Type       string   `json:"type"`
	Weight     string   `json:"weight"`
}

type C2S_CommodityStatusUpdateObject struct {
	Id     int32 `json:"commodity_id"`
	Status bool  `json:"status"`
}

/////////////////////////////////////////////////

func SetCommodityBaseObject(object *S2C_CommodityBaseObject, commodity *models.CommodityBase) {
	object.Id = commodity.Id
	object.Name = commodity.Name
	object.Price = commodity.Price
	object.Status = commodity.Status

	for _, image := range commodity.Images {
		object.Images = append(object.Images, fmt.Sprintf("http://%s:%s/static/img/%s", conf.RemoteIp, conf.RemotePort, image.Path))
	}
}

func SetCommodityObject(object *S2C_CommodityObject, commodity *models.CommodityBase) {
	object.Id = commodity.Id
	object.Name = commodity.Name
	object.Price = commodity.Price
	object.Status = commodity.Status

	object.Brand = commodity.CommodityMsg.Brand
	object.Color = commodity.CommodityMsg.Color
	object.Cpu = commodity.CommodityMsg.Cpu
	object.Date = commodity.CommodityMsg.Date
	object.Gpu = commodity.CommodityMsg.Gpu
	object.Bid = commodity.CommodityMsg.Bid
	object.Os = commodity.CommodityMsg.Os
	object.Position = commodity.CommodityMsg.Position
	object.Ram = commodity.CommodityMsg.Ram
	object.Rom = commodity.CommodityMsg.Rom
	object.RomType = commodity.CommodityMsg.RomType
	object.ScreenType = commodity.CommodityMsg.ScreenSize
	object.Series = commodity.CommodityMsg.Series
	object.Thickness = commodity.CommodityMsg.Thickness
	object.Type = commodity.CommodityMsg.Type
	object.Weight = commodity.CommodityMsg.Weight

	for _, image := range commodity.Images {
		object.Images = append(object.Images, fmt.Sprintf("http://%s:%s/static/img/%s", conf.RemoteIp, conf.RemotePort, image.Path))
	}
}

func SetAddressObject(object *S2C_AddressObject, address *models.Address) {
	object.Id = address.Id
	object.Address = address.Address
	object.RecvName = address.RecvName
	object.RecvTelphone = address.RecvTelphone
}

func SetImage(object *[]string, images []*models.Image) {
	for _, image := range images {
		*object = append(*object, image.Path)
	}
}
