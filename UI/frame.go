package ui

import (
	jsonutil "FinsEmu/JsonUtil"
	udp "FinsEmu/UDP"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type FrameName string

const (
	AddressFrameName     FrameName = "Address_F"
	CommandFrameName               = "Comm_F"
	CommandListFrameName           = "CommList_F"

	ConvenientFrameName = "Conv_F"

	// LogFrame
	LogTextFrameName = "LogText_F"

	// AddFrame
	AddDataGridFrameName = "AddDataGrid_F"
	AddDataFormFrameName = "AddDataForm_F"

	// DeleteFrame
	DeleteTableFrameName = "DeleteTable_F"
)

type Frame interface {
	MakeFrame() tview.Primitive
}

type Frames struct {
	App          *tview.Application
	Udp          *udp.Udp_Sock
	Connected    bool
	frame_map    map[FrameName]tview.Primitive
	setting_json *jsonutil.MyJson
}

func (self *Frames) FrameRegister(name FrameName, frame tview.Primitive) tview.Primitive {
	if self.frame_map == nil {
		self.frame_map = make(map[FrameName]tview.Primitive)
	}

	self.frame_map[name] = frame
	return self.frame_map[name]
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
