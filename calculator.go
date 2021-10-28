package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
	"strconv"
)
func showCalculatorApp() {
	//a:=app.New()
	//a.Settings().SetTheme(theme.DarkTheme())
	w:=a.NewWindow("Calculator")
	w.Resize(fyne.NewSize(500,300))
	var output=""
	var historyStr=""
	history:=widget.NewLabel("")
	var historyArr []string
	input:=widget.NewLabel(output)
	flag:=false
	historyBtn:=widget.NewButton("history", func() {
		historyStr=""
		if !flag {
			for i := 0; i < len(historyArr); i++ {
				historyStr += historyArr[i] + "\n"
			}
		}
		flag=!flag
		history.SetText(historyStr)

	})
	backBtn:=widget.NewButton("back",func(){
		if len(output)!=0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}
	})
	clearBtn:=widget.NewButton("clear", func() {
		output=""
		input.SetText(output)
	})
	openBtn:=widget.NewButton("(", func() {
		output+="("
		input.SetText(output)
	})
	closingBtn:=widget.NewButton(")", func() {
		output+=")"
		input.SetText(output)
	})
	divideBtn:=widget.NewButton("/", func() {
		output+="/"
		input.SetText(output)
	})
	multiplyBtn:=widget.NewButton("*", func() {
		output+="*"
		input.SetText(output)
	})
	minusBtn:=widget.NewButton("-", func() {
		output+="-"
		input.SetText(output)
	})
	plusBtn:=widget.NewButton("+", func() {
		output+="+"
		input.SetText(output)
	})
	equalsBtn:=widget.NewButton("=", func() {
		expression, err := govaluate.NewEvaluableExpression(output)
		historyStr+=output+"="
		if err==nil{
		result, err := expression.Evaluate(nil)
		if err==nil{
				output=strconv.FormatFloat(result.(float64),'f',-1,64)
		 }else {
			 output="error"
		  }
		}else{
			output="error"
		}
		historyStr+=output
		historyArr=append(historyArr, historyStr)
		historyStr=""
				input.SetText(output)
		// result is now set to "true", the bool value.

	})
	dotBtn:=widget.NewButton(".", func() {
		output+="."
		input.SetText(output)
	})
	oneBtn:=widget.NewButton("1",func(){
		output+="1"
		input.SetText(output)
	})
	twoBtn:=widget.NewButton("2",func(){
		output+="2"
		input.SetText(output)
	})
	threeBtn:=widget.NewButton("3",func(){
		output+="3"
		input.SetText(output)
	})
	fourBtn:=widget.NewButton("4",func(){
		output+="4"
		input.SetText(output)
	})
	fiveBtn:=widget.NewButton("5",func(){
		output+="5"
		input.SetText(output)
	})
	sixBtn:=widget.NewButton("6",func(){
		output+="6"
		input.SetText(output)
	})
	sevenBtn:=widget.NewButton("7",func(){
		output+="7"
		input.SetText(output)
	})
	eightBtn:=widget.NewButton("8",func(){
		output+="8"
		input.SetText(output)
	})
	nineBtn:=widget.NewButton("9",func(){
		output+="9"
		input.SetText(output)
	})
	zeroBtn:=widget.NewButton("0",func(){
		output+="0"
		input.SetText(output)
	})
	calcUI:=container.NewVBox(input,history,
		container.NewGridWithColumns(2,historyBtn,backBtn),
		container.NewGridWithColumns(4,clearBtn,openBtn,closingBtn,divideBtn),
		container.NewGridWithColumns(4,nineBtn,eightBtn,sevenBtn,multiplyBtn),
		container.NewGridWithColumns(4,sixBtn,fiveBtn,fourBtn,minusBtn),
		container.NewGridWithColumns(4,threeBtn,twoBtn,oneBtn,plusBtn),
		container.NewGridWithColumns(2,container.NewGridWithColumns(2,zeroBtn,dotBtn),equalsBtn))
	w.SetContent(container.NewBorder(deskBtn,nil,nil,nil,calcUI))
	w.Show()
}
