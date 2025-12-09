package app_gin

import (
	"github.com/gin-gonic/gin"
	"io/fs"
)

// RouteAPI api专用路由
func (agn *AppGin) RouteAPI(route *gin.Engine, ginHTML fs.FS, ginFiles fs.FS) {
	// info
	//appName := common.InterfaceToString(internal.GetConfigMap("app", "appName"))
	//appRights := common.InterfaceToString(internal.GetConfigMap("app", "appRights"))
	//appVersion := common.InterfaceToString(internal.GetConfigMap("app", "appVersion"))
	// 中间件
	//mdw := Middlewares{}
	//

}
