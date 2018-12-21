package codes

const (
	Code_Right int32 = 0

	Code_DataFormatError int32 = 100

	Code_NoLoginError          int32 = 200
	Code_PermissionError       int32 = 201
	Code_UsernameEmptyError    int32 = 202
	Code_UsernameFormatError   int32 = 203
	Code_UsernameUsedError     int32 = 204
	Code_UsernameNoExistError  int32 = 205
	Code_PasswordEmptyError    int32 = 206
	Code_PasswordFormatError   int32 = 207
	Code_TelphoneEmptyError    int32 = 208
	Code_TelphoneFormatError   int32 = 209
	Code_TelphoneUsedError     int32 = 210
	Code_TelphoneNoExistError  int32 = 211
	Code_UserRegisterFailError int32 = 212
	Code_UserLoginFailError    int32 = 213
	Code_AdminLoginFailError   int32 = 214
	Code_AdminGetUserError     int32 = 215
	Code_AdminDeleteUserError  int32 = 216
	Code_PasswordNoSameError   int32 = 217
	Code_UserNoExistError      int32 = 218
	Code_SexEmptyError         int32 = 219
	Code_SexFormatError        int32 = 220

	Code_GetCommodityError    int32 = 300
	Code_AddCommodityError    int32 = 301
	Code_UpdateCommodityError int32 = 302
	Code_DeleteCommodityError int32 = 303

	Code_GetAddressError    int32 = 400
	Code_AddAddressError    int32 = 401
	Code_UpdateAddressError int32 = 402
	Code_DeleteAddressError int32 = 403

	Code_GetCollectionError int32 = 500
	Code_AddCollectionError int32 = 501
	//Code_UpdateCollectionError int32 = 502
	Code_DeleteCollectionError int32 = 503

	Code_GetShopcartError    int32 = 600
	Code_AddShopcartError    int32 = 601
	Code_UpdateShopcartError int32 = 602
	Code_DeleteShopcartError int32 = 603

	Code_UserGetOrderError      int32 = 700
	Code_AdminGetOrderError     int32 = 701
	Code_UserAddOrderError      int32 = 702
	Code_AdminCancelOrderError  int32 = 703
	Code_UserComfirmOrderError  int32 = 704
	Code_AdminComfirmOrderError int32 = 705
	Code_UserDeleteOrderError   int32 = 706

	Code_GetCommentError         int32 = 800
	Code_UserAddCommentError     int32 = 801
	Code_UserDeleteCommentError  int32 = 802
	Code_AdminDeleteCommentError int32 = 803
)

func GetMsg(code int32) string {
	switch code {
	case Code_Right:
		return "正常返回"

	case Code_DataFormatError:
		return "数据格式错误"

	case Code_NoLoginError:
		return "未登陆"
	case Code_PermissionError:
		return "权限不对"
	case Code_UsernameEmptyError:
		return "用户名为空"
	case Code_UsernameFormatError:
		return "用户名格式错误"
	case Code_UsernameUsedError:
		return "用户名已被使用"
	case Code_UsernameNoExistError:
		return "用户名不存在"
	case Code_PasswordEmptyError:
		return "密码为空"
	case Code_PasswordFormatError:
		return "密码格式错误"
	case Code_TelphoneEmptyError:
		return "手机号码为空"
	case Code_TelphoneFormatError:
		return "手机号码格式错误"
	case Code_TelphoneUsedError:
		return "手机号码已被使用"
	case Code_TelphoneNoExistError:
		return "手机号码不存在"
	case Code_UserRegisterFailError:
		return "用户注册失败"
	case Code_UserLoginFailError:
		return "用户登录失败"
	case Code_AdminLoginFailError:
		return "管理员登录失败"
	case Code_AdminGetUserError:
		return "管理员获取用户信息失败"
	case Code_AdminDeleteUserError:
		return "管理员删除用户失败"
	case Code_PasswordNoSameError:
		return "两次输入密码不相同"
	case Code_UserNoExistError:
		return "用户不存在"
	case Code_SexEmptyError:
		return "性别为空"
	case Code_SexFormatError:
		return "性别格式错误"

	case Code_GetCommodityError:
		return "获取商品失败"
	case Code_AddCommodityError:
		return "添加商品失败"
	case Code_UpdateCommodityError:
		return "修改商品失败"
	case Code_DeleteCommodityError:
		return "删除商品失败"

	case Code_GetAddressError:
		return "获取地址失败"
	case Code_AddAddressError:
		return "添加地址失败"
	case Code_UpdateAddressError:
		return "修改地址失败"
	case Code_DeleteAddressError:
		return "删除地址失败"

	case Code_GetCollectionError:
		return "获取收藏失败"
	case Code_AddCollectionError:
		return "添加收藏失败"
	case Code_DeleteCollectionError:
		return "删除收藏失败"

	case Code_GetShopcartError:
		return "获取购物车失败"
	case Code_AddShopcartError:
		return "添加购物车失败"
	case Code_UpdateShopcartError:
		return "更新购物车失败"
	case Code_DeleteShopcartError:
		return "删除购物车失败"

	case Code_UserGetOrderError:
		return "用户获取订单失败"
	case Code_AdminGetOrderError:
		return "管理员获取订单失败"
	case Code_UserAddOrderError:
		return "用户添加订单失败"
	case Code_AdminCancelOrderError:
		return "管理员取消订单失败"
	case Code_UserComfirmOrderError:
		return "用户确认订单失败"
	case Code_AdminComfirmOrderError:
		return "管理员确认订单失败"
	case Code_UserDeleteOrderError:
		return "用户删除订单失败"

	case Code_GetCommentError:
		return "获取评论失败"
	case Code_UserAddCommentError:
		return "用户添加评论失败"
	case Code_UserDeleteCommentError:
		return "用户删除评论失败"
	case Code_AdminDeleteCommentError:
		return "管理员删除评论失败"

	default:
		return "服务器错误"
	}
}
