package ui

import (
	jsonutil "FinsEmu/JsonUtil"
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

type AddFormFrame struct {
	Change2LogFrameCall func()
	WriteLog            func(string)
	js                  *jsonutil.MyJson
	frames              *Frames
}

func NewAddFormFrame(f *Frames, j *jsonutil.MyJson) *AddFormFrame {
	return &AddFormFrame{
		js:     j,
		frames: f,
	}

}

func (self *AddFormFrame) checkHex(text string, ch rune) bool {
	if len(text) > 4 {
		return false
	}

	if ch >= '0' && ch <= '9' {
		return true
	}

	if (ch >= 'a' && ch <= 'f') ||
		(ch >= 'A' && ch <= 'F') {
		return true
	}

	return false
}

func (self *AddFormFrame) MakeFrame() tview.Primitive {
	add_grid := tview.NewGrid()
	add_form := tview.NewForm()

	add_grid.SetBorder(true).SetTitle("Add Data")
	add_form.AddInputField("DM", "", 20, tview.InputFieldInteger, nil)
	add_form.AddInputField("Data(Hex)", "", 20, self.checkHex, nil)

	add_form.AddButton("Save", func() {
		dm_no := add_form.GetFormItem(0).(*tview.InputField).GetText()
		data := add_form.GetFormItem(1).(*tview.InputField).GetText()

		if !(len(dm_no) > 0) || !(len(data) > 0) {
			self.Change2LogFrameCall()
			self.WriteLog("Failed Add Data")
			add_form.GetFormItem(0).(*tview.InputField).SetText("")
			add_form.GetFormItem(1).(*tview.InputField).SetText("")

			return
		}

		i_data, err := strconv.ParseInt(data, 16, 64)

		if err != nil {
			self.WriteLog(err.Error())
			add_form.GetFormItem(0).(*tview.InputField).SetText("")
			add_form.GetFormItem(1).(*tview.InputField).SetText("")
			self.Change2LogFrameCall()
			return
		}

		self.js.AddItem(dm_no, i_data)
		err = self.js.WriteJson()

		if err != nil {
			s := fmt.Sprint(err)
			self.WriteLog(s)
			return
		}

		add_form.GetFormItem(0).(*tview.InputField).SetText("")
		add_form.GetFormItem(1).(*tview.InputField).SetText("")
		self.WriteLog("Added\n")
		self.Change2LogFrameCall()
	})

	add_form.AddButton("Cancel", func() {
		add_form.GetFormItem(0).(*tview.InputField).SetText("")
		add_form.GetFormItem(1).(*tview.InputField).SetText("")
		self.Change2LogFrameCall()

	}).AddButton("Clear", func() {
		add_form.GetFormItem(0).(*tview.InputField).SetText("")
		add_form.GetFormItem(1).(*tview.InputField).SetText("")

	})

	add_grid.SetRows(1, 0).SetColumns(0)

	add_grid.AddItem(add_form, 1, 0, 1, 1, 0, 0, false)

	return add_grid

}
