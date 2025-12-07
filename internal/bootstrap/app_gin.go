package bootstrap

import (
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

type AppServicesForGin struct{}

// RunGin Http服务
func RunGin(app *application.App, window *application.WebviewWindow) {
	fmt.Println("[Waigo-Log]", "启动Gin")

}
