package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type CommandFrame struct {
	Frames *Frames
}

func NewCommandFrame(f *Frames) *CommandFrame {
	return &CommandFrame{
		Frames: f,
	}
}

func (self *CommandFrame) MakeFrame() tview.Primitive {
	command_main := tview.NewPages()
	command_main.SetBorder(true).SetTitle("C<O>mmand")
	reset_button := tview.NewButton("ResetLog")
	reset_button.SetSelectedFunc(func() {
		self.Frames.ResetLog()
	})

	main_util := tview.NewGrid()
	main_util.SetRows(0, 0, 0, 0).SetColumns(0, 0, 0)
	main_util.AddItem(reset_button, 4, 2, 1, 1, 0, 0, true)

	g2 := tview.NewGrid()
	g2.SetRows(0, 0, 0, 0).SetColumns(0, 0, 0)
	g2.AddItem(reset_button, 3, 1, 1, 1, 0, 0, true)
	command_main.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'A':
				command_main.SwitchToPage("Main")
			case 'C':
				command_main.SwitchToPage("Main2")
			}
		}

		return event
	})

	command_main.
		AddPage("Main", main_util, true, true).
		AddPage("Main2", g2, true, false)

	return command_main
}
