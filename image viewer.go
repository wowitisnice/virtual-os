package main

import (
	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"io/ioutil"
	"log"
	"strings"
)

func showImageViewerApp(w fyne.Window)  {
	//a:=app.New()
	//w:=a.NewWindow("image viewer")
	a.Settings().SetTheme(theme.DarkTheme())
	w.Resize(fyne.NewSize(1000,600))
	tabs := container.NewAppTabs()
	source:="C:\\Users\\user\\Pictures"
	files, err := ioutil.ReadDir(source)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		var fileName=file.Name()
		if file.IsDir()==false {
			ext := strings.Split(fileName, ".")[1]
			if ext == "png" || ext == "jpeg" {
				image:=canvas.NewImageFromFile(source+"\\"+file.Name())
				image.FillMode=canvas.ImageFillStretch
				tabs.Append(container.NewTabItem(file.Name(),image))
			}
		}
	}

	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewBorder(panelContent,nil,nil,nil,tabs))
	w.Show()
} 