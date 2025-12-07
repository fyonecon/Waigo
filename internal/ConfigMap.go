package internal

// ConfigMap 全局配置（固定）参数。只读。静态数据。
// 读取common.InterfaceToString(internal.ConfigMap["xx"]["xxx"])
var ConfigMap map[string]interface{} = map[string]interface{}{
	"app": map[string]interface{}{
		"appName":    "Waigo",         // AppName，推荐使用英文
		"appClass":   "Waigo_",        // 必须为string，且必须唯一，推荐使用英文
		"appRights":  "Datathink.Top", // AppRights，也是软件安装默认路径
		"appVersion": "1.0.0",         // AppVersion
	},
	"sys": map[string]interface{}{
		"appStateStartTime": 20251201010101,                // 最早时间
		"appStateEndTime":   20321201010101,                // 截止时间，一个版本：7年
		"cachePath":         "Datathink.Top/Wailgo/Cache/", // 缓存文件夹，必须/结尾
		"localPath":         "Datathink.Top/Wailgo/Files/", // 缓存文件夹，必须/结尾
	},
	"gin": map[string]interface{}{
		"ginDebug":                 "off",            // on 显示Gin日志，off 不显示Gin日志
		"webServerPort":            "9850",           // 内网端口，默认：9850
		"webServerKey":             "",               // 内网路由分组名
		"webServerGetFileDirName":  "/assets/files/", // 远程调用文件夹路径名，必须/结尾
		"webServerLoadHTMLDirName": "/assets/html/",  // 远程调用HTML路径名，必须/结尾
	},
	"wails": map[string]interface{}{
		"debug": "off",
	},
}
