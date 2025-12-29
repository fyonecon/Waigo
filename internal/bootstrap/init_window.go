package bootstrap

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/app/app_services"
	"datathink.top/Waigo/internal/app/app_tray"
	"datathink.top/Waigo/internal/app/app_window"
	"datathink.top/Waigo/internal/app/app_window/window_controller"
	"datathink.top/Waigo/internal/common/kits"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"io/fs"
	"time"
)

type AppServicesForWindow struct{}

// JSCallGo js调用go
// web端 runtime或window 需要关闭web的ssr，是能是browser模式下才能用
func (asw *AppServicesForWindow) JSCallGo(key string, dataDict map[string]interface{}) map[string]interface{} {
	//fmt.Println("JSCallGo=")
	var awd = window_controller.WindowController{}
	return awd.ListJSCallGo(key, dataDict)
}

// Test js调用go
func (asw *AppServicesForWindow) Test() string {
	return "Test"
}

// InitWindow 运行子程序视窗
func InitWindow(app *application.App, ginHTML fs.FS, ginFiles fs.FS) {
	fmt.Println("[Waigo-Log]", "启动主视窗的其它服务")

	// 子程序window
	window := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:      "master-window",
		Title:     "Master Window",
		Width:     960, // 960
		Height:    700, // 700
		MinWidth:  520,
		MinHeight: 520,
		//
		BackgroundColour: application.NewRGB(55, 55, 55),
		//
		Mac: application.MacWindow{
			InvisibleTitleBarHeight: 50,
			Backdrop:                application.MacBackdropTranslucent,
			TitleBar:                application.MacTitleBarHiddenInset,
			//LiquidGlass:             true,
		},
		Windows: application.WindowsWindow{
			DisableIcon:                       false,
			DisableFramelessWindowDecorations: false,
		},
		Linux: application.LinuxWindow{
			Icon: kits.IconData,
		},
		//
	})

	// 设置成全局变量
	internal.APP = app
	internal.WINDOW = window

	// window
	go func() {
		aw := app_window.AppWindow{}
		aw.WindowView(app, window)
	}()

	// Tray
	go func() {
		at := app_tray.AppTray{}
		at.TrayView(app, window)
	}()

	// Gin
	go func() {
		RunGin(app, window, ginHTML, ginFiles)
	}()

	// 周期服务
	go func() {
		// 启动周期定时器服务
		ticker := time.NewTicker(6 * time.Second) // 周期12s（默认10s，建议在[2,3,4,5,6,10,12,15,20]这些数值调。此处如果调整了，则CycleEventStateXXi00s(i string)函数里面的判断也必须调整。）
		asv := app_services.AppServices{}
		for _ = range ticker.C {
			asv.StartTimeInterval(internal.TIMEINTERVALNUM)
			internal.TIMEINTERVALNUM++
		}
	}()

	// Emit传递参数到前端js，前端必须接收
	//go func() {
	//	var key = "timer"
	//	application.RegisterEvent[string](key)      // 注册
	//	app.Event.Emit(key, map[string]interface{}{ // 传递
	//		"key": key,
	//		"data_dict": map[string]interface{}{
	//			"test": "12345",
	//		},
	//	})
	//}()
	//// Emit传递参数到前端js
	//let key = "timer"
	//Events.On(key, (result) => {
	//  console.log("GoEmitToJS="+key+"=", result.data);
	//});

}
