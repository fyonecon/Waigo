package bootstrap

import (
	"datathink.top.Waigo/internal/app_tray"
	"datathink.top.Waigo/internal/app_window"
	"datathink.top.Waigo/internal/common/kits"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type AppServicesForWindow struct{}

// JSCallGo js调用go
func (asw *AppServicesForWindow) JSCallGo(key string, dataDict map[string]interface{}) map[string]interface{} {
	var awd = app_window.AppWindow{}
	return awd.ListJSCallGo(key, dataDict)
}

// InitWindow 运行子程序视窗
func InitWindow(app *application.App) {
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
		RunGin(app, window)
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
