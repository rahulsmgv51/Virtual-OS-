package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)
var num_article int = 1
var news News
func showNewsApp(w fyne.Window) {

    // a := app.New()
    // w := a.NewWindow("News App")
    // w.Resize(fyne.NewSize(333, 444))

    res, _ := http.Get("https://gnews.io/api/v4/top-headlines?token=772105fd51198d5ab951b6e97a5de844&lang=en&max=100")
    defer res.Body.Close()

    body, _ := ioutil.ReadAll(res.Body)
    news, _ = UnmarshalNews(body)
    //fmt.Println(news)
    
    label1 := widget.NewLabel(fmt.Sprintf("No of articles:%d",news.TotalArticles))
    
    label2 := widget.NewLabel(fmt.Sprintf("%s", news.Articles[1].Title))
    label2.TextStyle = fyne.TextStyle{Bold: true}
    label2.Wrapping = fyne.TextWrapBreak

    label3 := widget.NewLabel(fmt.Sprintf("%s", news.Articles[1].Description))
    label3.Wrapping = fyne.TextWrapBreak

    btn := widget.NewButton("Next", func() {
        num_article += 1
        label2.Text = news.Articles[num_article].Title
        label3.Text = news.Articles[num_article].Description
        label2.Refresh()
        label3.Refresh()
    })
    label4 := canvas.NewText("News App", color.Black)//Heading
    label4.Alignment = fyne.TextAlignCenter
    label4.TextStyle = fyne.TextStyle{Bold: true}

    img := canvas.NewImageFromFile("C:\\Users\\DELL\\Desktop\\LRC\\LRC_SET_OS\\newsimg.jpg")
    img.FillMode = canvas.ImageFillOriginal

    conte := container.NewVBox(label4, label2, label3, btn)
    conte.Resize(fyne.NewSize(300, 400))

    panelContent := container.NewBorder(img, label1, nil, nil, conte)

    w.SetContent(panelContent)

    w.Show()
}
// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    news, err := UnmarshalNews(bytes)
//    bytes, err = news.Marshal()
func UnmarshalNews(data []byte) (News, error) {
    var r News
    err := json.Unmarshal(data, &r)
    return r, err
}
func (r *News) Marshal() ([]byte, error) {
    return json.Marshal(r)
}
type News struct {
    TotalArticles int64     `json:"totalArticles"`
    Articles      []Article `json:"articles"`
}
type Article struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    Content     string `json:"content"`
    URL         string `json:"url"`
    Image       string `json:"image"`
    PublishedAt string `json:"publishedAt"`
    Source      Source `json:"source"`
}
type Source struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}