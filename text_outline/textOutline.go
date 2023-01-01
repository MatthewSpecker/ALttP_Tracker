package text_outline

/*To Do:
-Add descriptions to functions
-Add default parameters to NewTextOutline
-Make NewTextOutline parameters easier to understand (primarily size, thick, and dense)
*/

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type TextOutline struct {
	text                 string
	textObject           []*canvas.Text
	outlineObject        []*canvas.Text
	TextOutlineContainer *fyne.Container
	textColor            color.Color
	outlineColor         color.Color
	textSize             float32
	thickness            float32
	density              float32
	coordinateCount      int
}

func NewTextOutline(textInput string, size float32, thick float32, dense int, colorFore color.Color, colorBack color.Color) *TextOutline {
	outlined := &TextOutline{
		text:                 textInput,
		TextOutlineContainer: container.New(&outlineLayout{thick}),
		textColor:            colorFore,
		outlineColor:         colorBack,
		textSize:             size * theme.Padding() / 5.0,
		thickness:            thick,
		density:              thick / (float32(dense)),
		coordinateCount:      (dense * 2) + 1,
	}

	combineContainer := []*fyne.Container{}
	var xCoordinate, yCoordinate float32
	var newPos fyne.Position

	for textIndex, element := range outlined.text {
		combineContainer = append(combineContainer, container.NewWithoutLayout())
		xCoordinate = -outlined.thickness
		for xIndex := 0; xIndex < outlined.coordinateCount; xIndex++ {
			yCoordinate = -outlined.thickness
			for yIndex := 0; yIndex < outlined.coordinateCount; yIndex++ {
				outlined.outlineObject = append(outlined.outlineObject, canvas.NewText(string(element), outlined.outlineColor))
				outlined.outlineObject[len(outlined.outlineObject)-1].TextSize = outlined.textSize
				newPos = outlined.outlineObject[len(outlined.outlineObject)-1].Position()
				newPos = newPos.AddXY(xCoordinate, yCoordinate)
				outlined.outlineObject[len(outlined.outlineObject)-1].Move(newPos)
				combineContainer[textIndex].Add(outlined.outlineObject[len(outlined.outlineObject)-1])
				yCoordinate += outlined.density
			}
			xCoordinate += outlined.density
		}
		outlined.textObject = append(outlined.textObject, canvas.NewText(string(element), outlined.textColor))
		outlined.textObject[textIndex].TextSize = outlined.textSize
		combineContainer[textIndex].Add(outlined.textObject[textIndex])
		outlined.TextOutlineContainer.Add(combineContainer[textIndex])
	}

	return outlined
}

func (t *TextOutline) Refresh(textInput string) {
	t.text = textInput
	combineContainer := []*fyne.Container{}
	var xCoordinate, yCoordinate float32
	var newPos fyne.Position
	t.textObject = t.textObject[:0]
	t.outlineObject = t.outlineObject[:0]
	t.TextOutlineContainer.RemoveAll()

	for textIndex, element := range t.text {
		combineContainer = append(combineContainer, container.NewWithoutLayout())
		xCoordinate = -t.thickness
		for xIndex := 0; xIndex < t.coordinateCount; xIndex++ {
			yCoordinate = -t.thickness
			for yIndex := 0; yIndex < t.coordinateCount; yIndex++ {
				t.outlineObject = append(t.outlineObject, canvas.NewText(string(element), t.outlineColor))
				t.outlineObject[len(t.outlineObject)-1].TextSize = t.textSize
				newPos = t.outlineObject[len(t.outlineObject)-1].Position()
				newPos = newPos.AddXY(xCoordinate, yCoordinate)
				t.outlineObject[len(t.outlineObject)-1].Move(newPos)
				combineContainer[textIndex].Add(t.outlineObject[len(t.outlineObject)-1])
				yCoordinate += t.density
			}
			xCoordinate += t.density
		}
		t.textObject = append(t.textObject, canvas.NewText(string(element), t.textColor))
		t.textObject[textIndex].TextSize = t.textSize
		combineContainer[textIndex].Add(t.textObject[textIndex])
		t.TextOutlineContainer.Add(combineContainer[textIndex])
	}

	t.TextOutlineContainer.Refresh()
}
