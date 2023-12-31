package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	StartHttp()
	// Create an instance of the app structure
	app := NewApp()
	// Start http server
	go StartHttp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:             "wails-app",
		Width:             1024,
		Height:            768,
		StartHidden:       false, // 启动时隐藏窗口
		HideWindowOnClose: false, // 关闭时隐藏窗口
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false, // 网页透明
			WindowIsTranslucent:  true,  // 窗口半透明
			DisableWindowIcon:    false, // 禁用窗口图标 true将删除标题栏左上角的图标
			Theme:                windows.SystemDefault,
		},
		Mac: &mac.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "测试工具 ",
				Message: "© 2023 pltrueover@gamil.com",
			},
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
		Frameless:       false, //无边框应用
		CSSDragProperty: "widows",
		CSSDragValue:    "1",
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
