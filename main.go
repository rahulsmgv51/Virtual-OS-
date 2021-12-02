package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("LRC SET OC")
var btn1 fyne.Widget
var btn2 fyne.Widget
var img fyne.CanvasObject;
var DeskBtn fyne.Widget

var panelContent *fyne.Container


func main(){
	//myApp.Settings().SetTheme(theme.LightTheme())
	img = canvas.NewImageFromFile("C:\\Users\\DELL\\Desktop\\LRC\\LRC_SET_OS\\dp.jpg")

	btn1 = widget.NewButtonWithIcon("News App", theme.InfoIcon(), func ()  {
		showNewsApp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Media Player", theme.HelpIcon(), func ()  {
		showMediaApp(myWindow)
	})
	
	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})
	panelContent = container.NewVBox(container.NewGridWithColumns(3, DeskBtn,btn1, btn2))
	myWindow.Resize(fyne.NewSize(1200, 720))
	myWindow.CenterOnScreen()
	myWindow.SetContent(container.NewBorder(panelContent, nil, nil,nil, img))
	myWindow.ShowAndRun()

}