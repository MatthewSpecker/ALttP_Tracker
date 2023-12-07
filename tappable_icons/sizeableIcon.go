package tappable_icons

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type SizeableIcon struct {
	widget.Icon
	resources      []fyne.Resource
	size 	       float32
}

func NewSizeableIcon(res []fyne.Resource, size float32) (*SizeableIcon, error) {
	if len(res) <= 1 {
		return nil, errors.New("'res' must contain 2 or more resources")
	}
	if size <= 0 {
		return nil, errors.New("'size' must be greater than 0")
	}

	icon := &SizeableIcon{
		resources:      res,
		size:        size,
	}

	icon.ExtendBaseWidget(icon)
	icon.SetResource(icon.resources[0])

	return icon, nil
}

func (s *SizeableIcon) Layout() *fyne.Container {
	tapIconContainer := container.New(layout.NewCenterLayout(), s)
	return tapIconContainer
}

func (s *SizeableIcon) Update(current int) {
	s.SetResource(s.resources[current])
}

func (s *SizeableIcon) MinSize() fyne.Size {
	return fyne.NewSize(theme.Padding()*s.size/4, theme.Padding()*s.size/4)
}
