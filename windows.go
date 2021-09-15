package main

import (
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewDelayWindow(a fyne.App, w float32, h float32, sec time.Duration) {
	sec = time.Second * sec
	time.Sleep(sec)

	win := a.NewWindow("Shown later")
	win.SetContent(widget.NewLabel("5 seconds later"))
	win.Resize(fyne.NewSize(w, h))
	win.Show()
}

// NewLayoutExample
// 基於 windows 設定 layout
func NewLayoutExample(mainWindow fyne.Window) {
	tt := container.New(
		layout.NewVBoxLayout(),
		//layout.NewSpacer(),
		widget.NewLabel("111"),
		widget.NewButton("顯示PID", func() {
			dialog.ShowInformation("確認", "PID:"+strconv.FormatInt(int64(os.Getpid()), 10), mainWindow)
		}),
		widget.NewButton("顯示PID", func() {
			mainWindow.SetContent(widget.NewLabel("5 seconds later"))
			mainWindow.Resize(fyne.NewSize(200, 200))
			mainWindow.Show()
		}),
		layout.NewSpacer(),
	)
	mainWindow.SetContent(tt)
}

// NewTabsExample
// 基於 windows 設定 tabs
func NewTabsExample(mainWindow fyne.Window) {
	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)
	tabs.Append(container.NewTabItem("test", widget.NewLabel("Home tab")))
	tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationTop)
	mainWindow.SetContent(tabs)
}
