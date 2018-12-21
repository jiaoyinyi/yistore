package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// 用户
type User struct {
	Id          int32  `orm:"auto;pk"`         //用户ID
	UserName    string `orm:"size(14);unique"` //用户名
	Password    string `orm:"size(16)"`        //密码
	Telphone    string `orm:"size(11);unique"` //电话号码
	Sex         string `orm:"size(2);null"`    //性别
	Description string `orm:"size(50);null"`   //简介
	Permission  int32                          //权限 0为普通用户 1为管理员
}

type CommodityBase struct {
	Id           int32  `orm:"auto;pk"`
	Name         string `orm:"size(200)"`
	Price        float32 //价格
	Status       bool    //0表示下架 1表示在售
	IsUsed       bool    //订单引用到商品就设置IsUsed
	IsDeleted    bool    //假删
	CommodityMsg *CommodityMsg `orm:"rel(one);null"`
	Images       []*Image      `orm:"reverse(many);null"`
}

type CommodityMsg struct {
	Id            int32          `orm:"auto;pk"`
	Brand         string         `orm:"size(45);null"`  //品牌
	Color         string         `orm:"size(45);null"`  //颜色
	Cpu           string         `orm:"size(45);null"`  //cpu
	Date          string         `orm:"size(45);null"`  //出产日期
	Gpu           string         `orm:"size(45);null"`  //gpu
	Bid           string         `orm:"size(45);null"`  //品牌id
	Os            string         `orm:"size(45);null"`  //操作系统
	Position      string         `orm:"size(45);null"`  //定位
	Ram           string         `orm:"size(45);null"`  //内存
	Rom           string         `orm:"size(45);null"`  //硬盘容量
	RomType       string         `orm:"size(45);null"`  //硬盘类型
	ScreenSize    string         `orm:"size(45);null"`  //屏幕大小
	Series        string         `orm:"size(45);null"`  //系列
	Thickness     string         `orm:"size(45);null"`  //厚度
	Type          string         `orm:"size(200);null"` //类型
	Weight        string         `orm:"size(45);null"`  //重量
	CommodityBase *CommodityBase `orm:"reverse(one)"`
}

//type Commodity struct {
//	Id         int32 `orm:"auto;pk"`
//	Status     bool                           //0表示下架 1表示在售
//	Brand      string  `orm:"size(45);null"`  //品牌
//	Color      string  `orm:"size(45);null"`  //颜色
//	Cpu        string  `orm:"size(45);null"`  //cpu
//	Date       string  `orm:"size(45);null"`  //出产日期
//	Gpu        string  `orm:"size(45);null"`  //gpu
//	Bid        string  `orm:"size(45);null"`  //品牌id
//	Name       string  `orm:"size(200);null"` //商品名
//	Os         string  `orm:"size(45);null"`  //操作系统
//	Position   string  `orm:"size(45);null"`  //定位
//	Price      float32 `orm:"null"`           //价格
//	Ram        string  `orm:"size(45);null"`  //内存
//	Rom        string  `orm:"size(45);null"`  //硬盘容量
//	RomType    string  `orm:"size(45);null"`  //硬盘类型
//	ScreenSize string  `orm:"size(45);null"`  //屏幕大小
//	Series     string  `orm:"size(45);null"`  //系列
//	Thickness  string  `orm:"size(45);null"`  //厚度
//	Type       string  `orm:"size(200);null"` //类型
//	Weight     string  `orm:"size(45);null"`  //重量
//	IsUsed     bool                           //订单引用到商品就设置IsUsed
//	IsDeleted  bool                           //假删
//}

// 收藏
type Collection struct {
	Id            int32          `orm:"auto;pk"` //收藏ID
	User          *User          `orm:"rel(fk)"` //引用用户一对多关系
	CommodityBase *CommodityBase `orm:"rel(fk)"` //引用商品一对多关系
}

// 地址
type Address struct {
	Id           int32  `orm:"auto;pk"`   //地址ID
	RecvName     string `orm:"size(10)"`  //收货人或发货人名字
	RecvTelphone string `orm:"size(11)"`  //电话
	Address      string `orm:"size(100)"` //地址
	User         *User  `orm:"rel(fk)"`   //引用用户一对多关系
	IsUsed       bool                     //订单引用到地址就设置IsUseed
	IsDeleted    bool                     //假删
}

//购物车
type ShopCart struct {
	Id            int32          `orm:"auto;pk"`
	User          *User          `orm:"rel(fk)"`
	CommodityBase *CommodityBase `orm:"rel(fk)"`
	Count         int32
}

//订单
type Order struct {
	Id            int32 `orm:"auto;pk"`          //订单ID
	Status        int32                          //订单状态 有三个状态 0用户下单 1管理员确认用户订单 2用户确认订单，订单完成
	CreateTime    time.Time                      //订单创建时间
	FinishTime    time.Time      `orm:"null"`    //订单结束时间
	User          *User          `orm:"rel(fk)"` //引用用户一对多关系
	RecvAddress   *Address       `orm:"rel(fk)"` //引用地址一对多关系
	CommodityBase *CommodityBase `orm:"rel(fk)"` //
	Count         int32
	UnitPrice     float32 //下单时的商品价格
	IsDeleted     bool    //假删
}

//评论
type Comment struct {
	Id            int32  `orm:"auto;pk"`
	Content       string `orm:"size(300)"`
	CreateTime    time.Time
	User          *User          `orm:"rel(fk)"`
	CommodityBase *CommodityBase `orm:"rel(fk)"`
}

//图片
type Image struct {
	Id            int32          `orm:"auto;pk"`   //图片ID
	Path          string         `orm:"size(300)"` //图片路径
	CommodityBase *CommodityBase `orm:"rel(fk)"`   //引用商品一对多关系
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/yistore?charset=utf8")
	orm.RegisterModel(new(User), new(Collection), new(Address),
		new(ShopCart), new(Order), new(CommodityBase), new(CommodityMsg), new(Image), new(Comment))
	orm.RunSyncdb("default", false, true)
}
