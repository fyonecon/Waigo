package app_gin

import (
	"datathink.top.Waigo/internal/app_gin/gin_controller"
	"github.com/gin-gonic/gin"
	"io/fs"
)

// RouteFile file专用路由
func (agn *AppGin) RouteFile(route *gin.Engine, ginHTML fs.FS, ginFiles fs.FS) {
	mdw := Middlewares{}
	gcl := gin_controller.GinController{}
	// 文件
	group := route.Group("/", mdw.HttpCorsFiles, mdw.HttpError500)
	group.POST("upload_files", gcl.UploadFile)
	//group.Any("get_file", http.FS(ginFiles))

}
