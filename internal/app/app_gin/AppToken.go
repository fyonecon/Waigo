package app_gin

import (
	"datathink.top/Waigo/internal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppToken struct{}

// CheckAppToken 校验
func (at *AppToken) CheckAppToken(ctx *gin.Context) {
	appToken := common.RequestInput(ctx, "app_token")
	appTokenState := common.CheckRandToken("20251229", appToken)
	if appTokenState {
		ctx.Next()
	} else {
		//
		ctx.JSON(http.StatusOK, gin.H{
			"state":   0,
			"msg":     "Auth Error",
			"content": map[string]interface{}{},
		})
		ctx.Abort()
	}
	return
}

// MakeAppToken 生成
func (at *AppToken) MakeAppToken() string {
	return common.MakeRandToken("20251229", 2*365*24*60*60)
}
