package tappable_icons

import (
	"errors"
	"image/color"
	"strconv"

	"tracker/save"
	"tracker/tooltip"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TappableIconWithNum struct {
	widget.Icon
	desktop.Hoverable
	resources      []fyne.Resource
	current        int
	number         int
	numberMax      int
	numberLabel    *widget.Label
	tapSize        float32
	toolTipText    string
	toolTipPopUp   *widget.PopUp
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableIconWithNum(res []fyne.Resource, num int, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableIconWithNum, error) {
	if len(res) <= 1 {
		return nil, errors.New("'res' must contain 2 or more resources")
	}
	if num < 0 {
		return nil, errors.New("'num' must be a non-negative integer")
	}
	if size <= 0 {
		return nil, errors.New("'size' must be float32 greater than 0")
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	icon := &TappableIconWithNum{
		resources:      res,
		current:        0,
		number:         0,
		numberMax:      num,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	icon.numberLabel = widget.NewLabel("")
	icon.toolTipText = tooltip.GetToolTipText(icon.resources[icon.current].Name())
	icon.ExtendBaseWidget(icon)
	icon.SetResource(icon.resources[icon.current])

	return icon, nil
}

func (t *TappableIconWithNum) Update() {
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, len(t.resources)-1, 0)
	t.number = t.saveFile.GetSaveInt(t.saveFileText + "_Number")
	t.number = intRangeCheck(t.number, t.numberMax, 0)
	t.toolTipText = tooltip.GetToolTipText(t.resources[t.current].Name())

	if t.number == 0 {
		t.numberLabel.SetText("")
	} else {
		t.numberLabel.SetText(strconv.Itoa(t.number))
	}

	t.Icon.SetResource(t.resources[t.current])
}

func (t *TappableIconWithNum) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Current", 0)
	t.saveFile.SetDefault(t.saveFileText+"_Number", 0)
}

func (t *TappableIconWithNum) GetSaveDefaults() {
	t.current = 0
	t.number = 0
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
	t.Update()
}

func (t *TappableIconWithNum) LayoutAdjust() *fyne.Container {
	container1 := container.New(layout.NewCenterLayout(), t)
	container2 := container.NewWithoutLayout(container1, t.numberLabel)
	iconSize := t.Size()
	iconChangePosition := fyne.NewPos(iconSize.Width/2, iconSize.Height/2)
	t.numberLabel.Move(iconChangePosition)

	return container2
}

func (t *TappableIconWithNum) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableIconWithNum) increment() {
	if t.current < (len(t.resources) - 1) {
		t.current++
		t.toolTipText = tooltip.GetToolTipText(t.resources[t.current].Name())
		t.Icon.SetResource(t.resources[t.current])
	}
	if t.number < t.numberMax {
		t.number++
		t.numberLabel.SetText(strconv.Itoa(t.number))
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
}

func (t *TappableIconWithNum) decrement() {
	if t.current > 0 && t.number <= t.current {
		t.current--
		t.toolTipText = tooltip.GetToolTipText(t.resources[t.current].Name())
		t.Icon.SetResource(t.resources[t.current])
	}
	if t.number == 1 {
		t.number--
		t.numberLabel.SetText("")
	} else if t.number > 1 {
		t.number--
		t.numberLabel.SetText(strconv.Itoa(t.number))
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
}

func (t *TappableIconWithNum) Tapped(_ *fyne.PointEvent) {
	if t.current < (len(t.resources) - 1) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}

func (t *TappableIconWithNum) TappedSecondary(_ *fyne.PointEvent) {
	if t.current > 0 {
		t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
	}
	t.decrement()
}

func (t *TappableIconWithNum) Keyed() {
	if t.current < (len(t.resources) - 1) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}

func (t *TappableIconWithNum) MouseIn(event *desktop.MouseEvent) {
	//t.toolTipPopUp = newToolTipTextTappableIconWithNum(event, t.toolTipText, t)
}

func (t *TappableIconWithNum) MouseMoved(_ *desktop.MouseEvent) {
}

func (t *TappableIconWithNum) MouseOut() {
	//t.toolTipPopUp.Hide()
}

func newToolTipTextTappableIconWithNum(event *desktop.MouseEvent, text string, object *TappableIconWithNum) *widget.PopUp {
	toolTipText := canvas.NewText(text, color.White)
	popUp := widget.NewPopUp(toolTipText, fyne.CurrentApp().Driver().CanvasForObject(object))
	var popUpPosition fyne.Position
	popUpPosition.X = event.AbsolutePosition.X + object.Size().Width/2
	popUpPosition.Y = event.AbsolutePosition.Y - object.Size().Height/2
	popUp.ShowAtPosition(popUpPosition)

	return popUp
}
