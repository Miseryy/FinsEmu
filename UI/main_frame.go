package ui

import (
	udp "FinsEmu/UDP"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainFrame struct {
	main_frame *tview.Grid
	frames     *Frames
}

type ChildFrames struct {
	address_frame    *tview.Flex
	command_frame    *tview.Pages
	convenient_frame *tview.Pages
}

func NewMainFrame(f *Frames) *MainFrame {
	return &MainFrame{
		main_frame: tview.NewGrid(),
		frames:     f,
	}
}

func (self *MainFrame) setComCB_FromConvinient(command_frame *CommandFrame, convinient_frame *ConvinientFrame) {
	command_frame.change_add_form_callback = func() {
		convinient_frame.Change2AddDataFrame()
	}

	command_frame.reset_log_callback = func() {
		convinient_frame.log_text_frame.ResetLog()
	}

	command_frame.connect_udp_callback = func() {
		err := self.frames.Udp.Listen()
		if err != nil {
			convinient_frame.log_text_frame.WriteLog(err.Error())
			return
		}
		self.frames.Connected = true

	}

	command_frame.close_udp_callback = func() {
		err := self.frames.Udp.Close()
		if err != nil {
			convinient_frame.log_text_frame.WriteLog(err.Error())
			return
		}
		self.frames.Connected = false

	}

	command_frame.change_delete_form_callback = func() {
		convinient_frame.Change2DeleteFrame()

	}

}

func (self *MainFrame) MakeFrame() tview.Primitive {
	self.frames.Udp = udp.New()
	self.frames.Connected = false

	convinient_frame := NewConvinientFrame(self.frames)
	command_frame := NewCommandFrame(self.frames)
	address_frame := NewAddressFrame(self.frames)

	child_frames := ChildFrames{
		convenient_frame: convinient_frame.MakeFrame().(*tview.Pages),
		command_frame:    command_frame.MakeFrame().(*tview.Pages),
		address_frame:    address_frame.MakeFrame().(*tview.Flex),
	}

	address_frame.write_log_call = func(text string) {
		convinient_frame.log_text_frame.WriteLog(text)
	}

	self.setComCB_FromConvinient(command_frame, convinient_frame)

	self.main_frame.SetBorder(true).SetTitle("FinsUDPEmurator").SetTitleAlign(0).SetTitleColor(tcell.ColorYellowGreen)
	self.main_frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			self.frames.App.SetFocus(child_frames.address_frame)

		case tcell.KeyCtrlS:
			self.frames.App.SetFocus(child_frames.command_frame)

		case tcell.KeyCtrlL:
		}

		return event
	})

	self.main_frame.SetRows(9, 0, 8).SetColumns(45, 0)
	// self.main_frame.SetBackgroundColor(tcell.ColorWhite)
	self.main_frame.AddItem(child_frames.address_frame, 0, 0, 1, 1, 0, 0, true)
	self.main_frame.AddItem(child_frames.command_frame, 1, 0, 3, 1, 0, 0, true)
	self.main_frame.AddItem(child_frames.convenient_frame, 0, 1, 4, 1, 0, 0, true)

	return self.main_frame
}
