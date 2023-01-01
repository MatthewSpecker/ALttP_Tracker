package tappable_icons

import (
	"errors"
	"image/color"

	"tracker/save"
	"tracker/text_outline"
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

type TappableIconWithString struct {
	widget.Icon
	desktop.Hoverable
	resources      []fyne.Resource
	obtained       bool
	textPath       []string
	textCurrent    int
	outlineText    *text_outline.TextOutline
	tapSize        float32
	toolTipText    string
	toolTipPopUp   *widget.PopUp
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableIconWithString(res []fyne.Resource, text []string, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableIconWithString, error) {
	if len(res) <= 1 {
		return nil, errors.New("'res' must contain 2 or more resources")
	}
	if len(text) == 0 {
		return nil, errors.New("'text' cannot be empty slice")
	}
	if size <= 0 {
		return nil, errors.New("'size' must be float32 greater than 0")
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	icon := &TappableIconWithString{
		resources:      res,
		obtained:       false,
		textCurrent:    0,
		textPath:       text,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	icon.outlineText = text_outline.NewTextOutline(icon.textPath[icon.textCurrent], icon.tapSize*theme.Padding()/6, 1.5, 6, color.White, color.Black)

	icon.ExtendBaseWidget(icon)
	if icon.obtained {
		icon.SetResource(icon.resources[1])
		icon.toolTipText = tooltip.GetToolTipText(icon.resources[1].Name())
	} else {
		icon.SetResource(icon.resources[0])
		icon.toolTipText = tooltip.GetToolTipText(icon.resources[0].Name())
	}

	return icon, nil
}

func (t *TappableIconWithString) Update() {
	t.obtained = t.saveFile.GetSaveBool(t.saveFileText + "_Obtained")
	t.textCurrent = t.saveFile.GetSaveInt(t.saveFileText + "_TextCurrent")
	t.textCurrent = intRangeCheck(t.textCurrent, len(t.resources)-1, 0)

	if t.obtained {
		t.Icon.SetResource(t.resources[1])
		t.toolTipText = tooltip.GetToolTipText(t.resources[1].Name())
	} else {
		t.Icon.SetResource(t.resources[0])
		t.toolTipText = tooltip.GetToolTipText(t.resources[0].Name())
	}

	t.outlineText.Refresh(t.textPath[t.textCurrent])
}

func (t *TappableIconWithString) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Obtained", false)
	t.saveFile.SetDefault(t.saveFileText+"_TextCurrent", 0)
}

func (t *TappableIconWithString) GetSaveDefaults() {
	t.obtained = false
	t.textCurrent = 0
	t.saveFile.SetSave(t.saveFileText+"_Obtained", t.obtained)
	t.saveFile.SetSave(t.saveFileText+"_TextCurrent", t.textCurrent)
	t.Update()
}

func (t *TappableIconWithString) Layout() *fyne.Container {
	tapIconContainer := t.layoutString()
	tapIconContainer = container.New(layout.NewCenterLayout(), tapIconContainer)
	return tapIconContainer
}

func (t *TappableIconWithString) layoutString() *fyne.Container {
	container1 := container.New(layout.NewCenterLayout(), t)
	container2 := container.New(layout.NewCenterLayout(), t.outlineText.TextOutlineContainer)
	container3 := container.NewWithoutLayout(container1, container2)
	iconSize := t.Size()
	container2Size := t.outlineText.TextOutlineContainer.Size()
	iconChangePosition := container2.Position()
	iconChangePosition = iconChangePosition.AddXY(iconSize.Width/2-container2Size.Width/2, iconSize.Height/8-container2Size.Height/2)
	container2.Move(iconChangePosition)

	return container3
}

func (t *TappableIconWithString) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableIconWithString) setIcon() {
	if t.obtained == true {
		t.obtained = false
		t.Icon.SetResource(t.resources[0])
	} else {
		t.obtained = true
		t.Icon.SetResource(t.resources[1])
	}
	t.saveFile.SetSave(t.saveFileText+"_Obtained", t.obtained)
}

func (t *TappableIconWithString) setTextIncrement() {
	if t.textCurrent == len(t.textPath)-1 {
		t.textCurrent = 0
	} else if t.textCurrent < len(t.textPath)-1 {
		t.textCurrent++
	}
	t.outlineText.Refresh(t.textPath[t.textCurrent])
	t.saveFile.SetSave(t.saveFileText+"_TextCurrent", t.textCurrent)
}

func (t *TappableIconWithString) setTextDecrement() {
	if t.textCurrent == 0 {
		t.textCurrent = len(t.textPath) - 1
	} else if t.textCurrent <= len(t.textPath)-1 {
		t.textCurrent--
	}
	t.outlineText.Refresh(t.textPath[t.textCurrent])
	t.saveFile.SetSave(t.saveFileText+"_TextCurrent", t.textCurrent)
}

func (t *TappableIconWithString) Tapped(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.setIcon, t.setIcon)
	t.setIcon()
}

func (t *TappableIconWithString) TappedSecondary(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.setTextDecrement, t.setTextIncrement)
	t.setTextIncrement()
}

func (t *TappableIconWithString) Keyed() {
	t.undoRedoStacks.StoreFunctions(t.setIcon, t.setIcon)
	t.setIcon()
}

func (t *TappableIconWithString) MouseIn(event *desktop.MouseEvent) {
	//t.toolTipPopUp = newToolTipTextTappableIconWithString(event, t.toolTipText, t)
}

func (t *TappableIconWithString) MouseMoved(_ *desktop.MouseEvent) {
}

func (t *TappableIconWithString) MouseOut() {
	//t.toolTipPopUp.Hide()
}

func newToolTipTextTappableIconWithString(event *desktop.MouseEvent, text string, object *TappableIconWithString) *widget.PopUp {
	toolTipText := canvas.NewText(text, color.White)
	popUp := widget.NewPopUp(toolTipText, fyne.CurrentApp().Driver().CanvasForObject(object))
	var popUpPosition fyne.Position
	popUpPosition.X = event.AbsolutePosition.X + object.Size().Width/2
	popUpPosition.Y = event.AbsolutePosition.Y - object.Size().Height/2
	popUp.ShowAtPosition(popUpPosition)

	return popUp
}
