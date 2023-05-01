package main

import (
	"fmt"
	"io"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Marsel App")
	w.Resize(fyne.NewSize(800, 500))

	entry := widget.NewMultiLineEntry()
	entry.Resize(fyne.NewSize(600, 300))
	entry.Move(fyne.NewPos(100, 135))

	btn := widget.NewButton("Open file", func() {
		dialog.ShowFileOpen(func(uc fyne.URIReadCloser, err error) {
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			data, _ := io.ReadAll(uc)
			entry.SetText(string(data))
		}, w)
	})
	btn.Resize(fyne.NewSize(150, 75))
	btn.Move(fyne.NewPos(325, 30))

	cont := container.NewWithoutLayout(entry, btn)

	w.SetContent(container.NewVBox(cont))
	w.Show()
	a.Run()
}
