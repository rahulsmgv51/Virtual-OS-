package main

import (
	"fmt"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)
var f *os.File
var format beep.Format
var streamer beep.StreamSeekCloser
var pause bool = false

func showMediaApp(w fyne.Window){

	go func (msg string)  {
		fmt.Println(msg)
		if streamer == nil {
			
		}else{
			fmt.Println(fmt.Sprint(streamer.Len()))
		}
	}("going")

	time.Sleep(time.Second)

	// a := app.New()
	// w := a.NewWindow("Media Player")
	w.Resize(fyne.NewSize(333, 333))

	logo := canvas.NewImageFromFile("C:\\Users\\DELL\\Desktop\\LRC\\LRC_SET_OS\\logo.jpg")
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.MediaPlayIcon(), func ()  {
			speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
			speaker.Play(streamer)
		}),
		widget.NewToolbarAction(theme.MediaPauseIcon(), func() {
			if !pause{
				pause = true
				speaker.Lock()
			}else if pause{
				pause =false
				speaker.Unlock()
			}
		}),
		widget.NewToolbarAction(theme.MediaStopIcon(), func(){
			speaker.Close()
		}),
		widget.NewToolbarSpacer(),
	)

	label := widget.NewLabel("Media Player..")
	label.Alignment = fyne.TextAlignCenter
	label2 := widget.NewLabel("play mp3..")
	label2.Alignment = fyne.TextAlignCenter

	browseFiles := widget.NewButton("chosse our music", func() {
		fileOpen := dialog.NewFileOpen(func (uc fyne.URIReadCloser, _ error)  {
			streamer, format, _ = mp3.Decode(uc)
			label2.Text = uc.URI().Name()
			label2.Refresh()
		}, w)
		fileOpen.Show()
		fileOpen.SetFilter(storage.NewExtensionFileFilter([]string{".mp3"}))
	})

	panelContent := container.NewVBox(label, browseFiles, label2, toolbar)
	w.SetContent(
		container.NewBorder(logo, nil, nil, nil, panelContent),
	)

	w.Show()
}