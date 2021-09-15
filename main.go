package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"time"
)

const appTitle = "daemon"

var AnotherWindow = true
var WindowToContainer = true

func main() {
	mainApp := setUp()
	mainWindow := func(fyne.App) fyne.Window {
		w := mainApp.NewWindow(appTitle)
		w.Resize(fyne.NewSize(300, 200))
		return w
	}(mainApp)

	if WindowToContainer {
		NewTabsExample(mainWindow)
		//NewMainWindow(mainWindow)
	}

	if AnotherWindow {
		go NewDelayWindow(mainApp, 200, 200, 5)
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
