package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

// Calculator represents a simple calculator with display and expression logic.
type Calculator struct {
	display *widget.Label // Display widget to show current expression or result
	expr    string        // Current expression being entered
}

// NewCalculator creates a new instance of Calculator with the given display widget.
func NewCalculator(display *widget.Label) *Calculator {
	return &Calculator{display: display}
}

// Press simulates pressing a button on the calculator.
func (c *Calculator) Press(key string) {
	switch key {
	case "C":
		c.expr = "" // Clear the expression
	case "=":
		c.evaluate() // Evaluate the expression
	default:
		c.expr += key // Append the key to the expression
	}
	c.updateDisplay() // Update the display with the current expression
}

// evaluate evaluates the current expression and updates the expression accordingly.
func (c *Calculator) evaluate() {
	expression, err := govaluate.NewEvaluableExpression(c.expr)
	if err != nil {
		c.expr = "Error" // Set expression to "Error" if evaluation fails
	} else {
		result, err := expression.Evaluate(nil)
		if err != nil {
			c.expr = "Error" // Set expression to "Error" if evaluation fails
		} else {
			c.expr = fmt.Sprintf("%v", result) // Convert the result to string and set it as the expression
		}
	}
}

// updateDisplay updates the display widget with the current expression.
func (c *Calculator) updateDisplay() {
	c.display.SetText(c.expr)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	display := widget.NewLabel("") // Create a label widget for displaying the expression/result
	display.SetText("")            // Set initial text to empty string

	calculator := NewCalculator(display) // Create a new calculator instance

	// Define the labels for buttons on the calculator
	buttons := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "+", "-", "*", "/", "C", "="}

	// Function to create buttons with given label and callback
	createButton := func(label string) *widget.Button {
		return widget.NewButton(label, func() {
			calculator.Press(label) // When the button is clicked, simulate pressing it on the calculator
		})
	}

	// Create buttons for each label
	var calcButtons []*widget.Button
	for _, label := range buttons {
		calcButtons = append(calcButtons, createButton(label))
	}

	// Set the content of the window with calculator display and buttons
	myWindow.SetContent(container.NewVBox(
		display,
		container.NewGridWithColumns(3,
			calcButtons[13], calcButtons[12], calcButtons[11], // Divide, Multiply, Subtract buttons
		),
		container.NewGridWithRows(3,
			container.NewGridWithColumns(3,
				calcButtons[6], calcButtons[7], calcButtons[8], // 7, 8, 9 buttons
			),
			container.NewGridWithColumns(3,
				calcButtons[3], calcButtons[4], calcButtons[5], // 4, 5, 6 buttons
			),
			container.NewGridWithColumns(3,
				calcButtons[0], calcButtons[1], calcButtons[2], // 1, 2, 3 buttons
			),
		),
		container.NewGridWithColumns(3,
			calcButtons[14], calcButtons[10], calcButtons[15], // Clear, Equals and Addition buttons
		),
		widget.NewLabelWithStyle("Made By Shalom", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), // Text box to display name
	))

	myWindow.ShowAndRun() // Show and run the application window
}
