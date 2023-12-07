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

type TappableTempIconWithIcon struct {
	widget.Icon
	resources      []fyne.Resource
	current        int
	smallResources []fyne.Resource
	iconSmall      *SizeableIcon
	tapSize        float32
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableTempIconWithIcon(res []fyne.Resource, smallRes []fyne.Resource, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableTempIconWithIcon, error) {
	if len(res) <= 1 {
		return nil, errors.New("'res' must contain 2 or more resources")
	}
	if len(smallRes) == 0 {
		return nil, errors.New("'smallRes' must contain 1 or more resources")
	}
	if size <= 0 {
		return nil, errors.New("'size' must be float32 greater than 0")
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	resSmallEmpty, _ := fyne.LoadResourceFromPath("")
	resSmallEmptySlice := []fyne.Resource{resSmallEmpty}
	smallRes = append(resSmallEmptySlice, smallRes...)

	icon := &TappableTempIconWithIcon{
		resources:      res,
		current:        0,
		smallResources: smallRes,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	icon.ExtendBaseWidget(icon)
	icon.iconSmall, _ = NewSizeableIcon(smallRes, size)
	icon.SetResource(icon.resources[icon.current])

	return icon, nil
}

func (t *TappableTempIconWithIcon) Update() {
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, len(t.smallResources)+len(t.resources)-1, 0)

	if t.current == len(t.smallResources)-1+len(t.resources) {
		t.iconSmall.Update(t.current-len(t.resources))
		t.SetResource(t.resources[len(t.resources)-1])
	} else if t.current > len(t.resources)-1 {
		t.iconSmall.Update(t.current-len(t.resources)+1)
		t.SetResource(t.resources[0])
	} else {
		t.iconSmall.Update(0)
		t.SetResource(t.resources[t.current])
	}
}

func (t *TappableTempIconWithIcon) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Current", 0)
}

func (t *TappableTempIconWithIcon) GetSaveDefaults() {
	t.current = 0
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.Update()
}

func (t *TappableTempIconWithIcon) Layout() *fyne.Container {
	tapIconContainer := t.layoutIcon()
	tapIconContainer = container.New(layout.NewCenterLayout(), tapIconContainer)
	return tapIconContainer
}

func (t *TappableTempIconWithIcon) layoutIcon() *fyne.Container {
	container1 := container.New(layout.NewCenterLayout(), t)
	container2 := container.New(layout.NewCenterLayout(), t.iconSmall)
	container3 := container.NewWithoutLayout(container1, container2)
	iconSize := t.Size()
	iconChangePosition := fyne.NewPos(iconSize.Width/2, iconSize.Height/2)
	container2.Move(iconChangePosition)

	return container3
}

func (t *TappableTempIconWithIcon) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableTempIconWithIcon) increment() {
	if t.current < (len(t.resources)-1) {
		t.current++
		t.Icon.SetResource(t.resources[t.current])
	} else if t.current < len(t.smallResources)-1+len(t.resources)-1 {
		t.current++
		t.Icon.SetResource(t.resources[0])
		t.iconSmall.Update(t.current-len(t.resources)+1)
	} else if t. current == len(t.smallResources)-1+len(t.resources)-1 {
		t.current++
		t.Icon.SetResource(t.resources[len(t.resources)-1])
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableTempIconWithIcon) decrement() {
	if t.current == (len(t.resources)+len(t.smallResources)-1) {
		t.current--
		t.Icon.SetResource(t.resources[0])
	} else if t.current > len(t.resources) {
		t.current--
		t.iconSmall.Update(t.current-len(t.resources)+1)
	} else if t.current == len(t.resources) {
		t.current--
		t.Icon.SetResource(t.resources[len(t.resources)-1])
		t.iconSmall.Update(0)
	} else if t.current > 0 {
		t.current--
		t.Icon.SetResource(t.resources[t.current])
	}
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableTempIconWithIcon) Tapped(_ *fyne.PointEvent) {
	if t.current < len(t.resources)+len(t.smallResources) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}

func (t *TappableTempIconWithIcon) TappedSecondary(_ *fyne.PointEvent) {
	if t.current > 0 {
		t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
	}
	t.decrement()
}

func (t *TappableTempIconWithIcon) Keyed() {
	if t.current < len(t.resources)+len(t.smallResources) {
		t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
	}
	t.increment()
}
