package main

import (
	"encoding/json"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io"
	"net/http"
	"strconv"
)

func showWeatherApp(w fyne.Window)  {
	//a:=app.New()
	//a.Settings().SetTheme(theme.DarkTheme())
	//w:=a.NewWindow("Weather App")
	//w.Resize(fyne.NewSize(300,300))
		//res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=88a3325d8b543b9103c71abe0ebc15ef",""))
		//if err != nil {
		//	fmt.Println(err)
		//}
		//body, err := io.ReadAll(res.Body)
		//defer res.Body.Close()
		//weather, err := UnmarshalWeather(body)
		//if err != nil {
		//	fmt.Println(err)
		//}

		img := canvas.NewImageFromFile("C:\\Users\\user\\Desktop\\test case\\dice\\download.jpg")
		img.FillMode = canvas.ImageFillOriginal
		label1 := canvas.NewText("Weather details", color.Black)
		label1.TextStyle = fyne.TextStyle{Italic: true}
		label2 := canvas.NewText(fmt.Sprintf("Country %s", "IN"), color.White)
		label3 := canvas.NewText(fmt.Sprintf("Wind Speed   %s", "--"), color.White)
		label4 := canvas.NewText(fmt.Sprintf("Humidity   %s", "--"), color.White)
		label5 := canvas.NewText(fmt.Sprintf("Temp   %s", "--"), color.White)
	combo := widget.NewSelect([]string{"Jaipur", "Mumbai","Delhi"}, func(value string) {
		res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&APPID=88a3325d8b543b9103c71abe0ebc15ef", value))
		if err != nil {
			fmt.Println(err)
		}
		body, err := io.ReadAll(res.Body)
		defer res.Body.Close()
		weather, err := UnmarshalWeather(body)
		if err != nil {
			fmt.Println(err)
		}
		label3.Text="Wind Speed   "+covertToString(weather.Wind.Speed)
		label3.Refresh()
		s:=strconv.FormatInt(weather.Main.Humidity, 10)
		label4.Text="Humidity   "+s
		label4.Refresh()
		label5.Text="Temp   "+covertToString(weather.Main.Temp)
		label5.Refresh()
	})
		vBox:=container.NewVBox(label1, img, combo, label2, label3, label4, label5)
		w.SetContent(container.NewBorder(panelContent,nil,nil,nil,vBox))
		w.Show()
	}

func covertToString(val float64,) string {
	res:=strconv.FormatFloat(val,'f',2,64)
	return res
}

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weather, err := UnmarshalWeather(bytes)
//    bytes, err = weather.Marshal()

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Coord      Coord            `json:"coord"`
	Weather    []WeatherElement `json:"weather"`
	Base       string           `json:"base"`
	Main       Main             `json:"main"`
	Visibility int64            `json:"visibility"`
	Wind       Wind             `json:"wind"`
	Clouds     Clouds           `json:"clouds"`
	Dt         int64            `json:"dt"`
	Sys        Sys              `json:"sys"`
	Timezone   int64            `json:"timezone"`
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	Cod        int64            `json:"cod"`
}

type Clouds struct {
	All int64 `json:"all"`
}

type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int64   `json:"pressure"`
	Humidity  int64   `json:"humidity"`
}

type Sys struct {
	Type    int64  `json:"type"`
	ID      int64  `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

type WeatherElement struct {
	ID          int64  `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int64   `json:"deg"`
}
