package bootstrap

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/app_gin"
	"datathink.top/Waigo/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"time"
)

type AppServicesForGin struct{}

// RouteMust 必要路由
func RouteMust(route *gin.Engine, ginHTML fs.FS, ginFiles fs.FS) {
	// info
	appName := common.InterfaceToString(internal.GetConfigMap("app", "appName"))
	appRights := common.InterfaceToString(internal.GetConfigMap("app", "appRights"))
	appVersion := common.InterfaceToString(internal.GetConfigMap("app", "appVersion"))

	// 中间件
	mdw := app_gin.Middlewares{}

	// 默认路由，api
	// 如：http://127.0.0.1:9850
	route.Any("/", mdw.HttpCorsApi, mdw.HttpError500, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"state": 1,
			"msg":   "YES",
			"content": map[string]interface{}{
				"appName":    appName,
				"appVersion": "v" + appVersion,
				"appRights":  appRights,
			},
		})
		return
	})

	// 默认路由，api
	// 如：http://127.0.0.1:9850/api
	route.Any("/api", mdw.HttpCorsApi, mdw.HttpError500, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"state": 1,
			"msg":   "YES API",
			"content": map[string]interface{}{
				"appName":    appName,
				"appVersion": "v" + appVersion,
				"appRights":  appRights,
			},
		})
		return
	})

	// 默认路由 html（html=file）
	// 如：http://127.0.0.1:9850/html/index.html
	group_html := route.Group("/html", mdw.HttpCorsHTML, mdw.HttpError500)
	group_html.GET("/:filename", func(ctx *gin.Context) {
		filename := ctx.Param("filename")
		rootPath := "ginassets/html/" // 结尾/，开头无/
		filepath := ""                // 完整路径
		//
		if len(filename) == 0 {
			filepath = rootPath + "index.html"
		} else {
			filepath = rootPath + filename
		}
		//
		theFile, err := ginHTML.Open(filepath) // 全路径：如 ginassets/html/favicon.ico
		if err != nil {
			ctx.JSON(404, gin.H{
				"state": 0,
				"msg":   "文件不存在",
				"content": map[string]interface{}{
					"filename":   filepath,
					"appName":    appName,
					"appVersion": "v" + appVersion,
					"appRights":  appRights,
				},
			})
			return
		}
		defer theFile.Close()
		//
		fileType := common.GetFileContentType(filename)
		ctx.Header("Content-Type", fileType)
		io.Copy(ctx.Writer, theFile) // 流式传输
		return
	})

	// 默认路由 file（html=file）
	// 如：http://127.0.0.1:9850/files/test.txt
	group_files := route.Group("/files", mdw.HttpCorsFiles, mdw.HttpError500)
	group_files.GET("/:filename", func(ctx *gin.Context) {
		filename := ctx.Param("filename")
		rootPath := "ginassets/files/"  // 结尾/，开头无/
		filepath := rootPath + filename // 完整路径
		//
		if len(filename) > 0 {
			//
			theFile, err := ginFiles.Open(filepath) // 全路径：如 ginassets/files/favicon.ico
			if err != nil {
				ctx.JSON(404, gin.H{
					"state": 0,
					"msg":   "文件不存在",
					"content": map[string]interface{}{
						"filename":   filepath,
						"appName":    appName,
						"appVersion": "v" + appVersion,
						"appRights":  appRights,
					},
				})
				return
			}
			defer theFile.Close()
			//
			fileType := common.GetFileContentType(filename)
			ctx.Header("Content-Type", fileType)
			io.Copy(ctx.Writer, theFile) // 流式传输
		} else {
			ctx.JSON(404, gin.H{
				"state": 0,
				"msg":   "文件不能为空",
				"content": map[string]interface{}{
					"filename":   filepath,
					"appName":    appName,
					"appVersion": "v" + appVersion,
					"appRights":  appRights,
				},
			})
		}
		return
	})

	// 默认路由，html
	// 如：http://127.0.0.1:9850/favicon.ico
	favicon := "favicon.ico"
	route.Any("/"+favicon, mdw.HttpCorsHTML, mdw.HttpError500, func(ctx *gin.Context) {
		fileType := common.GetFileContentType(favicon)
		//
		theFile, err := ginHTML.Open("ginassets/html/" + favicon) // 全路径：如 ginassets/html/favicon.ico
		if err != nil {                                           // 有文件
			ctx.JSON(404, gin.H{
				"state": 0,
				"msg":   "文件不存在",
				"content": map[string]interface{}{
					"filename":   favicon,
					"appName":    appName,
					"appVersion": "v" + appVersion,
					"appRights":  appRights,
				},
			})
			return
		}
		defer theFile.Close()
		//
		ctx.Header("Content-Type", fileType)
		io.Copy(ctx.Writer, theFile) // 流式传输
		return
	})

	// 默认路由，404
	route.NoRoute(mdw.HttpCorsApi, mdw.HttpError500, func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"state": 404,
			"msg":   "404",
			"content": map[string]interface{}{
				"appName":    appName,
				"appVersion": "v" + appVersion,
				"appRights":  appRights,
			},
		})
		return
	})
}

// RunGin Http服务
func RunGin(app *application.App, window *application.WebviewWindow, ginHTML fs.FS, ginFiles fs.FS) {
	fmt.Println("[Waigo-Log]", "启动Gin")

	// port
	ginPort := common.InterfaceToString(internal.GetConfigMap("gin", "ginPort"))
	ginDebug := common.InterfaceToString(internal.GetConfigMap("gin", "ginDebug"))

	// Gin日志
	if ginDebug == "off" {
		fmt.Println("ginDebug已关，不显示Web日志")
		gin.SetMode(gin.ReleaseMode)
		gin.DefaultWriter = ioutil.Discard
	} else {
		fmt.Println("Debug已开，显示Web日志")
		gin.SetMode(gin.DebugMode)
		fmt.Println("测试-File文件访问：", "http://127.0.0.1:"+ginPort+"/files/"+"test.txt")
	}
	// 创建Gin
	httpServer := gin.Default()

	// 限制最大文件缓存体积
	var maxFileSize int64 = common.InterfaceToInt(internal.GetConfigSetter("max_file_size"))
	httpServer.MaxMultipartMemory = maxFileSize * 1024 << 20

	// 注册必要路由
	RouteMust(httpServer, ginHTML, ginFiles)

	// 注册自定义路由
	agn := app_gin.AppGin{}
	agn.RouteAPI(httpServer, ginHTML, ginFiles)
	agn.RouteHTML(httpServer, ginHTML, ginFiles)
	agn.RouteFile(httpServer, ginHTML, ginFiles)

	// 启动
	err := httpServer.Run(":" + ginPort)
	if err != nil {
		fmt.Println("Error: 127.0.0.1:" + ginPort + " 开启失败，可能是 Gin服务已启用 或 Gin端口冲突 。")
		time.Sleep(1 * time.Second)
		app.Quit() // 服务不能启动就直接退出所有程序（防止重复开启软件）
	}

}
