package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/fredbi/uri"
	"io/ioutil"
	"strconv"
)
var count=1
func showTextEditorApp()  {
	//a:=app.New()
	w:=a.NewWindow("Text Editor")
	a.Settings().SetTheme(theme.DarkTheme())
	w.Resize(fyne.NewSize(400,400))
	label1:=widget.NewLabel("Text Editor")
	entry:=widget.NewMultiLineEntry()
	entry.SetPlaceHolder("enter text here")
	filesContainer:=container.NewVBox()
	filesContainer.Add(widget.NewButton("Add Files", func() {
		label:=widget.NewLabel("New File"+strconv.Itoa(count))
		filesContainer.Add(label)
		count++
	}))
	saveBtn:=widget.NewButton("Save Text File", func() {
		saveFileDialog:=dialog.NewFileSave(func(uc fyne.URIWriteCloser, err error) {
			textData:=[]byte(entry.Text)
			uc.Write(textData)
		},w)
		saveFileDialog.SetFileName("New File "+strconv.Itoa(count-1)+".txt")
		saveFileDialog.Show()
	})
	openBtn:=widget.NewButton("Open Text File", func() {
		openFileDialog:=dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
			textData,_:=ioutil.ReadAll(uc)
			viewData:=widget.NewMultiLineEntry()
			output:=fyne.NewStaticResource("New File",textData)
			viewData.SetText(string(output.StaticContent))
			jk:=uc.URI()
			//lm:=jk.String()
			//fmt.Println(jk)
			//dialog.ShowFileSave()
			w:=fyne.CurrentApp().NewWindow(output.StaticName)
			saveBtn:=widget.NewButton("Save Text File", func() {
				saveFileDialog:=dialog.NewFileSave(func(r fyne.URIWriteCloser, err error) {
					//r=lm
					textData:=[]byte(viewData.Text)
					r.Write(textData)
				//w.Close()
				},w)
				//list, _ := storage.ListerForURI(storage.NewFileURI())
				//saveFileDialog.SetLocation(list)
				//fmt.Println(jk.Path())
				saveFileDialog.SetFileName(jk.Name())
				saveFileDialog.Show()
				//saveFileDialog.
			})
			//likhna:=container.NewScroll(viewData)
			w.SetContent(container.NewVBox(viewData,saveBtn))
			w.Resize(fyne.NewSize(400,400))
			w.Show()
		},w)
		openFileDialog.SetFilter(storage.NewExtensionFileFilter([] string{".txt"}))
		openFileDialog.Show()
	})
	Hcontainer:=container.NewHBox(saveBtn,openBtn)
	content:=container.NewVBox(deskBtn,label1,filesContainer,entry,Hcontainer)

	w.SetContent(content)
	w.Show()
}