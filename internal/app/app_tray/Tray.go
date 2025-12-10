package app_tray

import (
	"datathink.top/Waigo/internal/common/kits"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// TrayView Tray
func (at *AppTray) TrayView(app *application.App, window *application.WebviewWindow) {
	//
	systray := app.SystemTray.New()
	systray.SetIcon(kits.ICON)
	systray.SetLabel("Waigo")
	systray.OnClick(func() {
		window.Show()
		window.Focus()
		//
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
}
