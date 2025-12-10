package app_gin

import (
	"datathink.top/Waigo/internal/app/app_gin/gin_controller"
	"github.com/gin-gonic/gin"
	"io/fs"
)

// RouteFile file专用路由
func (agn *AppGin) RouteFile(route *gin.Engine, ginHTML fs.FS, ginFiles fs.FS) {
	// info
	//appName := common.InterfaceToString(internal.GetConfigMap("app", "appName"))
	//appRights := common.InterfaceToString(internal.GetConfigMap("app", "appRights"))
	//appVersion := common.InterfaceToString(internal.GetConfigMap("app", "appVersion"))
	// 中间件
	mdw := Middlewares{}
	//
	gcl := gin_controller.GinController{}
	// 文件
	group := route.Group("/", mdw.HttpCorsFiles, mdw.HttpError500)
	group.POST("upload_files", gcl.UploadFile)
	//group.Any("get_file", http.FS(ginFiles))

}
