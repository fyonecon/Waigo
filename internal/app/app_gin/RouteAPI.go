package app_gin

import (
	"datathink.top/Waigo/internal/app/app_gin/gin_controller"
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
	mdw := Middlewares{}
	gcl := gin_controller.GinController{}
	at := AppToken{}

	// 默认路由，api
	// http://127.0.0.1:9850/api/spider/ithome
	route.Any("/api/spider/ithome", mdw.HttpCorsApi, mdw.HttpError500, at.CheckAppToken, gcl.SpiderITHome)

	// http://127.0.0.1:9850/api/get_play_audio_list
	route.Any("/api/get_play_audio_list", mdw.HttpCorsApi, mdw.HttpError500, at.CheckAppToken, gcl.GetPlayAudioList)

}
