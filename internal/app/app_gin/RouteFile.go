package app_gin

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/app/app_gin/gin_controller"
	"datathink.top/Waigo/internal/common"
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
	at := AppToken{}
	// 文件
	group := route.Group("/do_file", mdw.HttpCorsFiles, mdw.HttpError500)
	group.POST("/upload_files", gcl.UploadFile)
	//group.Any("get_file", http.FS(ginFiles))

	// 访问完整路径下的文件
	group2 := route.Group("/dir", mdw.HttpCorsFiles, mdw.HttpError500, at.CheckAppToken)
	group2.GET("/play_audio/*filepath", func(ctx *gin.Context) {
		filepath := ctx.Param("filepath")
		//
		//fileTokenState := common.RequestInput(ctx, "file_token") == common.MD5("filetoken#@"+common.URIDecode(filepath))
		//
		_apiURL := common.RequestFullURL(ctx) // 完整访问域名
		whiteHosts := internal.GetConfigMap("gin", "whiteHosts")
		apiURLState := common.ArrayInString(common.InterfaceToArrayString(whiteHosts), _apiURL) != -1
		//
		if apiURLState {
			if len(filepath) > 0 {
				if common.IsFile(filepath) {
					fileType := common.GetFileContentType(filepath)
					filename := common.GetFilename(filepath)
					ctx.Header("Content-Type", fileType)
					ctx.Header("filename", filename)
					ctx.File(filepath)
				} else {
					ctx.JSON(404, gin.H{
						"state": 0,
						"msg":   "文件不存在, 2",
						"content": map[string]interface{}{
							"filepath": filepath,
						},
					})
				}
			} else {
				ctx.JSON(404, gin.H{
					"state": 0,
					"msg":   "文件不能为空, 2",
					"content": map[string]interface{}{
						"filepath": filepath,
					},
				})
			}
		} else {
			ctx.JSON(200, gin.H{
				"state": 0,
				"msg":   "Auth Error",
				"content": map[string]interface{}{
					"filepath": filepath,
				},
			})
		}
		return
	})

}
