package customWidgets

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type floatEntry struct {
	widget.Entry
}

func newFloatEntry() *floatEntry {
	entry := &floatEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *floatEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' {
		e.Entry.TypedRune(r)
	}
}

func (e *floatEntry) TypedShortcut(shortcut fyne.Shortcut) {
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

func (e *floatEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
