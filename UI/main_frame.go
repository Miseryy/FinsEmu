package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainFrame struct {
	main_frame *tview.Grid
	frames     *Frames
}

type ChildFrames struct {
	address_frame    *tview.Grid
	command_frame    *tview.Pages
	convenient_frame *tview.Pages
}

func NewMainFrame(f *Frames) *MainFrame {
	return &MainFrame{
		main_frame: tview.NewGrid(),
		frames:     f,
	}
}

func (self *MainFrame) MakeFrame() tview.Primitive {
	self.frames.log_text_frame = tview.NewTextView()
	self.frames.log_text = string("")

	child_frames := ChildFrames{
		address_frame:    NewAddressFrame(self.frames).MakeFrame().(*tview.Grid),
		command_frame:    NewCommandFrame(self.frames).MakeFrame().(*tview.Pages),
		convenient_frame: NewConvinientFrame(self.frames).MakeFrame().(*tview.Pages),
	}
	self.main_frame.SetBorder(true).SetTitle("FinsUDPEmurator").SetTitleAlign(0).SetTitleColor(tcell.ColorYellowGreen)
	self.main_frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			self.frames.App.SetFocus(child_frames.address_frame)

		case tcell.KeyCtrlO:
			self.frames.App.SetFocus(child_frames.command_frame)

		case tcell.KeyCtrlL:
		}

		return event
	})

	self.main_frame.SetRows(7, 0, 8).SetColumns(44, 0)
	// self.main_frame.SetBackgroundColor(tcell.ColorWhite)
	self.main_frame.AddItem(child_frames.address_frame, 0, 0, 1, 1, 0, 0, true)
	self.main_frame.AddItem(child_frames.command_frame, 1, 0, 3, 1, 0, 0, true)
	self.main_frame.AddItem(child_frames.convenient_frame, 0, 1, 4, 1, 0, 0, true)

	return self.main_frame
}
