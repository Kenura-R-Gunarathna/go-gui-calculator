package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var window fyne.Window

func mainContainer() {

	label := widget.NewLabel("Select the area calculation")

	options := []string{"Circle", "Triangle", "Rectangle", "Parallelogram", "Trapezium"}

	combo := widget.NewSelect(options, func(value string) {
			
		switch value {
		case "Circle":
			circle()
		case "Triangle":
			triangle()
		case "Rectangle":
			notImplemented()
		case "Parallelogram":
			notImplemented()
		case "Trapezium":
			notImplemented()
		default:
			notImplemented()
		}

	})

	window.SetContent(container.NewVBox(label, combo))
}

func notImplemented() {

	label := widget.NewLabel("Not implemented yet")

	backButton := widget.NewButton("Back", func() {
		mainContainer()
	})

	window.SetContent(container.NewVBox(label, backButton))
}

func circle() {
	const PI float64 = 3.14
	var area float64

	radiusInput := widget.NewEntry()
	radiusInput.SetPlaceHolder("Enter the radius...")
	radiusText := radiusInput.Text
	r, err := strconv.ParseFloat(radiusText, 64)

	areaLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Calculate", func() {
		if err != nil {
			log.Println("Error converting text to float64 of `radius`: ", err)
			areaLabel.SetText("Error converting text to float64")
			return
		}

		area = PI * r * r
		
		areaStr := strconv.FormatFloat(area, 'f', 2, 64)
		areaLabel.SetText("Area is: " + areaStr)
	})

	
	backButton := widget.NewButton("Back", func() {
		mainContainer()
	})

	buttonGroup := container.New(layout.NewHBoxLayout(), calculateButton, backButton)

	window.SetContent(container.NewVBox(radiusInput, buttonGroup, areaLabel))
}

func triangle() {
    var area float64

	perpendicularHeightInput := widget.NewEntry()
	perpendicularHeightInput.SetPlaceHolder("Enter the perpendicular height...")
	perpendicularHeightText := perpendicularHeightInput.Text
	perpendicularHeight, perpendicularHeightErr := strconv.ParseFloat(perpendicularHeightText, 64)

	baseLengthInput := widget.NewEntry()
	baseLengthInput.SetPlaceHolder("Enter the base length...")
	baseLengthText := baseLengthInput.Text
	baseLength, baseLengthErr := strconv.ParseFloat(baseLengthText, 64)

	areaLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Calculate", func() {

		if perpendicularHeightErr != nil {
			log.Println("Error converting text to float64 of `perpendicular height`: ", perpendicularHeightErr)
			areaLabel.SetText("Error converting text to float64")
			return
		}

		if baseLengthErr != nil {
			log.Println("Error converting text to float64 of `base length.`: ", baseLengthErr)
			areaLabel.SetText("Error converting text to float64")
			return
		}

    	area = perpendicularHeight * baseLength / 2

		areaStr := strconv.FormatFloat(area, 'f', 2, 64)
		areaLabel.SetText("Area is: " + areaStr)
	})

	
	backButton := widget.NewButton("Back", func() {
		mainContainer()
	})

	buttonGroup := container.New(layout.NewHBoxLayout(), calculateButton, backButton)

	window.SetContent(container.NewVBox(perpendicularHeightInput, baseLengthInput, buttonGroup, areaLabel))
}

func main() {
	myApp := app.New()
	window = myApp.NewWindow("Area Calculator")
	window.Resize(fyne.Size{Width: 400, Height: 300})

	mainContainer()

	window.ShowAndRun()
}
