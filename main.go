package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

//var
var a=app.New()
//w:=a.NewWindow("ok")
var myWindow = a.NewWindow("LRC OS")
var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var img  fyne.CanvasObject
var deskBtn  fyne.Widget
var panelContent *fyne.Container
func main()  {
	a.Settings().SetTheme(theme.DarkTheme())
	img:=canvas.NewImageFromFile("C:\\Users\\user\\Desktop\\test case\\dice\\desktop.jpg")
	//w1:=canvas.NewImageFromFile()
	weatherIcon,_:=fyne.LoadResourceFromPath("C:\\Users\\user\\Desktop\\test case\\dice\\weather.png")
	galleryIcon,_:=fyne.LoadResourceFromPath("C:\\Users\\user\\Desktop\\test case\\dice\\gallery.jpg")
	notepadIcon,_:=fyne.LoadResourceFromPath("C:\\Users\\user\\Desktop\\test case\\dice\\notepad.jpg")
	calculatorIcon,_:=fyne.LoadResourceFromPath("C:\\Users\\user\\Desktop\\test case\\dice\\calc.png")
	homeIcon,_:=fyne.LoadResourceFromPath("C:\\Users\\user\\Desktop\\test case\\dice\\home.png")
	btn1=widget.NewButtonWithIcon("Weather App",weatherIcon, func() {
			showWeatherApp(myWindow)
	})
	//calcIcon,_:=fyne.LoadResourceFromPath("C:\\Users\\user\\Desktop\\test case\\dice\\calc.png")
	btn2=widget.NewButtonWithIcon("Calculator",calculatorIcon, func() {
		showCalculatorApp()
	})
	//ikon,_:=fyne.LoadResourceFromPath("C:\\Users\\user\\Desktop\\test case\\dice\\images.jpg")

	btn3=widget.NewButtonWithIcon("Image Viewer",galleryIcon, func() {
		showImageViewerApp(myWindow)
	})
	btn4=widget.NewButtonWithIcon("Text Editor",notepadIcon, func() {
		showTextEditorApp()
	})
	deskBtn=widget.NewButtonWithIcon("Home", homeIcon,func() {
		myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,nil,img))
	})
	panelContent=container.NewGridWithColumns(5,deskBtn,btn1,btn2,btn3,btn4)
	myWindow.Resize(fyne.NewSize(1280,720))
	myWindow.CenterOnScreen()
	myWindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	myWindow.ShowAndRun()
}

