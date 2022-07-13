package ui

import (
	"github.com/rivo/tview"
)

type ConvinientFrame struct {
	frames           *Frames
	log_text_view    *tview.TextView
	add_form         *tview.Form
	convinient_frame *tview.Pages
	log_text         string
}

func NewConvinientFrame(f *Frames) *ConvinientFrame {
	return &ConvinientFrame{
		frames:        f,
		log_text_view: tview.NewTextView(),
		add_form:      tview.NewForm(),
	}
}

func (self *ConvinientFrame) MakeFrame() tview.Primitive {
	self.convinient_frame = tview.NewPages()
	self.log_text_view.SetBorder(true).SetTitle("Log")
	self.log_text_view.SetText("test")

	self.add_form.SetBorder(true).SetTitle("Add Data")
	self.add_form.AddInputField("DM", "", 20, nil, nil)
	self.add_form.AddInputField("Data", "", 20, nil, nil)
	self.add_form.AddButton("Save", func() {
		// Write Json File

		self.add_form.Clear(true)
		self.Change2LogFrame()
	})

	self.add_form.AddButton("Cancel", func() {
		self.add_form.GetFormItem(0).(*tview.InputField).SetText("")
		self.add_form.GetFormItem(1).(*tview.InputField).SetText("")
		self.Change2LogFrame()

	}).AddButton("Clear", func() {
		self.add_form.GetFormItem(0).(*tview.InputField).SetText("")
		self.add_form.GetFormItem(1).(*tview.InputField).SetText("")

	})

	self.convinient_frame.
		AddPage("Log", self.log_text_view, true, true).
		AddPage("Add", self.add_form, true, false)

	return self.convinient_frame
}

func (self *ConvinientFrame) Change2LogFrame() {
	self.convinient_frame.SwitchToPage("Log")

}

func (self *ConvinientFrame) Change2AddDataFrame() {
	self.convinient_frame.SwitchToPage("Add")
	self.frames.App.SetFocus(self.add_form)
}

func (self *ConvinientFrame) ResetLog() {
	self.log_text_view.SetText("")
}

func (self *ConvinientFrame) WriteLog(text string) {
	self.log_text += text
	self.log_text_view.SetText(self.log_text)

}
