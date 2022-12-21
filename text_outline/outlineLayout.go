package text_outline

import (
	"fyne.io/fyne/v2"
)

type outlineLayout struct {
	thickness float32
}

func (o *outlineLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := float32(0), float32(0)
	for _, element := range objects {
		childSize := element.MinSize()

		w += childSize.Width + o.thickness
		if childSize.Height > h {
			h = childSize.Height + o.thickness
		}
	}
	return fyne.NewSize(w, h)
}

func (o *outlineLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(o.thickness/2, 0)
	for _, element := range objects {
		size := element.MinSize()
		element.Resize(size)
		element.Move(pos)

		pos = pos.Add(fyne.NewPos(size.Width+o.thickness, 0))
	}
}
