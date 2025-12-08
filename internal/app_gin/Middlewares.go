package app_gin

import (
	"datathink.top.Waigo/internal"
	"datathink.top.Waigo/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

type Middlewares struct{}

// CheckRequestApi 验证请求
func (mdw *Middlewares) CheckRequestApi(ctx *gin.Context) {
	webServerPort := common.InterfaceToString(internal.ConfigMap["webServerPort"])
	webServerKey := common.InterfaceToString(internal.ConfigMap["webServerKey"])
	href := ctx.Request.Host + ctx.Request.RequestURI // 完整网址

	if common.StringHasString(href, "127.0.0.1:"+webServerPort+"/"+webServerKey) != -1 { // 校验href，仅限电脑内网访问，拦截广域网和局域网
		ctx.Next()
	} else {
		ctx.String(404, "500-api")
		ctx.Abort()
		return
	}
}

// CheckRequestFiles 验证请求
func (mdw *Middlewares) CheckRequestFiles(ctx *gin.Context) {
	webServerPort := common.InterfaceToString(internal.ConfigMap["webServerPort"])
	webServerKey := common.InterfaceToString(internal.ConfigMap["webServerKey"])
	webServerGetFile := common.InterfaceToString(internal.ConfigMap["webServerGetFile"])
	webServerGetFileDirName := common.InterfaceToString(internal.ConfigMap["webServerGetFileDirName"])
	href := ctx.Request.Host + ctx.Request.RequestURI // 完整网址

	okHref := "127.0.0.1:" + webServerPort + "/" + webServerKey + "/" + webServerGetFile + webServerGetFileDirName
	if common.StringHasString(href, okHref) != -1 { // 校验href，仅限电脑内网访问，拦截广域网和局域网
		ctx.Next()
	} else {
		ctx.String(404, "500-files")
		ctx.Abort()
		return
	}
}

// CheckRequestHTML 验证请求
func (mdw *Middlewares) CheckRequestHTML(ctx *gin.Context) {
	webServerPort := common.InterfaceToString(internal.ConfigMap["webServerPort"])
	webServerLoadHTML := common.InterfaceToString(internal.ConfigMap["webServerLoadHTML"])
	webServerLoadHTMLDirName := common.InterfaceToString(internal.ConfigMap["webServerLoadHTMLDirName"])
	href := ctx.Request.Host + ctx.Request.RequestURI // 完整网址

	okHref := "" + webServerPort + "/" + webServerLoadHTML + webServerLoadHTMLDirName
	if common.StringHasString(href, okHref) != -1 { // 校验href，不拦截网址
		ctx.Next()
	} else {
		ctx.String(404, "500-html")
		ctx.Abort()
		return
	}
}

// HttpError500 抛出500错误
func (mdw *Middlewares) HttpError500(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			//打印错误堆栈信息
			fmt.Printf("\n Http panic: %v \n\n", err)

			pc := make([]uintptr, 8)
			runtime.Callers(2, pc)
			//f := runtime.FuncForPC(pc[0])

			_fc0 := runtime.FuncForPC(pc[0]).Name()
			_fc1 := runtime.FuncForPC(pc[1]).Name()
			_fc2 := runtime.FuncForPC(pc[2]).Name()
			_fc3 := runtime.FuncForPC(pc[3]).Name()
			_fc4 := runtime.FuncForPC(pc[4]).Name()
			_fc5 := runtime.FuncForPC(pc[5]).Name()
			_fc6 := runtime.FuncForPC(pc[6]).Name()

			errorFunc1 := gin.H{
				"0": _fc0,
				"1": _fc1,
				"2": _fc2,
				"3": _fc3,
				"4": _fc4,
				"5": _fc5,
				"6": _fc6,
			}
			fmt.Println("代码500错误：", errorFunc1)

			// 返回
			ctx.JSON(500, gin.H{
				"state": 0,
				"msg":   "500",
				"content": map[string]interface{}{
					"errorFunc1": errorFunc1,
				},
			})
			ctx.Abort()
		}
	}()
	//加载完 defer recover，继续后续接口调用并返回JSON提示
	ctx.Next()
}

// HttpCorsApi 处理http-header信息
func (mdw *Middlewares) HttpCorsApi(ctx *gin.Context) { // 面向Api
	method := ctx.Request.Method
	//
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//
	ctx.Header("Content-type", "application/json; charset=utf-8")
	ctx.Header("Cache-Control", "max-age=20")
	//
	ctx.Header("Framework-Author", "fyonecon")
	ctx.Header("Date", common.GetTimeDate("Y-m-d H:i:s"))
	//是否允许后续请求携带认证信息,该值只能是true,否则不返回
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Content-Disposition", "inline")
	// 二请
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	ctx.Next()
}

// HttpCorsFiles 处理http-header信息
func (mdw *Middlewares) HttpCorsFiles(ctx *gin.Context) { // 面向Api
	method := ctx.Request.Method
	//
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Cache-Control", "max-age=1200")
	ctx.Header("Framework-Author", "fyonecon")
	ctx.Header("Date", common.GetTimeDate("Y-m-d H:i:s"))
	//是否允许后续请求携带认证信息,该值只能是true,否则不返回
	ctx.Header("Access-Control-Allow-Credentials", "true")
	//ctx.Header("Content-Disposition", "attachment")
	ctx.Header("Content-Disposition", "inline")
	// 二请
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	ctx.Next()
}

// HttpCorsHTML 处理http-header信息
func (mdw *Middlewares) HttpCorsHTML(ctx *gin.Context) { // 面向Api
	method := ctx.Request.Method
	//
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ctx.Header("Cache-Control", "max-age=120")
	ctx.Header("Framework-Author", "fyonecon")
	ctx.Header("Date", common.GetTimeDate("Y-m-d H:i:s"))
	//是否允许后续请求携带认证信息,该值只能是true,否则不返回
	ctx.Header("Access-Control-Allow-Credentials", "true")
	ctx.Header("Content-Disposition", "inline")
	// 二请
	if method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	ctx.Next()
}
