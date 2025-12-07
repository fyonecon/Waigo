package bootstrap

import (
	"datathink.top.Waigo/internal/app_tray"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type AppServicesForTray struct{}

// RunTray 状态栏托盘服务
func RunTray(app *application.App, window *application.WebviewWindow) {
	fmt.Println("[Waigo-Log]", "启动Tray")
	at := app_tray.AppTray{}
	at.TrayView(app, window)
}
