package customWidgets

import (
	"fyne.io/fyne/v2/widget"
)

const defaultPlaceHolder string = "(Select one or more items)"

type MultipleSelect struct {
	widget.Select
}

func NewMultipleSelect(options []string, changed func(string)) *MultipleSelect {
	s := &MultipleSelect{}
	s.OnChanged = changed
	s.Options = options
	s.PlaceHolder = defaultPlaceHolder
	s.ExtendBaseWidget(s)
	return s
}

func (s *MultipleSelect) updateSelected(text string) {
	s.Selected = text

	if s.OnChanged != nil {
		s.OnChanged(s.Selected)
	}

	s.Refresh()
}

// SelectedIndex returns the index value of the currently selected item in Options list.
// It will return -1 if there is no selection.
func (s *MultipleSelect) SelectedIndex() int {
	for i, option := range s.Options {
		if s.Selected == option {
			return i
		}
	}
	return -1 // not selected/found
}

// SetSelected sets the current option of the select widget
func (s *MultipleSelect) SetSelected(text string) {
	for _, option := range s.Options {
		if text == option {
			s.updateSelected(text)
		}
	}
}

// SetSelectedIndex will set the Selected option from the value in Options list at index position.
func (s *MultipleSelect) SetSelectedIndex(index int) {
	if index < 0 || index >= len(s.Options) {
		return
	}

	s.updateSelected(s.Options[index])
}
