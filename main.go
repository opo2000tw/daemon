package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const appTitle = "daemon"

var AnotherWindow = true

func main() {
	mainApp := setUp()
	mainWindow := func(fyne.App) fyne.Window {
		w := mainApp.NewWindow(appTitle)
		w.Resize(fyne.NewSize(300, 200))
		return w
	}(mainApp)

	mainWindow.SetContent(
		container.New(
			layout.NewVBoxLayout(),
			//layout.NewSpacer(),
			widget.NewLabel("111"),
			widget.NewButton("顯示PID", func() {
				dialog.ShowInformation("確認", "PID:"+strconv.FormatInt(int64(os.Getpid()), 10), mainWindow)
			}),
			layout.NewSpacer(),
		),
	)

	if AnotherWindow {
		go setAnotherWin(mainApp)
	}

	mainWindow.ShowAndRun()
	tearDown()
}

func setAnotherWin(a fyne.App) {
	time.Sleep(time.Second * 5)

	win := a.NewWindow("Shown later")
	win.SetContent(widget.NewLabel("5 seconds later"))
	win.Resize(fyne.NewSize(200, 200))
	win.Show()
}

func setUp() fyne.App {
	myApp := app.New()
	myApp.Settings().SetTheme(&myTheme{})
	return myApp
}

func tearDown() {
	fmt.Println("Exited")
}
