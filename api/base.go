package api

import (
	"server/global"
	"server/model/request"
	response2 "server/model/response"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type BaseApi struct {
}

var store = base64Captcha.DefaultMemStore

func (baseApi *BaseApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(
		global.Config.Captcha.Height,
		global.Config.Captcha.Width,
		global.Config.Captcha.Length,
		global.Config.Captcha.MaxSkew,
		global.Config.Captcha.DotCount,
	)

	captcha := base64Captcha.NewCaptcha(driver, store)

	id, b64s, _, err := captcha.Generate()
	if err != nil {
		global.Log.Error("验证码生成失败", zap.Error(err))
		response2.FailWithMessage("验证码生成失败", c)
		return
	}
	response2.OkWithData(response2.Captcha{
		CaptchaID: id,
		PicPath:   b64s,
	}, c)
}

func (baseApi *BaseApi) SendEmailVerificationCode(c *gin.Context) {
	var req request.SendEmailVerificationCode
	if err := c.ShouldBindJSON(&req); err != nil {
		response2.FailWithMessage(err.Error(), c)
		return
	}

	if store.Verify(req.CaptchaID, req.Captcha, true) {
		err := baseService.SendEmailVerificationCode(c, req.Email)
		if err != nil {
			global.Log.Error("发送验证码失败", zap.Error(err))
			response2.FailWithMessage(err.Error(), c)
			return
		}
		response2.OkWithMessage("发送验证码成功", c)
		return
	}
	response2.FailWithMessage("验证码错误", c)
}

// QQLoginURL 返回 QQ 登录链接
func (baseApi *BaseApi) QQLoginURL(c *gin.Context) {
	url := global.Config.QQ.QQLoginURL()
	response2.OkWithData(url, c)
}
