package ui

import (
	jsonutil "FinsEmu/JsonUtil"
	"fmt"
	"strconv"

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
	self.add_form.AddInputField("DM", "", 20, tview.InputFieldInteger, nil)
	self.add_form.AddInputField("Data(0x)", "", 20, nil, nil)

	self.convinient_frame.
		AddPage("Log", self.log_text_view, true, true).
		AddPage("Add", self.add_form, true, false)

	js := jsonutil.New()
	err := js.LoadJson(json_path)

	if err != nil {
		self.WriteLog("Json Load Failed")
		s := fmt.Sprint(err)
		self.WriteLog(s)

	}

	self.add_form.AddButton("Save", func() {
		dm_no := self.add_form.GetFormItem(0).(*tview.InputField).GetText()
		data := self.add_form.GetFormItem(1).(*tview.InputField).GetText()

		if !(len(dm_no) > 0) || !(len(data) > 0) {
			self.Change2LogFrame()
			self.WriteLog("Failed Add Data")
			self.add_form.GetFormItem(0).(*tview.InputField).SetText("")
			self.add_form.GetFormItem(1).(*tview.InputField).SetText("")

			return
		}

		i_data, err := strconv.ParseInt(data, 16, 64)

		if err != nil {
			self.WriteLog(err.Error())
			self.add_form.GetFormItem(0).(*tview.InputField).SetText("")
			self.add_form.GetFormItem(1).(*tview.InputField).SetText("")
			return
		}

		js.AddItem(dm_no, i_data)
		err = js.WriteJson(json_path)

		if err != nil {
			self.WriteLog(err.Error())
			return
		}

		self.add_form.GetFormItem(0).(*tview.InputField).SetText("")
		self.add_form.GetFormItem(1).(*tview.InputField).SetText("")
		self.WriteLog("Added\n")
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
	self.log_text += text + "\n"
	self.log_text_view.SetText(self.log_text)

}
