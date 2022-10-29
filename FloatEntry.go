package customWidgets

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type FloatEntry struct {
	widget.Entry
}

func NewFloatEntry() *FloatEntry {
	entry := &FloatEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *FloatEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' {
		e.Entry.TypedRune(r)
	}
}

func (e *FloatEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *FloatEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func (e *FloatEntry) Value() float64 {
	val, _ := strconv.ParseFloat(e.Entry.Text, 64)
	return val
}

func (e *FloatEntry) SetValue(val float64) {
	e.Entry.SetText(strconv.FormatFloat(val, 'f', 2, 64))
}
