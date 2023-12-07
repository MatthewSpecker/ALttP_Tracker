package tappable_icons

import (
	"errors"

	"tracker/save"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TappableIcon struct {
	widget.Icon
	resources      []fyne.Resource
	current        int
	tapSize        float32
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableIcon(res []fyne.Resource, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableIcon, error) {
	if len(res) <= 1 {
		return nil, errors.New("'res' must contain 2 or more resources")
	}
	if size <= 0 {
		return nil, errors.New("'size' must be greater than 0")
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	icon := &TappableIcon{
		resources:      res,
		current:        0,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	icon.ExtendBaseWidget(icon)
	icon.SetResource(icon.resources[icon.current])

	return icon, nil
}

func (t *TappableIcon) Update() {
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, len(t.resources)-1, 0)

	t.Icon.SetResource(t.resources[t.current])
}

func (t *TappableIcon) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Current", 0)
}

func (t *TappableIcon) GetSaveDefaults() {
	t.current = 0
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.Update()
}

func (t *TappableIcon) Layout() *fyne.Container {
	tapIconContainer := container.New(layout.NewCenterLayout(), t)
	return tapIconContainer
}

func (t *TappableIcon) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableIcon) increment() {
	if t.current < (len(t.resources) - 1) {
		t.current++
		t.Icon.SetResource(t.resources[t.current])
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableIcon) decrement() {
	if t.current > 0 {
		t.current--
		t.Icon.SetResource(t.resources[t.current])
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableIcon) Tapped(_ *fyne.PointEvent) {
	if t.current < (len(t.resources) - 1) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}

func (t *TappableIcon) TappedSecondary(_ *fyne.PointEvent) {
	if t.current > 0 {
		t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
	}
	t.decrement()
}

func (t *TappableIcon) Keyed() {
	if t.current < (len(t.resources) - 1) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}
