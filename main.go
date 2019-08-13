package main

import (
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func main() {
	os.Setenv("FYNE_SCALE", "2.0")
	app := app.New()

	pathEntry := &widget.Entry{PlaceHolder: "Enter folder path here."}

	gallery := widget.NewScrollContainer(fyne.NewContainerWithLayout(layout.NewGridLayout(10),
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
		&canvas.Image{FillMode: canvas.ImageFillOriginal, File: "pic.jpg"},
	))
	gallery.Resize(fyne.Size{Width: 2000, Height: 2000})

	navbar := widget.NewHBox(
		widget.NewButton("Prev", func() {
			app.Quit()
		}),
		layout.NewSpacer(),
		widget.NewButton("Crop", func() {
			app.Quit()
		}),
		widget.NewButton("Del", func() {
			app.Quit()
		}),
		layout.NewSpacer(),
		widget.NewButton("Next", func() {
			app.Quit()
		}),
	)
	fmt.Print(navbar.MinSize())

	w := app.NewWindow("Golabel")
	w.SetContent(fyne.NewContainerWithLayout(
		layout.NewBorderLayout(pathEntry, navbar, nil, nil),
		pathEntry,
		gallery,
		navbar,
	))
	w.Resize(fyne.Size{Width: 2000, Height: 2000})

	w.ShowAndRun()
}
