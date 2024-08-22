package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("refreshing fyne knowledge")
	label := widget.NewLabel("Hello world!!")
	helloButton := widget.NewButton("Hello!", func() {
		window := a.NewWindow("Button click")
		window.SetContent(widget.NewLabel("Hello Amogh"))
		window.Show()
	})
	contents := container.NewVBox(label, helloButton)
	w.SetContent(contents)
	w.ShowAndRun()

}
