package app_window

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/app/app_gin"
	"datathink.top/Waigo/internal/app/app_window/window_controller"
	"datathink.top/Waigo/internal/common"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

// RouteWindow window专用路由
func (awd *AppWindow) RouteWindow(route *gin.Engine, ginHTML fs.FS, ginFiles fs.FS) {
	// info
	//appName := common.InterfaceToString(internal.GetConfigMap("app", "appName"))
	//appRights := common.InterfaceToString(internal.GetConfigMap("app", "appRights"))
	//appVersion := common.InterfaceToString(internal.GetConfigMap("app", "appVersion"))

	// 中间件
	mdw := app_gin.Middlewares{}
	//gcl := gin_controller.GinController{}

	// 重写js_call_go
	// 如：http://127.0.0.1:9850/api/js_call_go
	route.Any("/api/js_call_go", mdw.HttpCorsApi, mdw.HttpError500, func(ctx *gin.Context) {
		//
		_appClass := common.RequestInput(ctx, "app_class")
		_appVersion := common.RequestInput(ctx, "app_version")
		//
		_windowToken := common.RequestInput(ctx, "window_token")
		windowTokenState := common.CheckRandID(_windowToken)
		//
		_apiURL := common.RequestFullURL(ctx) // 完整访问域名
		whiteHosts := internal.GetConfigMap("gin", "whiteHosts")
		apiURLState := common.ArrayInString(common.InterfaceToArrayString(whiteHosts), _apiURL) != -1
		//
		key := common.RequestInput(ctx, "key")
		_dataDict := common.RequestInput(ctx, "data_dict")
		dataDict := common.JsonStringToMap(_dataDict)
		//
		//fmt.Println("js_call_go=", _windowToken, windowTokenState, _apiURL, apiURLState, key, _dataDict)
		//
		if windowTokenState && apiURLState { // 校验通过就返回对照表
			var awd = window_controller.WindowController{}
			ctx.JSON(http.StatusOK, awd.ListJSCallGo(key, dataDict))
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"state": 0,
				"msg":   "参数错误或过期",
				"content": map[string]interface{}{
					"windowToken": _windowToken,
					"whiteHosts":  whiteHosts,
					"appClass":    _appClass,
					"appVersion":  "v" + _appVersion,
				},
			})
		}
		return
	})

	// view_js必要参数
	// 如：http://127.0.0.1:9850/js_must_data.js
	route.Any("/js_must_data.js", mdw.HttpCorsHTML, mdw.HttpError500, awd.JSMustData)
	//

}
