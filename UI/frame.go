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
	App       *tview.Application
	Udp       *udp.Udp_Sock
	Connected bool
	frame_map map[string]tview.Primitive
}

func (self *Frames) MakePageFrame(name string, frame *tview.Pages) *tview.Pages {
	self.frame_map[name] = frame
	return frame
}

func (self *Frames) FrameRegister(name string, frame *tview.Primitive) *tview.Primitive {
	self.frame_map[name] = *frame
	return frame
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
