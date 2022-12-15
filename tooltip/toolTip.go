package tooltip

import (
  "strings"

  "fyne.io/fyne/v2"
  "fyne.io/fyne/v2/canvas"
  "fyne.io/fyne/v2/widget"
  "fyne.io/fyne/v2/theme"
)

func GetToolTipText(text string) string {
  toolTipText := strings.Replace(text, "_", " ", -1)
  toolTipText = strings.Replace(toolTipText, ".png", "", 1)
  toolTipText = strings.Replace(toolTipText, "Gray", "", -1)

  return toolTipText
}

type ToolTip struct {
  widget.BaseWidget

  Content fyne.CanvasObject
  Canvas  fyne.Canvas
  TapFunction func (_ *fyne.PointEvent)
  TapSecFunction func(_ *fyne.PointEvent)

  innerPos     fyne.Position
  innerSize    fyne.Size
  overlayShown bool
}

// Hide this widget, if it was previously visible
func (t *ToolTip) Hide() {
  if t.overlayShown {
    t.Canvas.Overlays().Remove(t)
    t.overlayShown = false
  }
  t.BaseWidget.Hide()
}

// Move the widget to a new position. A ToolTip position is absolute to the top, left of its canvas.
// For ToolTip this actually moves the content so checking Position() will not return the same value as is set here.
func (t *ToolTip) Move(pos fyne.Position) {
  t.innerPos = pos
  t.Refresh()
}

// Resize changes the size of the ToolTip's content.
// ToolTips always have the size of their canvas, but this call updates the
// size of the content portion.
//
// Implements: fyne.Widget
func (t *ToolTip) Resize(size fyne.Size) {
  t.innerSize = size
  // The canvas size might not have changed and therefore the Resize won't trigger a layout.
  // Until we have a widget.Relayout() or similar, the renderer's refresh will do the re-layout.
  t.Refresh()
}

// Show this pop-up as overlay if not already shown.
func (t *ToolTip) Show() {
  if !t.overlayShown {
    t.Canvas.Overlays().Add(t)
    t.overlayShown = true
  }
  t.Refresh()
  t.BaseWidget.Show()
}

// ShowAtPosition shows this pop-up at the given position.
func (t *ToolTip) ShowAtPosition(pos fyne.Position) {
  t.Move(pos)
  t.Show()
}

// MinSize returns the size that this widget should not shrink below
func (t *ToolTip) MinSize() fyne.Size {
  t.ExtendBaseWidget(t)
  return t.BaseWidget.MinSize()
}

// ShowToolTipAtPosition creates a new ToolTip for the specified content at the specified absolute position.
// It will then display the ToolTip on the passed canvas.
func ShowToolTipAtPosition(content fyne.CanvasObject, canvas fyne.Canvas, tap func (_ *fyne.PointEvent), tapSec func(_ *fyne.PointEvent), pos fyne.Position) {
  newToolTip(content, canvas, tap, tapSec).ShowAtPosition(pos)
}

// Tapped is called when the user taps the ToolTip background - if not modal then dismiss this widget
func (t *ToolTip) Tapped(point *fyne.PointEvent) {
  t.BaseWidget.Hide()
  t.TapFunction(point)
}

// TappedSecondary is called when the user right/alt taps the background - if not modal then dismiss this widget
func (t *ToolTip) TappedSecondary(point *fyne.PointEvent) {
  t.BaseWidget.Hide()
  t.TapSecFunction(point)
}

func newToolTip(content fyne.CanvasObject, canvas fyne.Canvas, tap func (_ *fyne.PointEvent), tapSec func(_ *fyne.PointEvent)) *ToolTip {
  ret := &ToolTip{Content: content, Canvas: canvas, TapFunction: tap, TapSecFunction: tapSec}
  ret.ExtendBaseWidget(ret)
  return ret
}

// NewToolTip creates a new ToolTip for the specified content and displays it on the passed canvas.
func NewToolTip(content fyne.CanvasObject, canvas fyne.Canvas, tap func (_ *fyne.PointEvent), tapSec func(_ *fyne.PointEvent)) *ToolTip {
  return newToolTip(content, canvas, tap, tapSec)
}

// ShowToolTip creates a new ToolTip for the specified content and displays it on the passed canvas.
func ShowToolTip(content fyne.CanvasObject, canvas fyne.Canvas, tap func (_ *fyne.PointEvent), tapSec func(_ *fyne.PointEvent)) {
  newToolTip(content, canvas, tap, tapSec).Show()
}

type ToolTipBaseRenderer struct {
  ToolTip      *ToolTip
  background *canvas.Rectangle
}

func (r *ToolTipBaseRenderer) padding() fyne.Size {
  return fyne.NewSize(theme.Padding()*2, theme.Padding()*2)
}

func (r *ToolTipBaseRenderer) offset() fyne.Position {
  return fyne.NewPos(theme.Padding(), theme.Padding())
}

type ToolTipRenderer struct {
  //*widget.ShadowingRenderer
  ToolTipBaseRenderer
}

func (r *ToolTipRenderer) Destroy() {
}

func (r *ToolTipRenderer) Layout(_ fyne.Size) {
  innerSize := r.ToolTip.innerSize.Max(r.ToolTip.MinSize())
  r.ToolTip.Content.Resize(innerSize.Subtract(r.padding()))

  innerPos := r.ToolTip.innerPos
  if innerPos.X+innerSize.Width > r.ToolTip.Canvas.Size().Width {
    innerPos.X = r.ToolTip.Canvas.Size().Width - innerSize.Width
    if innerPos.X < 0 {
      innerPos.X = 0 // TODO here we may need a scroller as it's wider than our canvas
    }
  }
  if innerPos.Y+innerSize.Height > r.ToolTip.Canvas.Size().Height {
    innerPos.Y = r.ToolTip.Canvas.Size().Height - innerSize.Height
    if innerPos.Y < 0 {
      innerPos.Y = 0 // TODO here we may need a scroller as it's longer than our canvas
    }
  }
  r.ToolTip.Content.Move(innerPos.Add(r.offset()))

  r.background.Resize(innerSize)
  r.background.Move(innerPos)
  //r.LayoutShadow(innerSize, innerPos)
}

func (r *ToolTipRenderer) MinSize() fyne.Size {
  return r.ToolTip.Content.MinSize().Add(r.padding())
}

func (r *ToolTipRenderer) Refresh() {
  r.background.FillColor = theme.BackgroundColor()
  expectedContentSize := r.ToolTip.innerSize.Max(r.ToolTip.MinSize()).Subtract(r.padding())
  shouldRelayout := r.ToolTip.Content.Size() != expectedContentSize

  if r.background.Size() != r.ToolTip.innerSize || r.background.Position() != r.ToolTip.innerPos || shouldRelayout {
    r.Layout(r.ToolTip.Size())
  }
  if r.ToolTip.Canvas.Size() != r.ToolTip.BaseWidget.Size() {
    r.ToolTip.BaseWidget.Resize(r.ToolTip.Canvas.Size())
  }
  r.ToolTip.Content.Refresh()
  r.background.Refresh()
  //r.ShadowingRenderer.RefreshShadow()
}

func (r *ToolTipRenderer) Objects() []fyne.CanvasObject {
  background := canvas.NewRectangle(theme.BackgroundColor())
  return []fyne.CanvasObject{background, r.ToolTip.Content}
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (t *ToolTip) CreateRenderer() fyne.WidgetRenderer {
  t.ExtendBaseWidget(t)
  background := canvas.NewRectangle(theme.BackgroundColor())
  return &ToolTipRenderer{
    //widget.NewShadowingRenderer(objects, widget.ToolTipLevel),
    ToolTipBaseRenderer{ToolTip: t, background: background},
  }
}

