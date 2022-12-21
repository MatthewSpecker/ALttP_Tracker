package tappable_icons

import (
	"errors"
	"image/color"

	"tracker/save"
	"tracker/tooltip"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TappableIconCycled struct {
	widget.Icon
	desktop.Hoverable
	resources      []fyne.Resource
	current        int
	tapSize        float32
	toolTipText    string
	toolTipPopUp   *widget.PopUp
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableIconCycled(res []fyne.Resource, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableIconCycled, error) {
	if len(res) <= 1 {
		return nil, errors.New("'res' must contain 2 or more resources")
	}
	if size <= 0 {
		return nil, errors.New("'size' must be float32 greater than 0")
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	icon := &TappableIconCycled{
		resources:      res,
		current:        0,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	icon.toolTipText = tooltip.GetToolTipText(icon.resources[icon.current].Name())
	icon.ExtendBaseWidget(icon)
	icon.SetResource(icon.resources[icon.current])

	return icon, nil
}

func (t *TappableIconCycled) Update() {
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, len(t.resources)-1, 0)
	t.toolTipText = tooltip.GetToolTipText(t.resources[t.current].Name())

	t.Icon.SetResource(t.resources[t.current])
}

func (t *TappableIconCycled) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Current", 0)
}

func (t *TappableIconCycled) GetSaveDefaults() {
	t.current = 0
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.Update()
}

func (t *TappableIconCycled) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableIconCycled) increment() {
	if t.current < (len(t.resources) - 1) {
		t.current++
	} else {
		t.current = 0
	}
	t.Icon.SetResource(t.resources[t.current])
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableIconCycled) decrement() {
	if t.current > 0 {
		t.current--
	} else {
		t.current = len(t.resources) - 1
	}
	t.Icon.SetResource(t.resources[t.current])
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableIconCycled) Tapped(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	t.increment()
}

func (t *TappableIconCycled) TappedSecondary(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
	t.decrement()
}

func (t *TappableIconCycled) Keyed() {
	t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	t.increment()
}

func (t *TappableIconCycled) MouseIn(event *desktop.MouseEvent) {
	//t.toolTipPopUp = newToolTipTextTappableIconCycled(event, t.toolTipText, t)
}

func (t *TappableIconCycled) MouseMoved(_ *desktop.MouseEvent) {
}

func (t *TappableIconCycled) MouseOut() {
	//t.toolTipPopUp.Hide()
}

func newToolTipTextTappableIconCycled(event *desktop.MouseEvent, text string, object *TappableIconCycled) *widget.PopUp {
	toolTipText := canvas.NewText(text, color.White)
	popUp := widget.NewPopUp(toolTipText, fyne.CurrentApp().Driver().CanvasForObject(object))
	var popUpPosition fyne.Position
	popUpPosition.X = event.AbsolutePosition.X + object.Size().Width/2
	popUpPosition.Y = event.AbsolutePosition.Y - object.Size().Height/2
	popUp.ShowAtPosition(popUpPosition)

	return popUp
}
