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

type TappableBossIcon struct {
	widget.Icon
	desktop.Hoverable
	resourcesGray   []fyne.Resource
	resourcesBoss   []fyne.Resource
	current         int
	bossStart       int
	passedPrizeIcon *TappablePrizeIcon
	tapSize         float32
	toolTipText     string
	toolTipPopUp    *tooltip.ToolTip
	undoRedoStacks  *undo_redo.UndoRedoStacks
	saveFile        *save.SaveFile
	saveFileText    string
}

func NewTappableBossIcon(bossNum int, size float32, prizeIcon *TappablePrizeIcon, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableBossIcon, error) {
	if bossNum < 0 || bossNum > 9 {
		return nil, errors.New("'bossNum' must be int from 0 to 9")
	}
	if size <= 0 {
		return nil, errors.New("'size' must be float32 greater than 0")
	}
	if saveName == "" {
		return nil, errors.New("'saveName' cannot be empty string")
	}

	icon := &TappableBossIcon{
		resourcesGray:   []fyne.Resource{resourceArmosGrayPng, resourceLanmolasGrayPng, resourceMoldormGrayPng, resourceHelmasaurGrayPng, resourceArrghusGrayPng, resourceMothulaGrayPng, resourceBlindGrayPng, resourceKholdstareGrayPng, resourceVitreousGrayPng, resourceTrinexxGrayPng},
		resourcesBoss:   []fyne.Resource{resourceArmosPng, resourceLanmolasPng, resourceMoldormPng, resourceHelmasaurPng, resourceArrghusPng, resourceMothulaPng, resourceBlindPng, resourceKholdstarePng, resourceVitreousPng, resourceTrinexxPng},
		current:         bossNum,
		bossStart:       bossNum,
		passedPrizeIcon: prizeIcon,
		tapSize:         size,
		undoRedoStacks:  undoRedo,
		saveFile:        save,
		saveFileText:    saveName + "_Boss",
	}

	icon.toolTipText = tooltip.GetToolTipText(icon.resourcesBoss[icon.current].Name())
	icon.passedPrizeIcon.bossIcon = icon

	icon.ExtendBaseWidget(icon)
	if icon.passedPrizeIcon.obtained {
		icon.SetResource(icon.resourcesBoss[icon.current])
	} else {
		icon.SetResource(icon.resourcesGray[icon.current])
	}

	return icon, nil
}

func (t *TappableBossIcon) Update() {
	t.current = t.saveFile.GetSaveInt(t.saveFileText + "_Current")
	t.current = intRangeCheck(t.current, len(t.resourcesBoss)-1, 0)
	t.toolTipText = tooltip.GetToolTipText(t.resourcesBoss[t.current].Name())

	if t.passedPrizeIcon.obtained {
		t.Icon.SetResource(t.resourcesBoss[t.current])
	} else {
		t.Icon.SetResource(t.resourcesGray[t.current])
	}
}

func (t *TappableBossIcon) SetSaveDefaults() {
	t.saveFile.SetDefault(t.saveFileText+"_Current", t.bossStart)
}

func (t *TappableBossIcon) GetSaveDefaults() {
	t.current = t.bossStart
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
	t.Update()
}

func (t *TappableBossIcon) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableBossIcon) bossIncrement() {
	if t.current == len(t.resourcesBoss)-1 {
		t.current = 0
	} else {
		t.current++
	}
	if t.passedPrizeIcon.obtained == false {
		t.Icon.SetResource(t.resourcesGray[t.current])
	} else {
		t.Icon.SetResource(t.resourcesBoss[t.current])
	}
	t.toolTipText = tooltip.GetToolTipText(t.resourcesBoss[t.current].Name())
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableBossIcon) bossDecrement() {
	if t.current == 0 {
		t.current = len(t.resourcesBoss) - 1
	} else {
		t.current--
	}
	if t.passedPrizeIcon.obtained == false {
		t.Icon.SetResource(t.resourcesGray[t.current])
	} else {
		t.Icon.SetResource(t.resourcesBoss[t.current])
	}
	t.toolTipText = tooltip.GetToolTipText(t.resourcesBoss[t.current].Name())
	t.saveFile.SetSave(t.saveFileText+"_Current", t.current)
}

func (t *TappableBossIcon) bossSet() {
	if t.passedPrizeIcon.obtained == false {
		t.Icon.SetResource(t.resourcesGray[t.current])
	} else {
		t.Icon.SetResource(t.resourcesBoss[t.current])
	}
}

func (t *TappableBossIcon) Tapped(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.bossDecrement, t.bossIncrement)
	t.bossIncrement()
}

func (t *TappableBossIcon) TappedSecondary(_ *fyne.PointEvent) {
	t.undoRedoStacks.StoreFunctions(t.bossIncrement, t.bossDecrement)
	t.bossDecrement()
}

func (t *TappableBossIcon) Keyed() {
	//t.undoRedoStacks.StoreFunctions(t.PrizeSet, t.PrizeSet)
	//t.PrizeSet()
}

func (t *TappableBossIcon) MouseIn(event *desktop.MouseEvent) {
	//t.toolTipPopUp = newToolTipTextTappableBossIcon(event, t.toolTipText, t)
}

func (t *TappableBossIcon) MouseMoved(_ *desktop.MouseEvent) {
}

func (t *TappableBossIcon) MouseOut() {
	//t.toolTipPopUp.Hide()
}

func newToolTipTextTappableBossIcon(event *desktop.MouseEvent, text string, object *TappableBossIcon) *tooltip.ToolTip {
	toolTipText := canvas.NewText(text, color.White)
	toolTip := tooltip.NewToolTip(toolTipText, fyne.CurrentApp().Driver().CanvasForObject(object), object.Tapped, object.TappedSecondary)
	var toolTipPosition fyne.Position
	toolTipPosition.X = event.AbsolutePosition.X + object.Size().Width/2
	toolTipPosition.Y = event.AbsolutePosition.Y - object.Size().Height/2
	toolTip.ShowAtPosition(toolTipPosition)

	return toolTip
}
