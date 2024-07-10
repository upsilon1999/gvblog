package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils"
	jwts "gvb_server/utils/jwt"

	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"userName" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}


	//验证用户是否存在
	var userModel models.UserModel
	//我们页面上传入的用户名、邮箱等实际上都被cr.UserName接收
	count := global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.UserName).RowsAffected
	if count == 0 {
		// 没找到
		global.Log.Warn("用户名不存在")
		res.FailWithMessage("用户名不存在", c)
		return
	}
	// 校验密码
	isCheck := utils.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		res.FailWithMessage("用户密码错误", c)
		return
	}
	// 登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("token生成失败", c)
		return
	}
	res.OkWithData(token, c)

}