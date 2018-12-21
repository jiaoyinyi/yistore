package routers

import (
	"github.com/astaxie/beego"
	"yistore/controllers"
)

func init() {
	//普通用户
	beego.Router("/user/register", &controllers.UserController{}, "post:RegisterPost")
	beego.Router("/user/login", &controllers.UserController{}, "get:LoginGet;post:LoginPost")
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
	beego.Router("/user/get", &controllers.UserController{}, "get:UserGetMsg")
	beego.Router("/user/update", &controllers.UserController{}, "post:UserUpdateMsg")

	beego.Router("/user/check_username", &controllers.UserController{}, "post:CheckUserName")
	beego.Router("/user/check_telphone", &controllers.UserController{}, "post:CheckTelphone")
	//管理员
	beego.Router("/admin/login", &controllers.AdminController{}, "post:LoginPost")
	beego.Router("/admin/logout", &controllers.AdminController{}, "get:Logout")

	////商品
	beego.Router("/commodity/get", &controllers.CommodityController{}, "get:GetCommodityBase")
	beego.Router("/commodity/get_detail", &controllers.CommodityController{}, "get:GetCommodity")

	//地址
	beego.Router("/address/get", &controllers.AddressController{}, "get:GetAddress")
	beego.Router("/address/add", &controllers.AddressController{}, "post:AddAddress")
	beego.Router("/address/update", &controllers.AddressController{}, "post:UpdateAddress")
	beego.Router("/address/delete", &controllers.AddressController{}, "post:DeleteAddress")

	//收藏
	beego.Router("/collection/get", &controllers.CollectionController{}, "get:GetCollection")
	beego.Router("/collection/add", &controllers.CollectionController{}, "post:AddCollection")
	beego.Router("/collection/delete", &controllers.CollectionController{}, "post:DeleteCollection")

	//购物车
	beego.Router("/shopcart/get", &controllers.ShopCartController{}, "get:ShopCartGet")
	beego.Router("/shopcart/add", &controllers.ShopCartController{}, "post:ShopCartAdd")
	beego.Router("/shopcart/update", &controllers.ShopCartController{}, "post:ShopCartUpdate")
	beego.Router("/shopcart/delete", &controllers.ShopCartController{}, "post:ShopCartDelete")

	//订单
	beego.Router("/order/user/get", &controllers.OrderController{}, "get:UserGetOrder")
	beego.Router("/order/user/add", &controllers.OrderController{}, "post:AddOrder")
	beego.Router("/order/user/comfirm", &controllers.OrderController{}, "post:UserComfirmOrder")
	beego.Router("/order/admin/comfirm", &controllers.OrderController{}, "post:AdminComfirmOrder")
	beego.Router("/order/admin/cancel", &controllers.OrderController{}, "post:AdminCancelOrder")

	//评论
	beego.Router("/comment/user/get", &controllers.CommentController{}, "get:UserGetComment")
	beego.Router("/comment/admin/get", &controllers.CommentController{}, "get:AdminGetComment")
	beego.Router("/comment/add", &controllers.CommentController{}, "post:AddComment")
	beego.Router("/comment/user/delete", &controllers.CommentController{}, "post:UserDeleteComment")
	beego.Router("/comment/admin/delete", &controllers.CommentController{}, "post:AdminDeleteComment")

	//图片
	beego.Router("/image/add", &controllers.ImageController{}, "post:AddImage")
}
