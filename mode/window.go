package mode

import (
	cmd "github.com/satran/e/commands"
	"github.com/satran/e/editor"
	"github.com/nsf/termbox-go"
)

type WindowMode struct {
	editor *editor.Editor
	count  int
}

func NewWindowMode(editor *editor.Editor, count int) WindowMode {
	return WindowMode{editor: editor, count: count}
}

func (m WindowMode) Enter(e *editor.Editor) {
}

func (m WindowMode) OnKey(ev *termbox.Event) {
	switch ev.Ch {
	case 'h':
		m.editor.Commands <- cmd.NearestVSplit{cmd.Backward}
	case 'j':
		m.editor.Commands <- cmd.NearestHSplit{cmd.Forward}
	case 'k':
		m.editor.Commands <- cmd.NearestHSplit{cmd.Backward}
	case 'l':
		m.editor.Commands <- cmd.NearestVSplit{cmd.Forward}
	case '=':
		// TODO viewTree.normalizeSplit
	}
	m.editor.SetMode(NewNormalMode(m.editor))
}

func (m WindowMode) Exit() {
}
