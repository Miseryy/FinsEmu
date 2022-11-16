package ui

import "github.com/rivo/tview"

type LogTextViewFrame struct {
	text_view *tview.TextView
	log_text  string
}

func NewLogTextViewFrame() *LogTextViewFrame {
	return &LogTextViewFrame{}
}

func (self *LogTextViewFrame) MakeFrame() tview.Primitive {
	self.text_view = tview.NewTextView()
	self.text_view.SetDynamicColors(true)
	self.text_view.SetBorder(true).SetTitle("Log")

	return self.text_view
}

func (self *LogTextViewFrame) WriteLog(text string, new_line bool) {
	if new_line {
		self.log_text += text + "\n"
	} else {
		self.log_text += text
	}

	self.text_view.SetText(self.log_text)

}

func (self *LogTextViewFrame) ResetLog() {
	self.log_text = ""
	self.text_view.SetText("")
}
