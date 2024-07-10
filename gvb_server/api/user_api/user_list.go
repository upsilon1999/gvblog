package user_api

import (
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	// 如何判断是管理员
	//判断并解析token
	token := c.Request.Header.Get("token")
	if token == "" {
		res.FailWithMessage("未携带token", c)
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMessage("token错误", c)
		return
	}


	//分页获取用户列表数据
	var page models.PageInfo
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 管理员
			user.UserName = ""
		}
		user.Tel = desens.DesensitizationTel(user.Tel)
		user.Email = desens.DesensitizationEmail(user.Email)
		// 脱敏
		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}