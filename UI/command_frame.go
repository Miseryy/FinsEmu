package ui

import (
	"github.com/rivo/tview"
)

type CommandFrame struct {
	frames                      *Frames
	change_add_form_callback    func()
	change_delete_form_callback func()
	reset_log_callback          func()
	connect_udp_callback        func()
	close_udp_callback          func()
}

func NewCommandFrame(f *Frames) *CommandFrame {
	return &CommandFrame{
		frames: f,
	}
}

func (self *CommandFrame) MakeFrame() tview.Primitive {
	command_main := tview.NewPages()
	command_main.SetBorder(true).SetTitle("Commands <S>")

	main_commandlist_frame := tview.NewFlex().SetDirection(tview.FlexRow)
	command_list := tview.NewList()
	command_list.
		AddItem("Add Data", "Add DM Data", 'a', func() {
			self.change_add_form_callback()
		}).
		AddItem("Delete Data", "Delete DM Data", 'd', func() {
			self.change_delete_form_callback()

		}).
		AddItem("Connect", "Connect UDP", 'c', func() {
			self.connect_udp_callback()

		}).
		AddItem("Connect Close", "Close UDP", 'e', func() {
			self.close_udp_callback()

		}).
		AddItem("ResetLog", "", 'r', func() {
			self.reset_log_callback()

		}).
		AddItem("Quit", "", 'q', func() {
			self.frames.App.Stop()
		})

	main_commandlist_frame.
		AddItem(command_list, 0, 1, true)

	command_main.
		AddPage("Main", main_commandlist_frame, true, true)

	return command_main
}
