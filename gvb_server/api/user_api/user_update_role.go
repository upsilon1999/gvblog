package user_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type UserRole struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"权限参数错误"`
	NickName string     `json:"nickName"` // 防止用户昵称非法，管理员有能力修改
	UserID   uint       `json:"userId" binding:"required" msg:"用户id错误"`
}

// UserUpdateRoleView 用户权限变更
/*
	通过中间件控制这个接口只有管理员可访问

	这个功能只有管理员才有权限，主要的目的有下
	1.修改用户的权限，例如禁言或升级为vip
	2.将用户的非法昵称进行调整
*/
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	//1.获得前端数据 需要userId
	var cr UserRole
	if err := c.ShouldBindJSON(&cr); err != nil {
		res.FailWithError(err, &cr, c)
		return
	}


	var user models.UserModel
	count := global.DB.Take(&user, cr.UserID).RowsAffected
	if count==0 {
		res.FailWithMessage("用户id错误，用户不存在", c)
		return
	}
	err := global.DB.Model(&user).Updates(map[string]any{
		"role":      cr.Role,
		"nick_name": cr.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("修改权限失败", c)
		return
	}
	res.OkWithMessage("修改权限成功", c)
}