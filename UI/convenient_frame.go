package ui

import (
	"github.com/rivo/tview"
)

type ConvinientFrame struct {
	frames *Frames
}

func NewConvinientFrame(f *Frames) *ConvinientFrame {
	return &ConvinientFrame{
		frames: f,
	}
}

func (self *ConvinientFrame) MakeFrame() tview.Primitive {
	convinient_frame := tview.NewPages()
	self.frames.log_text_frame.SetBorder(true).SetTitle("Log")
	self.frames.log_text_frame.SetText("test")

	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Add Data")
	form.AddInputField("DM", "", 20, nil, nil)
	form.AddInputField("Data", "", 20, nil, nil)
	form.AddButton("Save", func() {
		// Write Json File
	})
	form.AddButton("Cancel", func() {
		form.Clear(true)
		convinient_frame.SwitchToPage("log")

	})

	convinient_frame.
		AddPage("Log", self.frames.log_text_frame, true, true).
		AddPage("Add", form, true, false)

	return convinient_frame
}
