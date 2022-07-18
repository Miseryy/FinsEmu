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
	delete_frame     *DeleteFormFrame
	convinient_frame *tview.Pages
	add_form         *tview.Grid
	delete_table     *tview.Table
	log_text         string
}

func NewConvinientFrame(f *Frames) *ConvinientFrame {
	return &ConvinientFrame{
		frames: f,
	}
}

type PageName string

const (
	Log = string("Log")
	Add = string("Add")
	Del = string("Del")
)

func (self *ConvinientFrame) MakeFrame() tview.Primitive {
	js := jsonutil.New(json_path)
	self.convinient_frame = self.frames.MakePageFrame("convinient", tview.NewPages())
	self.log_text_frame = NewLogTextViewFrame()
	self.add_frame = NewAddFormFrame(self.frames, js)
	self.delete_frame = NewDeleteFormFrame(js)

	self.add_form = self.add_frame.MakeFrame().(*tview.Grid)
	log_text_view := self.log_text_frame.MakeFrame().(*tview.TextView)
	self.delete_table = self.delete_frame.MakeFrame().(*tview.Table)

	err := js.LoadJson()

	if err != nil {
		self.log_text_frame.WriteLog("Json Load Failed", true)
		s := fmt.Sprint(err)
		self.log_text_frame.WriteLog(s, true)
	}

	self.add_frame.Change2LogFrameCall = func() {
		self.Change2LogFrame()
	}

	self.add_frame.WriteLog = func(text string) {
		self.log_text_frame.WriteLog(text, true)
	}

	self.delete_frame.change2LogFrame_call = func() {
		self.Change2LogFrame()
	}

	self.convinient_frame.
		AddPage(Log, log_text_view, true, true).
		AddPage(Add, self.add_form, true, false).
		AddPage(Del, self.delete_table, true, false)

	return self.convinient_frame
}

func (self *ConvinientFrame) Change2LogFrame() {
	self.convinient_frame.SwitchToPage(Log)

}

func (self *ConvinientFrame) Change2AddDataFrame() {
	self.convinient_frame.SwitchToPage(Add)
	self.frames.App.SetFocus(self.add_form)
}

func (self *ConvinientFrame) Change2DeleteFrame() {
	self.convinient_frame.SwitchToPage(Del)
	self.delete_frame.makeCells(self.delete_table)
	self.frames.App.SetFocus(self.delete_table)

}
