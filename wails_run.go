package main

import (
	"datathink.top.Waigo/internal"
	"embed"
	_ "embed"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"log"
	"time"
)

// Wails uses Go's `embed` package to embed the frontend files into the binary.
// Any files in the frontend/dist folder will be embedded into the binary and
// made available to the frontend.
// See https://pkg.go.dev/embed for more information.

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	// Register a custom event whose associated data type is string.
	// This is not required, but the binding generator will pick up registered events
	// and provide a strongly typed JS/TS API for them.
	application.RegisterEvent[string]("time")
}

func RunWails() {
	// Create a new Wails application by providing the necessary options.
	// Variables 'Name' and 'Description' are for application metadata.
	// 'Assets' configures the asset server with the 'FS' variable pointing to the frontend files.
	// 'Bind' is a list of Go struct instances. The frontend has access to the methods of these instances.
	// 'Mac' options tailor the application when running an macOS.
	app := application.New(application.Options{
		Name:        "Waigo",
		Description: "A demo of using raw HTML & CSS",
		Icon:        internal.ICON,
		//
		Services: []application.Service{
			application.NewService(&WailsServices{}),
			application.NewService(TestGin()),
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
		//	var err error
		//	db, err = sql.Open("sqlite3", "app.db")
		//	if err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	// Run migrations
		//	if err := runMigrations(db); err != nil {
		//		log.Fatal(err)
		//	}
		//},
		OnShutdown: func() {
			fmt.Println("on shutdown")
		},
	})

	// Create a new window with the necessary options.
	// 'Title' is the title of the window.
	// 'Mac' options tailor the window when running on macOS.
	// 'BackgroundColour' is the background colour of the window.
	// 'URL' is the URL that will be loaded into the webview.
	window := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Name:             "master-window",
		Title:            "Master Window",
		URL:              "/",
		BackgroundColour: application.NewRGB(27, 38, 54),
		// Size and Position
		Width:     1000,
		Height:    800,
		MinWidth:  800,
		MinHeight: 600,
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
			Icon: internal.ICON,
		},
		//
	})
	window.Center()
	window.Show()

	// Tray
	go func() {
		//
		systray := app.SystemTray.New()
		systray.SetIcon(internal.ICON)
		systray.SetLabel("Waigo")
		systray.OnClick(func() {
			window.Show()
			window.Focus()
		})
		//
		menu := app.NewMenu()
		menu.Add("Show").OnClick(func(ctx *application.Context) {
			window.Show()
			window.Focus()
			//
		})
		menu.Add("Quit").OnClick(func(ctx *application.Context) {
			app.Quit()
		})
		systray.SetMenu(menu)
	}()

	// Create a goroutine that emits an event containing the current time every second.
	// The frontend can listen to this event and update the UI accordingly.
	go func() {
		for {
			now := time.Now().Format(time.RFC1123)
			app.Event.Emit("time", now)
			time.Sleep(time.Second)
		}
	}()

	//
	app.Window.OnCreate(func(window application.Window) {
		fmt.Println("OnCreate================================")
	})

	// Run the application. This blocks until the application has been exited.
	err := app.Run()

	// If an error occurred while running the application, log it and exit.
	if err != nil {
		log.Fatal(err)
	}
}
