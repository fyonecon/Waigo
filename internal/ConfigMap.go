package internal

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

// 全局变量
var RUNNNINGID = ""           // 本次软件启动携带的唯一性ID
var TIMEINTERVALNUM int64 = 0 // 周期计数
var APP *application.App
var WINDOW *application.WebviewWindow

func GetConfigMap(group string, key string) interface{} {
	//
	var inter = ConfigMap[group]
	if inter != nil {
		config := inter.(map[string]interface{}) // interface转map
		return config[key]
	} else {
		return nil
	}
}

// ConfigMap 全局配置（固定）参数。只读。静态数据。
// appConfig := common.InterfaceToMap(internal.ConfigMap["app"]); appClass := common.InterfaceToString(appConfig["appClass"])
var ConfigMap map[string]interface{} = map[string]interface{}{
	"app": map[string]interface{}{
		"appName":    "Waigo",         // AppName，推荐使用英文
		"appClass":   "Waigo_",        // 必须为string，且必须唯一，推荐使用英文
		"appRights":  "Datathink.Top", // AppRights，也是软件安装默认路径
		"appVersion": "1.6.4",         // AppVersion
	},
	"sys": map[string]interface{}{
		"appStateStartTime": 20251201010101,        // 最早时间 YmdHis
		"appStateEndTime":   20341201010101,        // 截止时间，一个版本：9年。（这是软件及扩展更新的要求）
		"cachePath":         "top.datathink.Waigo", // 缓存文件夹，结尾无/
		"dataPath":          "top.datathink.Waigo", // 数据持久化文件夹，结尾无/
		"dataPathDirsName": []string{
			"running", "local_database", "gin_ssl", "user", // 必要
			// 其它
		},
	},
	"gin": map[string]interface{}{
		"ginDebug":        "on",     // on 显示Gin日志，off 不显示Gin日志
		"ginPort":         "9850",   // 内网+外网端口（非wails服务），默认：9850
		"maxFileSize":     "32",     // 最大上传文件体积，GB
		"windowTokenSalt": "NbPlus", //
		"whiteHosts": []string{
			"http://127.0.0.1",
			"https://127.0.0.1",
		}, // 白名单域名或IP，格式：协议+IPv4+port、协议+域名
		"view_url": "http://127.0.0.1:9850", // 生产环境：api主网址或额外视图网址（协议+网址+端口+路径，如：http(s)://127.0.0.1:port ）。注意ssl是否已开启。
		"ssl":      "OFF",                   // "ON"、"OFF"，是否开启ssl
	},
	"wails": map[string]interface{}{
		"debug": "off",
	},
}
