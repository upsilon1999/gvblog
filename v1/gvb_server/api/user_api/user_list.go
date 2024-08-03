package user_api

import (
	"fmt"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/common"
	"gvb_server/utils"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
)

func (UserApi) UserListView(c *gin.Context) {
	//获取从中间件来的数据
	_claims,_ := c.Get("claims")
    fmt.Println("token是",_claims)
	//由于Get获取的值是any类型，所以要进行类型断言
	claims := _claims.(*jwts.CustomClaims)
	//分页获取用户列表数据
	var page models.PageInfo
	//前端必须得传limit，即条数，因为现在gorm默认limit为0
	if err := c.ShouldBindQuery(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
		Debug: true,
	})
	fmt.Println(list)
	for _, user := range list {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 如果不是管理员，就不能看到userName
			user.UserName = ""
		}
		user.Tel = utils.DesensitizationTel(user.Tel)
		user.Email = utils.DesensitizationEmail(user.Email)
		// 脱敏
		users = append(users, user)
	}

	res.OkWithList(users, count, c)
}

