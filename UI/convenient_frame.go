package ui

import (
	jsonutil "FinsEmu/JsonUtil"
	"fmt"

	"github.com/rivo/tview"
)

type ConvinientFrame struct {
	frames           *Frames
	log_text_frame   *LogTextViewFrame
	add_frame        *AddFormFrame
	convinient_frame *tview.Pages
	add_form         *tview.Form
	log_text         string
}

func NewConvinientFrame(f *Frames) *ConvinientFrame {
	return &ConvinientFrame{
		frames: f,
	}
}

func (self *ConvinientFrame) MakeFrame() tview.Primitive {
	self.convinient_frame = tview.NewPages()
	self.log_text_frame = NewLogTextViewFrame()
	log_text_view := self.log_text_frame.MakeFrame().(*tview.TextView)

	js := jsonutil.New()
	err := js.LoadJson(json_path)

	if err != nil {
		self.log_text_frame.WriteLog("Json Load Failed")
		s := fmt.Sprint(err)
		self.log_text_frame.WriteLog(s)
	}

	self.add_frame = NewAddFormFrame(js)
	self.add_form = self.add_frame.MakeFrame().(*tview.Form)

	self.add_frame.Change2LogFrameCall = func() {
		self.Change2LogFrame()
	}

	self.add_frame.WriteLog = func(text string) {
		self.log_text_frame.WriteLog("Add")
	}

	self.convinient_frame.
		AddPage("Log", log_text_view, true, true).
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
