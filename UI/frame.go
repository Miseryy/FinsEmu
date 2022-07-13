package ui

import (
	udp "FinsEmu/UDP"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Frame interface {
	MakeFrame() tview.Primitive
}

type Frames struct {
	App *tview.Application
	Udp *udp.Udp_Sock
	Connected bool
}

func (self *Frames) Focus(elements []tview.Primitive, reverse bool) {
	max := len(elements)
	for i, el := range elements {
		if !el.HasFocus() {
			continue
		}

		if reverse {
			i = i - 1
			if i < 0 {
				i = max - 1
			}
		} else {
			i = i + 1
			i = i % max
		}

		self.App.SetFocus(elements[i])
		return
	}

	self.App.SetFocus(elements[0])

}

func (self *Frames) FocusEvent(elements []tview.Primitive, event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyCtrlN:
		self.Focus(elements, false)
	case tcell.KeyCtrlP:
		self.Focus(elements, true)

	}
}

// func (self *Frames) WriteLog(text string) {
// 	self.log_text += text
// 	self.log_text_frame.SetText(self.log_text)
//
// }
//
// func (self *Frames) ResetLog() {
// 	self.log_text = ""
// 	self.log_text_frame.SetText(self.log_text)
//
// }

// func (self *Frames) ConvinientFrameChangePage(name string) {
// 	self.convinient_frame.SwitchToPage(name)
// 	self.App.SetFocus(self.c)
// }
