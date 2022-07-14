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
}

func NewAddFormFrame(j *jsonutil.MyJson) *AddFormFrame {
	return &AddFormFrame{
		js: j,
	}

}

func (self *AddFormFrame) MakeFrame() tview.Primitive {
	add_form := tview.NewForm()

	add_form.SetBorder(true).SetTitle("Add Data")
	add_form.AddInputField("DM", "", 20, tview.InputFieldInteger, nil)
	add_form.AddInputField("Data(0x)", "", 20, nil, nil)

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
			return
		}

		fmt.Println(i_data)

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

	return add_form

}
