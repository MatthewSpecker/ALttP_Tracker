package tappable_icons

import (
  "errors"
  "strconv"
  "image/color"

  "tracker/save"
  "tracker/tooltip"
  "tracker/undo_redo"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/canvas"
  "fyne.io/fyne/v2/container"
  "fyne.io/fyne/v2/layout"
  "fyne.io/fyne/v2/theme"
  "fyne.io/fyne/v2/driver/desktop"

)

type TappableNumIconWithIcon struct {
  widget.Icon
  desktop.Hoverable
  text *canvas.Text
  resources []fyne.Resource
  current int
  iconSmall *widget.Icon
  number int
  numberMax int
  ascending bool
  tapSize float32
  toolTipText string
  toolTipPopUp *widget.PopUp
  undoRedoStacks *undo_redo.UndoRedoStacks
  saveFile *save.SaveFile
  saveFileText string
}

func NewTappableNumIconWithIcon(res []fyne.Resource, num int, increase bool, size float32, undoRedo *undo_redo.UndoRedoStacks, save *save.SaveFile, saveName string) (*TappableNumIconWithIcon, error) {
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

  icon := &TappableNumIconWithIcon{
    resources: res,
    current: 0,
    number: 0,
    numberMax: num,
    ascending: increase,
    tapSize: size,
    undoRedoStacks: undoRedo,
    saveFile: save,
    saveFileText: saveName,
  }

  icon.iconSmall = widget.NewIcon(icon.resources[icon.current])
  icon.text = canvas.NewText(strconv.Itoa(icon.number), color.White)
  icon.toolTipText = tooltip.GetToolTipText(icon.resources[icon.current].Name())

  icon.text.TextStyle.Bold = true
  icon.text.TextSize = size * theme.Padding()/2.5
  if icon.ascending == false {
    icon.number = icon.numberMax
    icon.text.Text = strconv.Itoa(icon.number)
  }
  if icon.numberMax == 0 {
    icon.text.Color = color.NRGBA{R:0, G:255, B:0, A:255}
  }

  icon.text.Refresh()
  icon.ExtendBaseWidget(icon)
  resEmpty, _ := fyne.LoadResourceFromPath("")
  icon.SetResource(resEmpty)

  return icon, nil
}

func (t *TappableNumIconWithIcon) Update() {
  t.number = t.saveFile.GetSaveInt(t.saveFileText + "_Number")
  t.number = intRangeCheck(t.number, t.numberMax, 0)
  t.text.Text = strconv.Itoa(t.number)

  if t.ascending {
    if t.number == t.numberMax {
      t.text.Color = color.NRGBA{R:0, G:255, B:0, A:255}
    } else {
      t.text.Color = color.White
    }
    if t.number == 0 && t.numberMax != 0{
      t.current = 0
    } else {
      t.current = 1
    }
  } else {
    if t.number == 0 {
      t.text.Color = color.NRGBA{R:0, G:255, B:0, A:255}
    } else {
      t.text.Color = color.White
    }
    if t.number == 0 {
      t.current = 1
    } else {
      t.current = 0
    }
  }

  t.iconSmall.SetResource(t.resources[t.current])
  t.text.Refresh()
}

func (t *TappableNumIconWithIcon) SetSaveDefaults() {
  if t.ascending {
    t.saveFile.SetDefault(t.saveFileText + "_Number", 0)
  } else {
    t.saveFile.SetDefault(t.saveFileText + "_Number", t.numberMax)
  }
}

func (t *TappableNumIconWithIcon) GetSaveDefaults() {
  if t.ascending {
    t.number = 0
  } else {
    t.number = t.numberMax
  }
  t.saveFile.SetSave(t.saveFileText + "_Number", t.number)
  t.Update()
}

func (t *TappableNumIconWithIcon) LayoutAdjust() *fyne.Container {
  container1 := container.New(layout.NewCenterLayout(), t)
  container2 := container.New(layout.NewCenterLayout(), t.text)
  container3 := container.New(layout.NewMaxLayout(), t.iconSmall)
  container4 := container.NewWithoutLayout(container3, container2, container1)
  container3Size := container3.Size()
  container3Resize := fyne.NewSize(container3Size.Width/1.5, container3Size.Height/1.5)
  container3.Resize(container3Resize)
  iconSize := t.text.Size()
  iconChangePosition := fyne.NewPos(iconSize.Width/2, iconSize.Height/2)
  container3.Move(iconChangePosition)

  return container4
}

func (t *TappableNumIconWithIcon) MinSize() fyne.Size {
  return fyne.NewSize(theme.Padding()*t.tapSize/2, theme.Padding()*t.tapSize/2)
}

func (t *TappableNumIconWithIcon) increment() {
  if t.number < t.numberMax {
    t.number++
    t.text.Text = (strconv.Itoa(t.number))
    if t.number == t.numberMax && t.ascending == true {
      t.text.Color = color.NRGBA{R:0, G:255, B:0, A:255}
      t.current++
      t.iconSmall.SetResource(t.resources[t.current])
    }
    if t.number == 1 && t.ascending == false {
      t.text.Color = color.White
      t.current--
      t.iconSmall.SetResource(t.resources[t.current])
    }
    t.text.Refresh()
  }
  t.saveFile.SetSave(t.saveFileText + "_Number", t.number)
}

func (t *TappableNumIconWithIcon) decrement() {
  if t.number > 0 {
    t.number--
    t.text.Text = (strconv.Itoa(t.number))
    if t.number == t.numberMax - 1 && t.ascending == true {
      t.text.Color = color.White
      t.current--
      t.iconSmall.SetResource(t.resources[t.current])
    }
    if t.number == 0 && t.ascending == false {
      t.text.Color = color.NRGBA{R:0, G:255, B:0, A:255}
      t.current++
      t.iconSmall.SetResource(t.resources[t.current])
    }
    t.text.Refresh()
  }
  t.saveFile.SetSave(t.saveFileText + "_Number", t.number)
}

func (t *TappableNumIconWithIcon) Tapped(_ *fyne.PointEvent) {
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

func (t *TappableNumIconWithIcon) TappedSecondary(_ *fyne.PointEvent) {
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

func (t *TappableNumIconWithIcon) Keyed() {
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

func (t *TappableNumIconWithIcon) MouseIn(event *desktop.MouseEvent) {
  //t.toolTipPopUp = newToolTipTextTappableNumIconWithIcon(event, t.toolTipText, t)
}

func (t *TappableNumIconWithIcon) MouseMoved(_ *desktop.MouseEvent) {
}

func (t *TappableNumIconWithIcon) MouseOut() {
  //t.toolTipPopUp.Hide()
}

func newToolTipTextTappableNumIconWithIcon (event *desktop.MouseEvent, text string, object *TappableNumIconWithIcon) *widget.PopUp {
  toolTipText := canvas.NewText(text, color.White)
  popUp := widget.NewPopUp(toolTipText, fyne.CurrentApp().Driver().CanvasForObject(object))
  var popUpPosition fyne.Position
  popUpPosition.X = event.AbsolutePosition.X + object.Size().Width/2
  popUpPosition.Y = event.AbsolutePosition.Y - object.Size().Height/2
  popUp.ShowAtPosition(popUpPosition)

  return popUp
}