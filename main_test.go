package main

import (
	"fyne.io/fyne/v2/test"
	"testing"
)

func TestAdd(t *testing.T) {
	calc := newCalculator()
	calc.loadUI(test.NewApp())

	test.Tap(calc.buttons["1"])
	test.Tap(calc.buttons["+"])
	test.Tap(calc.buttons["1"])
	test.Tap(calc.buttons["="])

	assert.Equal(t, "2", calc.output.Text)
}
