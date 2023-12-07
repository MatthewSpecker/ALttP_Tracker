package tappable_icons

import (
	"errors"
	"image/color"
	"strconv"

	"tracker/save"
	"tracker/undo_redo"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type TappableNumIconWithNum struct {
	widget.Icon
	text           *canvas.Text
	number         int
	numberMax      int
	numberLabel    *widget.Label
	ascending      bool
	tapSize        float32
	undoRedoStacks *undo_redo.UndoRedoStacks
	saveFile       *save.SaveFile
	saveFileText   string
}

func NewTappableNumIconWithNum(name string, num int, increase bool, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableNumIconWithNum, error) {
	if name == "" {
		return nil, errors.New("'name' cannot be empty string")
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

	icon := &TappableNumIconWithNum{
		number:         0,
		numberMax:      num,
		ascending:      increase,
		tapSize:        size,
		undoRedoStacks: undoRedo,
		saveFile:       save,
		saveFileText:   saveName,
	}

	icon.text = canvas.NewText(strconv.Itoa(icon.number), color.White)
	icon.numberLabel = widget.NewLabel(strconv.Itoa(icon.numberMax))

	icon.text.TextStyle.Bold = true
	icon.text.TextSize = size * theme.Padding() / 2
	if icon.ascending == false {
		icon.number = icon.numberMax
		icon.text.Text = strconv.Itoa(icon.number)
	}
	if icon.numberMax == 0 {
		icon.text.Color = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	}

	icon.ExtendBaseWidget(icon)
	res, _ := fyne.LoadResourceFromPath("")
	icon.SetResource(res)

	return icon, nil
}

func (t *TappableNumIconWithNum) Update() {
	t.number = t.saveFile.GetSaveInt(t.saveFileText + "_Number")
	t.number = intRangeCheck(t.number, t.numberMax, 0)
	t.text.Text = strconv.Itoa(t.number)

	if t.ascending {
		if t.number == t.numberMax {
			t.text.Color = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
		} else {
			t.text.Color = color.White
		}
	} else {
		if t.number == 0 {
			t.text.Color = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
		} else {
			t.text.Color = color.White
		}
	}

	t.text.Refresh()
}

func (t *TappableNumIconWithNum) SetSaveDefaults() {
	if t.ascending {
		t.saveFile.SetDefault(t.saveFileText+"_Number", 0)
	} else {
		t.saveFile.SetDefault(t.saveFileText+"_Number", t.numberMax)
	}
}

func (t *TappableNumIconWithNum) GetSaveDefaults() {
	if t.ascending {
		t.number = 0
	} else {
		t.number = t.numberMax
	}
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
	t.Update()
}

func (t *TappableNumIconWithNum) Layout() *fyne.Container {
	tapIconContainer := t.layoutNum()
	tapIconContainer = container.New(layout.NewCenterLayout(), tapIconContainer)
	return tapIconContainer
}

func (t *TappableNumIconWithNum) layoutNum() *fyne.Container {
	container1 := container.New(layout.NewCenterLayout(), t)
	container2 := container.New(layout.NewCenterLayout(), t.text)
	container3 := container.New(layout.NewCenterLayout(), t.numberLabel)
	container4 := container.NewWithoutLayout(container1, container2, container3)
	iconSize := t.Size()
	iconChangePosition := fyne.NewPos(iconSize.Width/2, iconSize.Height)
	container3.Move(iconChangePosition)

	return container4
}

func (t *TappableNumIconWithNum) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableNumIconWithNum) increment() {
	if t.number < t.numberMax {
		t.number++
		t.text.Text = (strconv.Itoa(t.number))
		if t.number == t.numberMax && t.ascending == true {
			t.text.Color = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
		}
		if t.number == 1 && t.ascending == false {
			t.text.Color = color.White
		}
		t.text.Refresh()
	}
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
}

func (t *TappableNumIconWithNum) decrement() {
	if t.number > 0 {
		t.number--
		t.text.Text = (strconv.Itoa(t.number))
		if t.number == t.numberMax-1 && t.ascending == true {
			t.text.Color = color.White
		}
		if t.number == 0 && t.ascending == false {
			t.text.Color = color.NRGBA{R: 0, G: 255, B: 0, A: 255}
		}
		t.text.Refresh()
	}
	t.saveFile.SetSave(t.saveFileText+"_Number", t.number)
}

func (t *TappableNumIconWithNum) Tapped(_ *fyne.PointEvent) {
	if t.ascending {
		if t.number < t.numberMax {
			t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
		}
		t.increment()
	} else {
		if t.number > 0 {
			t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
		}
		t.decrement()
	}
}

func (t *TappableNumIconWithNum) TappedSecondary(_ *fyne.PointEvent) {
	if t.ascending {
		if t.number > 0 {
			t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
		}
		t.decrement()
	} else {
		if t.number < t.numberMax {
			t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
		}
		t.increment()
	}
}

func (t *TappableNumIconWithNum) Keyed() {
	if t.ascending {
		if t.number < t.numberMax {
			t.undoRedoStacks.StoreFunctions(t.decrement, t.increment)
		}
		t.increment()
	} else {
		if t.number > 0 {
			t.undoRedoStacks.StoreFunctions(t.increment, t.decrement)
		}
		t.decrement()
	}
}
