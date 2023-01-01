package tappable_icons

import (
	"errors"
	"image/color"
	"strconv"

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

type TappableIconWithBottomCenteredString struct {
	widget.Icon
	desktop.Hoverable
	resources      []fyne.Resource
	current        int
	text           string
	obtained       bool
	outlineText    *text_outline.TextOutline
	tapSize        float32
	toolTipText    string
	toolTipPopUp   *widget.PopUp
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableIconWithBottomCenteredString(res []fyne.Resource, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableIconWithBottomCenteredString, error) {
	if len(res) != 2 {
		return nil, errors.New("'res' must contain exactly 2 resources")
	}
	if size <= 0 {
		return nil, errors.New("'size' must be float32 greater than 0")
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	icon := &TappableIconWithBottomCenteredString{
		resources:      res,
		current:        -1,
		text:           "?",
		obtained:       false,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	icon.toolTipText = tooltip.GetToolTipText(icon.resources[0].Name())
	icon.outlineText = text_outline.NewTextOutline(icon.text, icon.tapSize, 1.5, 6, color.White, color.Black)
	icon.ExtendBaseWidget(icon)
	icon.SetResource(icon.resources[0])

	return icon, nil
}

func (t *TappableIconWithBottomCenteredString) Update() {
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, 8, -1)
	t.toolTipText = tooltip.GetToolTipText(t.resources[0].Name())
	if t.current == -1 {
		t.text = "?"
	} else if t.current == 8 {
		t.text = "AD"
	} else {
		t.text = strconv.Itoa(t.current)
	}

	t.outlineText.Refresh(t.text)

	if t.obtained {
		t.Icon.SetResource(t.resources[1])
	} else {
		t.Icon.SetResource(t.resources[0])
	}
}

func (t *TappableIconWithBottomCenteredString) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Current", 7)
	t.saveFile.SetDefault(t.saveFileText+"_Obtained", false)
}

func (t *TappableIconWithBottomCenteredString) GetSaveDefaults() {
	t.current = 7
	t.obtained = false
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.saveFile.SetDefault(t.saveFileText+"_Obtained", t.obtained)
	t.Update()
}

func (t *TappableIconWithBottomCenteredString) Layout() *fyne.Container {
	tapIconContainer := t.layoutBottomCenteredString()
	tapIconContainer = container.New(layout.NewCenterLayout(), tapIconContainer)
	return tapIconContainer
}

func (t *TappableIconWithBottomCenteredString) layoutBottomCenteredString() *fyne.Container {
	container1 := container.New(layout.NewCenterLayout(), t)
	container2 := container.New(layout.NewCenterLayout(), t.outlineText.TextOutlineContainer)
	container3 := container.NewWithoutLayout(container1, container2)
	iconSize := t.Icon.Size()
	container2Size := t.outlineText.TextOutlineContainer.Size()
	iconChangePosition := container2.Position()
	iconChangePosition = iconChangePosition.AddXY(iconSize.Width/2-container2Size.Width/2, iconSize.Height/2-container2Size.Height/2)
	container2.Move(iconChangePosition)

	return container3
}

func (t *TappableIconWithBottomCenteredString) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableIconWithBottomCenteredString) changeResource() {
	if t.obtained {
		t.obtained = false
		t.Icon.SetResource(t.resources[0])
	} else {
		t.obtained = true
		t.Icon.SetResource(t.resources[1])
	}

	t.saveFile.SetSave(t.saveFileText+"_Obtained", t.obtained)
}

func (t *TappableIconWithBottomCenteredString) increment() {
	if t.current < 7 {
		t.current++
		t.text = strconv.Itoa(t.current)
	} else if t.current == 7 {
		t.current++
		t.text = "AD"
	} else {
		t.current = -1
		t.text = "?"
	}
	t.outlineText.Refresh(t.text)

	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableIconWithBottomCenteredString) decrement() {
	if t.current > 0 {
		t.current--
		t.text = strconv.Itoa(t.current)
	} else if t.current == 0 {
		t.current--
		t.text = "?"
	} else {
		t.current = 8
		t.text = "AD"
	}
	t.outlineText.Refresh(t.text)

	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableIconWithBottomCenteredString) Tapped(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.changeResource, t.changeResource)
	t.changeResource()
}

func (t *TappableIconWithBottomCenteredString) TappedSecondary(_ *fyne.PointEvent) {
	if t.current <= 8 {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}

func (t *TappableIconWithBottomCenteredString) Keyed() {
	t.undoRedoStacks.StoreFunctions(t.changeResource, t.changeResource)
	t.changeResource()
}

func (t *TappableIconWithBottomCenteredString) MouseIn(event *desktop.MouseEvent) {
	//t.toolTipPopUp = newToolTipTextTappableIconWithBottomCenteredString(event, t.toolTipText, t)
}

func (t *TappableIconWithBottomCenteredString) MouseMoved(_ *desktop.MouseEvent) {
}

func (t *TappableIconWithBottomCenteredString) MouseOut() {
	//t.toolTipPopUp.Hide()
}

func newToolTipTextTappableIconWithBottomCenteredString(event *desktop.MouseEvent, text string, object *TappableIconWithBottomCenteredString) *widget.PopUp {
	toolTipText := canvas.NewText(text, color.White)
	popUp := widget.NewPopUp(toolTipText, fyne.CurrentApp().Driver().CanvasForObject(object))
	var popUpPosition fyne.Position
	popUpPosition.X = event.AbsolutePosition.X + object.Size().Width/2
	popUpPosition.Y = event.AbsolutePosition.Y - object.Size().Height/2
	popUp.ShowAtPosition(popUpPosition)

	return popUp
}
