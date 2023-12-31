// Code generated by hertz generator.

package sys

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/liukunc9/hertz_template/biz/model/sys"
	"github.com/liukunc9/hertz_template/dao/model"
	"github.com/liukunc9/hertz_template/global"
	"github.com/liukunc9/hertz_template/response"
	"github.com/mitchellh/mapstructure"
	"strings"
)

// SaveUser .
// @router /api/v1/sys/user [POST]
func SaveUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.SaveUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	resp := new(sys.SaveUserResp)

	var sysUser model.SysUser
	err = mapstructure.Decode(req.UserInfo, &sysUser)
	if err != nil {
		response.FailWithMessage(c, err.Error())
	}
	sysUser.UserID = global.Snowflake.Generate().String()

	err = global.DB.Create(&sysUser).Error
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	resp.Data = sysUser.UserID
	response.OkWithData(c, resp.Data)
}

// GetUser .
// @router /api/v1/sys/user [GET]
func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.GetUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	userIds := strings.Split(req.UserIds, ",")

	resp := new(sys.GetUserResp)
	var users []model.SysUser
	global.DB.Where("user_id in ?", userIds).Find(&users)
	for _, sysUser := range users {
		var userInfo sys.UserInfo
		err = mapstructure.Decode(sysUser, &userInfo)
		if err != nil {
			response.FailWithMessage(c, err.Error())
			return
		}
		resp.Data = append(resp.Data, &userInfo)
	}
	response.OkWithData(c, resp.Data)
}

// UpdateUser .
// @router /api/v1/sys/user/:id [PUT]
func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req sys.UpdateUserReq
	err = c.BindAndValidate(&req)
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}

	resp := new(sys.UpdateUserResp)
	var sysUser model.SysUser
	if err = mapstructure.Decode(req.UserInfo, &sysUser); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	// sysUser.UserID = "" 如果将值设置为该类型的默认值，gorm也不会更新该字段
	global.DB.Table(sysUser.TableName()).
		Where("user_id = ?", req.UserID).
		Omit("id", "user_id").
		Updates(&sysUser)

	response.OkWithData(c, resp.Data)
}
