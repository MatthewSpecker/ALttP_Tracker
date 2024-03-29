package tappable_icons

import (
	"errors"
	"fmt"

	"tracker/save"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TappableIconVariedSize struct {
	widget.Icon
	resources      []fyne.Resource
	current        int
	tapSizePath    []float32
	tapSize        float32
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableIconVariedSize(res []fyne.Resource, size []float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableIconVariedSize, error) {
	if len(res) <= 1 {
		return nil, errors.New("'res' must contain 2 or more resources")
	}
	for index, element := range size {
		if element <= 0 {
			return nil, errors.New(fmt.Sprintf("'size[%d]' must be float32 greater than 0", index))
		}
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	icon := &TappableIconVariedSize{
		resources:      res,
		current:        0,
		tapSizePath:    size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	icon.tapSize = icon.tapSizePath[icon.current]
	icon.ExtendBaseWidget(icon)
	icon.SetResource(icon.resources[icon.current])

	return icon, nil
}

func (t *TappableIconVariedSize) Update() {
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, len(t.resources)-1, 0)
	t.tapSize = t.tapSizePath[t.current]

	t.Icon.SetResource(t.resources[t.current])
}

func (t *TappableIconVariedSize) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Current", 0)
}

func (t *TappableIconVariedSize) GetSaveDefaults() {
	t.current = 0
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.Update()
}

func (t *TappableIconVariedSize) Layout() *fyne.Container {
	tapIconContainer := container.New(layout.NewCenterLayout(), t)
	return tapIconContainer
}

func (t *TappableIconVariedSize) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableIconVariedSize) increment() {
	if t.current < (len(t.resources) - 1) {
		t.current++
		t.tapSize = t.tapSizePath[t.current]
		t.Icon.SetResource(t.resources[t.current])
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableIconVariedSize) decrement() {
	if t.current > 0 {
		t.current--
		t.tapSize = t.tapSizePath[t.current]
		t.Icon.SetResource(t.resources[t.current])
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableIconVariedSize) Tapped(_ *fyne.PointEvent) {
	if t.current < (len(t.resources) - 1) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}

func (t *TappableIconVariedSize) TappedSecondary(_ *fyne.PointEvent) {
	if t.current > 0 {
		t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
	}
	t.decrement()
}

func (t *TappableIconVariedSize) Keyed() {
	if t.current < (len(t.resources) - 1) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}
