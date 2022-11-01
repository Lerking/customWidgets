package customWidgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MultipleSelect struct {
	widget.Select
	Selected []string
	Options  []*widget.Check
}

func NewMultipleSelect() *MultipleSelect {
	selct := &MultipleSelect{}
	selct.ExtendBaseWidget(selct)
	return selct
}

func (s *MultipleSelect) SetOptions(opts []string) []*widget.Check {
	var ch *widget.Check
	for _, opt := range opts {
		ch = widget.NewCheck(opt, func(b bool) {
			s.SetSelected(opt)
		})
		s.Options = append(s.Options, ch)
	}
	return s.Options
}

func (s *MultipleSelect) SetSelected(opt string) {
	for i, v := range s.Selected {
		if v == opt {
			s.Selected = append(s.Selected[:i], s.Selected[i+1:]...)
			return
		}
	}
	s.Selected = append(s.Selected, opt)
}

func (s *MultipleSelect) GetSelected() []string {
	return s.Selected
}

func (s *MultipleSelect) GetOptions() []*widget.Check {
	return s.Options
}

func (s *MultipleSelect) ClearSelected() {
	s.Selected = nil
}

func (s *MultipleSelect) ClearOptions() {
	s.Options = nil
}

func (s *MultipleSelect) ClearAll() {
	s.ClearSelected()
	s.ClearOptions()
}

func (s *MultipleSelect) Refresh() {
	s.Select.Refresh()
}

func (s *MultipleSelect) Hide() {
	s.Select.Hide()
}

func (s *MultipleSelect) Show() {
	s.Select.Show()
}

func (s *MultipleSelect) Resize(size fyne.Size) {
	s.Select.Resize(size)
}

func (s *MultipleSelect) Move(pos fyne.Position) {
	s.Select.Move(pos)
}

func (s *MultipleSelect) MinSize() fyne.Size {
	return s.Select.MinSize()
}

func (s *MultipleSelect) Tapped(_ *fyne.PointEvent) {
	s.Select.Tapped(nil)
}

func (s *MultipleSelect) SetPlaceHolder() {
	txt := ""
	for i, opt := range s.Options {
		if i == len(s.Options)-1 && opt.Checked {
			txt += opt.Text
			s.Select.PlaceHolder = txt
			return
		} else if i < len(s.Options)-1 && opt.Checked {
			txt += opt.Text + ", "
		} else if i == len(s.Options)-1 && !opt.Checked {
			s.Select.PlaceHolder = txt
			return
		}
	}
}
