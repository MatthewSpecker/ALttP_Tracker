package tappable_icons

import (
	"errors"
	"image/color"
	"strconv"

	"tracker/save"
	"tracker/text_outline"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TappableIconWithCenteredNum struct {
	widget.Icon
	resources      []fyne.Resource
	current        int
	number         int
	numberMax      int
	text           string
	outlineText    *text_outline.TextOutline
	tapSize        float32
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableIconWithCenteredNum(res []fyne.Resource, num int, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableIconWithCenteredNum, error) {
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

	icon := &TappableIconWithCenteredNum{
		resources:      res,
		current:        0,
		number:         0,
		numberMax:      num,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	if icon.current == 0 {
		icon.text = ""
	} else {
		icon.text = strconv.Itoa(icon.number)
	}

	icon.outlineText = text_outline.NewTextOutline(icon.text, icon.tapSize, 1.5, 6, color.White, color.Black)
	icon.ExtendBaseWidget(icon)
	icon.SetResource(icon.resources[icon.current])

	return icon, nil
}

func (t *TappableIconWithCenteredNum) Update() {
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, len(t.resources)-1, 0)
	t.number = t.saveFile.GetSaveInt(t.saveFileText + "_Number")
	t.number = intRangeCheck(t.number, t.numberMax, 0)
	if t.number == 0 {
		t.text = ""
	} else {
		t.text = strconv.Itoa(t.number)
	}

	t.outlineText.Refresh(t.text)

	t.Icon.SetResource(t.resources[t.current])
}

func (t *TappableIconWithCenteredNum) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Current", 0)
	t.saveFile.SetDefault(t.saveFileText+"_Number", 0)
}

func (t *TappableIconWithCenteredNum) GetSaveDefaults() {
	t.current = 0
	t.number = 0
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
	t.Update()
}

func (t *TappableIconWithCenteredNum) Layout() *fyne.Container {
	tapIconContainer := t.layoutCenteredNum()
	tapIconContainer = container.New(layout.NewCenterLayout(), tapIconContainer)
	return tapIconContainer
}

func (t *TappableIconWithCenteredNum) layoutCenteredNum() *fyne.Container {
	container1 := container.New(layout.NewCenterLayout(), t)
	container2 := container.New(layout.NewCenterLayout(), t.outlineText.TextOutlineContainer)
	container3 := container.NewWithoutLayout(container1, container2)
	iconSize := t.Size()
	container2Size := t.outlineText.TextOutlineContainer.Size()
	iconChangePosition := container2.Position()
	iconChangePosition = iconChangePosition.AddXY(iconSize.Width/2-container2Size.Width/2, iconSize.Height*0.6-container2Size.Height/2)
	container2.Move(iconChangePosition)

	return container3
}

func (t *TappableIconWithCenteredNum) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableIconWithCenteredNum) increment() {
	if t.current < (len(t.resources) - 1) {
		t.current++
		t.Icon.SetResource(t.resources[t.current])
	}
	if t.number < t.numberMax {
		t.number++
		t.text = strconv.Itoa(t.number)
		t.outlineText.Refresh(t.text)
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
}

func (t *TappableIconWithCenteredNum) decrement() {
	if t.current > 0 && t.number <= t.current {
		t.current--
		t.Icon.SetResource(t.resources[t.current])
	}
	if t.number == 1 {
		t.number--
		t.text = ""
		t.outlineText.Refresh(t.text)
	} else if t.number > 1 {
		t.number--
		t.text = strconv.Itoa(t.number)
		t.outlineText.Refresh(t.text)
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
}

func (t *TappableIconWithCenteredNum) Tapped(_ *fyne.PointEvent) {
	if t.current < (len(t.resources) - 1) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}

func (t *TappableIconWithCenteredNum) TappedSecondary(_ *fyne.PointEvent) {
	if t.current > 0 {
		t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
	}
	t.decrement()
}

func (t *TappableIconWithCenteredNum) Keyed() {
	if t.current < (len(t.resources) - 1) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}
