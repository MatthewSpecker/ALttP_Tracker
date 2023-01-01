package tappable_icons

import (
	"errors"
	"image/color"

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

type TappablePrizeIcon struct {
	widget.Icon
	desktop.Hoverable
	resourcesGray  []fyne.Resource
	resourcesPrize []fyne.Resource
	current        int
	obtained       bool
	bossIcon       *TappableBossIcon
	tapSize        float32
	toolTipText    string
	toolTipPopUp   *tooltip.ToolTip
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappablePrizeIcon(size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappablePrizeIcon, error) {
	if size <= 0 {
		return nil, errors.New("'size' must be float32 greater than 0")
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	icon := &TappablePrizeIcon{
		resourcesGray:  []fyne.Resource{resourceBlueCrystalGrayPng, resourceRedCrystalGrayPng, resourceBlueRedPendantGrayPng, resourceGreenPendantGrayPng},
		resourcesPrize: []fyne.Resource{resourceBlueCrystalPng, resourceRedCrystalPng, resourceBlueRedPendantPng, resourceGreenPendantPng},
		current:        0,
		obtained:       false,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName + "_Prize",
	}

	icon.toolTipText = tooltip.GetToolTipText(icon.resourcesPrize[icon.current].Name())
	icon.ExtendBaseWidget(icon)
	icon.SetResource(icon.resourcesGray[icon.current])

	return icon, nil
}

func (t *TappablePrizeIcon) Update() {
	t.obtained = t.saveFile.GetSaveBool(t.saveFileText + "_Obtained")
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, len(t.resourcesPrize)-1, 0)
	t.toolTipText = tooltip.GetToolTipText(t.resourcesPrize[t.current].Name())

	if t.obtained {
		t.Icon.SetResource(t.resourcesPrize[t.current])
	} else {
		t.Icon.SetResource(t.resourcesGray[t.current])
	}
}

func (t *TappablePrizeIcon) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Obtained", false)
	t.saveFile.SetDefault(t.saveFileText+"_Current", 0)
}

func (t *TappablePrizeIcon) GetSaveDefaults() {
	t.obtained = false
	t.current = 0
	t.saveFile.SetSave(t.saveFileText+"_Obtained", t.obtained)
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.Update()
}

func (t *TappablePrizeIcon) Layout() *fyne.Container {
	tapIconContainer := container.New(layout.NewCenterLayout(), t)
	return tapIconContainer
}

func (t *TappablePrizeIcon) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappablePrizeIcon) prizeIncrement() {
	if t.current == len(t.resourcesPrize)-1 {
		t.current = 0
	} else {
		t.current++
	}
	if t.obtained == false {
		t.Icon.SetResource(t.resourcesGray[t.current])
	} else {
		t.Icon.SetResource(t.resourcesPrize[t.current])
	}
	t.toolTipText = tooltip.GetToolTipText(t.resourcesPrize[t.current].Name())
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappablePrizeIcon) prizeDecrement() {
	if t.current == 0 {
		t.current = len(t.resourcesPrize) - 1
	} else {
		t.current--
	}
	if t.obtained == false {
		t.Icon.SetResource(t.resourcesGray[t.current])
	} else {
		t.Icon.SetResource(t.resourcesPrize[t.current])
	}
	t.toolTipText = tooltip.GetToolTipText(t.resourcesPrize[t.current].Name())
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappablePrizeIcon) prizeSet() {
	if t.obtained == true {
		t.obtained = false
		t.Icon.SetResource(t.resourcesGray[t.current])
	} else {
		t.obtained = true
		t.Icon.SetResource(t.resourcesPrize[t.current])
	}

	if t.bossIcon != nil {
		t.bossIcon.bossSet()
	}
	t.saveFile.SetSave(t.saveFileText+"_Obtained", t.obtained)
}

func (t *TappablePrizeIcon) Tapped(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.prizeSet, t.prizeSet)
	t.prizeSet()
}

func (t *TappablePrizeIcon) TappedSecondary(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.prizeDecrement, t.prizeIncrement)
	t.prizeIncrement()
}

func (t *TappablePrizeIcon) Keyed() {
	t.undoRedoStacks.StoreFunctions(t.prizeSet, t.prizeSet)
	t.prizeSet()
}

func (t *TappablePrizeIcon) MouseIn(event *desktop.MouseEvent) {
	//t.toolTipPopUp = newToolTipTextTappablePrizeIcon(event, t.toolTipText, t)
}

func (t *TappablePrizeIcon) MouseMoved(_ *desktop.MouseEvent) {
}

func (t *TappablePrizeIcon) MouseOut() {
	//t.toolTipPopUp.Hide()
}

func newToolTipTextTappablePrizeIcon(event *desktop.MouseEvent, text string, object *TappablePrizeIcon) *tooltip.ToolTip {
	toolTipText := canvas.NewText(text, color.White)
	toolTip := tooltip.NewToolTip(toolTipText, fyne.CurrentApp().Driver().CanvasForObject(object), object.Tapped, object.TappedSecondary)
	var toolTipPosition fyne.Position
	toolTipPosition.X = event.AbsolutePosition.X + object.Size().Width/2
	toolTipPosition.Y = event.AbsolutePosition.Y - object.Size().Height/2
	toolTip.ShowAtPosition(toolTipPosition)

	return toolTip
}
