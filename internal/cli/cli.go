package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/huh"
	commonModel "github.com/horonlee/ginhub/internal/model/common"
	"github.com/horonlee/ginhub/internal/tui"
)

// DoGinHubInfo 打印 GinHub 信息
func DoGinHubInfo() {
	if _, err := fmt.Fprintln(os.Stdout, tui.GetGinHubInfo()); err != nil {
		fmt.Fprintf(os.Stderr, "failed to print GinHub info: %v\n", err)
	}
}

// DoVersion 打印版本信息
func DoVersion() {
	item := struct{ Title, Msg string }{
		Title: "📦 当前版本",
		Msg:   "v" + commonModel.Version,
	}
	tui.PrintCLIWithBox(item)
}

// DoHello 打印 Ech0 Logo
func DoHello() {
	tui.ClearScreen()
	tui.PrintCLIBanner()
}

// DoTui 执行 TUI
func DoTui() {
	// 清除屏幕当前字符
	tui.ClearScreen()
	// 打印 ASCII 风格 Banner
	tui.PrintCLIBanner()

	for {
		// 换行
		fmt.Println()

		var action string
		var options []huh.Option[string]

		if s == nil {
			options = append(options, huh.NewOption("🚀 启动 Web 服务", "serve"))
		} else {
			options = append(options, huh.NewOption("🛑 停止 Web 服务", "stopserve"))
		}

		options = append(options,
			huh.NewOption("🦖 查看信息", "info"),
			huh.NewOption("📌 查看版本", "version"),
			huh.NewOption("❌ 退出", "exit"),
		)

		err := huh.NewSelect[string]().
			Title("欢迎使用 GinHub TUI .").
			Options(options...).
			Value(&action).
			WithTheme(huh.ThemeCatppuccin()).
			Run()
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "serve":
			tui.ClearScreen()
			DoServe()
		case "stopserve":
			tui.ClearScreen()
			DoStopServe()
		case "info":
			tui.ClearScreen()
			DoGinHubInfo()
		case "version":
			tui.ClearScreen()
			DoVersion()
		case "exit":
			fmt.Println("👋 感谢使用 GinHub TUI，期待下次再见")
			return
		}
	}
}
