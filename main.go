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
			rectangle()
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

	label := widget.NewLabel("Calculate area of a circle")

	radiusInput := widget.NewEntry()
	radiusInput.SetPlaceHolder("Enter the radius...")

	areaLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Calculate", func() {

		// radius
		radiusText := radiusInput.Text
		r, err := strconv.ParseFloat(radiusText, 64)
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

	window.SetContent(container.NewVBox(label, radiusInput, buttonGroup, areaLabel))
}

func triangle() {
    var area float64

	label := widget.NewLabel("Calculate area of a triangle")

	perpendicularHeightInput := widget.NewEntry()
	perpendicularHeightInput.SetPlaceHolder("Enter the perpendicular height...")

	baseLengthInput := widget.NewEntry()
	baseLengthInput.SetPlaceHolder("Enter the base length...")

	areaLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Calculate", func() {

		// perpendicular height
		perpendicularHeightText := perpendicularHeightInput.Text
		perpendicularHeight, perpendicularHeightErr := strconv.ParseFloat(perpendicularHeightText, 64)
		if perpendicularHeightErr != nil {
			log.Println("Error converting text to float64 of `perpendicular height`: ", perpendicularHeightErr)
			areaLabel.SetText("Error converting text to float64")
			return
		}


		// base length
		baseLengthText := baseLengthInput.Text
		baseLength, baseLengthErr := strconv.ParseFloat(baseLengthText, 64)
		if baseLengthErr != nil {
			log.Println("Error converting text to float64 of `base length`: ", baseLengthErr)
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

	window.SetContent(container.NewVBox(label, perpendicularHeightInput, baseLengthInput, buttonGroup, areaLabel))
}

func rectangle(){
	var area float64

	label := widget.NewLabel("Calculate area of a rectangle")

	lengthInput := widget.NewEntry()
	lengthInput.SetPlaceHolder("Enter the length...")

	widthInput := widget.NewEntry()
	widthInput.SetPlaceHolder("Enter the width...")

	areaLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Calculate", func() {

		// length
		lengthText := lengthInput.Text
		length, lengthErr := strconv.ParseFloat(lengthText, 64)
		if lengthErr != nil {
			log.Println("Error converting text to float64 of `length`: ", lengthErr)
			areaLabel.SetText("Error converting text to float64")
			return
		}


		// width
		widthText := widthInput.Text
		width, widthErr := strconv.ParseFloat(widthText, 64)
		if widthErr != nil {
			log.Println("Error converting text to float64 of `width`: ", widthErr)
			areaLabel.SetText("Error converting text to float64")
			return
		}

  		area = length * width

		areaStr := strconv.FormatFloat(area, 'f', 2, 64)
		areaLabel.SetText("Area is: " + areaStr)
	})
	
	backButton := widget.NewButton("Back", func() {
		mainContainer()
	})

	buttonGroup := container.New(layout.NewHBoxLayout(), calculateButton, backButton)

	window.SetContent(container.NewVBox(label, lengthInput, widthInput, buttonGroup, areaLabel))
}

func main() {
	myApp := app.New()
	window = myApp.NewWindow("Area Calculator")
	window.Resize(fyne.Size{Width: 400, Height: 300})

	mainContainer()

	window.ShowAndRun()
}
