package admin

import (
	"blog/common/utils"
	"blog/dto"
	"blog/middleware"
	"blog/service"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	var q dto.QueryUser
	if err := c.ShouldBind(&q); err != nil {
		middleware.ResponseFail(c, 201, err.Error())
		return
	}
	users, total, err := service.GetUserList(q)
	if err != nil {
		middleware.ResponseFail(c, 202, err.Error())
	} else {
		middleware.ResponseSucc(c, "获取用户列表成功", map[string]interface{}{
			"items": users,
			"total": total,
		})
	}
	return
}
func CreateUser(c *gin.Context) {
	var user dto.UserInfoIn
	err := c.BindJSON(&user)
	if err != nil {
		middleware.ResponseFail(c, 201, err.Error())
		return
	}
	err = service.SaveUser(user)
	if err != nil {
		middleware.ResponseFail(c, 202, err.Error())
	} else {
		middleware.ResponseSucc(c, "添加用户成功", true)
	}
	return
}

func UpdateUser(c *gin.Context) {
	var user dto.UserInfoIn
	err := c.BindJSON(&user)
	if err != nil {
		middleware.ResponseFail(c, 201, err.Error())
		return
	}
	err = service.SaveUser(user)
	if err != nil {
		middleware.ResponseFail(c, 202, err.Error())
	} else {
		middleware.ResponseSucc(c, "更新用户成功", true)
	}
	return
}

func DelUser(c *gin.Context) {
	var user dto.QueryUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		middleware.ResponseFail(c, 201, err.Error())
		return
	}
	err = service.DeleteUser(utils.StrToInt(user.ID))
	if err != nil {
		middleware.ResponseFail(c, 202, err.Error())
	} else {

		middleware.ResponseSucc(c, "删除用户成功", true)

	}
	return
}
