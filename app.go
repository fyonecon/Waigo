package main

import (
	"datathink.top/Waigo/internal/bootstrap"
	"datathink.top/Waigo/internal/common/kits"
	"embed"
	_ "embed"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

// 视图UI
//
//go:embed all:frontend/dist
var assets embed.FS

// 配套的Gin
//
//go:embed all:ginassets/html
var ginHTML embed.FS

//go:embed all:ginassets/files
var ginFiles embed.FS

// InitSYS 先检查系统运行环境
func InitSYS() {
	state, msg, content := bootstrap.InitCheckSYS()
	if state == 1 {
		InitApp()
	} else {
		fmt.Println("\nXXX 系统自检失败 => ", msg, content)
		// alert
	}
}

// InitApp 先创建主程序App，再创建子程序window，所以其它服务建议在window创建后再注册。这是为了适配Mac下主程序唯一性原则，同时适配的其它平台。
func InitApp() {
	// 主程序App
	app := application.New(application.Options{
		Name:        "Waigo",
		Description: "This is Waigo.",
		Icon:        kits.IconData,
		//
		Services: []application.Service{
			application.NewService(&bootstrap.AppServicesForWindow{}),
		},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
		Windows: application.WindowsOptions{},
		Linux:   application.LinuxOptions{},
		//OnStartup: func(ctx context.Context) {
		// 此方法 Wails (v3.0.0-alpha.41) 里无，但是文档有
		//},
		OnShutdown: func() {
			fmt.Println("[Waigo-Log]", "关闭了主程序。。。")
		},
	})

	// 注册子程序window
	bootstrap.InitWindow(app, ginHTML, ginFiles)

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("[Waigo-Log]", "主动结束了主进程App。。。")
	}
}
