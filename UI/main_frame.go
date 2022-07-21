package ui

import (
	fins "FinsEmu/Fins"
	jsonutil "FinsEmu/JsonUtil"
	udp "FinsEmu/UDP"
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type MainFrame struct {
	main_frame   *tview.Grid
	frames       *Frames
	child_frames *ChildFrames
	address_json *jsonutil.MyJson
}

type ChildFrames struct {
	address_frame    *AddressFrame
	command_frame    *CommandFrame
	convinient_frame *ConvinientFrame
}

type ChildPrimitive struct {
	address_flex    *tview.Flex
	command_page    *tview.Pages
	convenient_page *tview.Pages
}

func NewMainFrame(f *Frames) *MainFrame {
	return &MainFrame{
		main_frame: tview.NewGrid(),
		frames:     f,
	}
}

func (self *MainFrame) setCallBacks() {
	self.child_frames.command_frame.change_add_form_callback = func() {
		self.child_frames.convinient_frame.Change2AddDataFrame()
	}

	self.child_frames.command_frame.reset_log_callback = func() {
		self.child_frames.convinient_frame.log_text_frame.ResetLog()
	}

	self.child_frames.command_frame.connect_udp_callback = func() {
		addr := self.child_frames.address_frame.address_form.GetFormItem(0).(*tview.InputField).GetText()
		port := self.child_frames.address_frame.address_form.GetFormItem(1).(*tview.InputField).GetText()

		if self.frames.Connected {
			s := fmt.Sprintf("!!Connected \nAddress::%s\nPort   ::%s", addr, port)
			self.WriteLog(s, true)
		}

		int_port, _ := strconv.Atoi(port)

		self.frames.Udp.SetAddr(addr, int_port)
		err := self.frames.Udp.Listen()

		if err != nil {
			self.WriteLog(err.Error(), true)
			return
		}

		self.address_json.
			AddItemString("address", addr).
			AddItemInt("port", int64(int_port)).
			WriteJson()

		self.child_frames.convinient_frame.Change2LogFrame()

		s := fmt.Sprintf("Connect \nAddress::%s\nPort   ::%s\n\n", addr, port)
		self.WriteLog(s, true)
		self.frames.Connected = true

		update_draw := func(text string) {
			self.frames.App.QueueUpdateDraw(func() {
				self.WriteLog(text, true)
			})
		}

		go func() {
			for {
				if !self.frames.Connected {
					break
				}

				recv_buff, addr, err := fins.RecvHostData(self.frames.Udp)

				if err != nil {
					update_draw(err.Error())
					continue
				}

				str_port := strconv.Itoa(addr.Port)
				s := fmt.Sprintf("Recv [%s:%s]:%X", addr.IP, str_port, recv_buff)
				update_draw(s)

				recv_param, err := fins.CheckFinsCommand(recv_buff)

				if err != nil {
					update_draw(err.Error())
					continue
				}

				send_buff, err := fins.MakeSendCommand(self.frames.Udp, recv_param, addr.IP.String(), data_json_path)

				if err != nil {
					update_draw(err.Error())
					continue
				}

				_, err = self.frames.Udp.WriteTo(send_buff, addr)

				if err != nil {
					update_draw(err.Error())
					continue
				}

				t := fmt.Sprintf("Send [%s:%d]:%X", addr.IP, addr.Port, send_buff)
				update_draw(t)

				// self.frames.App.QueueUpdateDraw(func() {
				// 	str_port := strconv.Itoa(addr.Port)
				// 	s := fmt.Sprintf("F [%s:%s]:%X", addr.IP, str_port, recv_buff)
				// 	update_draw(s)

				// 	// self.frames.Udp.WriteTo([]byte(buff[:num]), addr)
				// 	// s = fmt.Sprintf("Send [%s:%s]:%s", addr.IP, str_port, string(buff[:num]))
				// 	// self.child_frames.convinient_frame.log_text_frame.WriteLog(s, true)
				// })
			}
		}()
	}

	self.child_frames.command_frame.close_udp_callback = func() {
		err := self.frames.Udp.Close()
		if err != nil {
			self.child_frames.convinient_frame.log_text_frame.WriteLog(err.Error(), true)
			self.frames.Connected = false
			return
		}
		self.frames.Connected = false

	}

	self.child_frames.command_frame.change_delete_form_callback = func() {
		self.child_frames.convinient_frame.Change2DeleteFrame()
	}

	self.child_frames.address_frame.write_log_call = func(text string) {
		self.child_frames.convinient_frame.log_text_frame.WriteLog(text, true)
	}

}

func (self *MainFrame) MakeFrame() tview.Primitive {
	self.frames.Udp = udp.New()
	self.frames.Connected = false
	self.address_json = jsonutil.New(setting_json_path)

	self.child_frames = &ChildFrames{
		convinient_frame: NewConvinientFrame(self.frames),
		command_frame:    NewCommandFrame(self.frames),
		address_frame:    NewAddressFrame(self.address_json, self.frames),
	}

	child_frames := ChildPrimitive{
		convenient_page: self.child_frames.convinient_frame.MakeFrame().(*tview.Pages),
		command_page:    self.child_frames.command_frame.MakeFrame().(*tview.Pages),
		address_flex:    self.child_frames.address_frame.MakeFrame().(*tview.Flex),
	}

	self.setCallBacks()

	self.main_frame.SetBorder(true).SetTitle("FinsUDPEmurator").SetTitleAlign(0).SetTitleColor(tcell.ColorYellowGreen)
	self.main_frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			self.frames.App.SetFocus(child_frames.address_flex)
			self.child_frames.convinient_frame.Change2LogFrame()

		case tcell.KeyCtrlS:
			self.frames.App.SetFocus(child_frames.command_page)
			self.child_frames.convinient_frame.Change2LogFrame()

		case tcell.KeyCtrlL:
			self.child_frames.convinient_frame.Change2LogFrame()
			self.frames.App.SetFocus(self.frames.frame_map[LogTextFrameName])
		}

		return event
	})

	self.main_frame.SetRows(7, 0, 8).SetColumns(45, 0)
	// self.main_frame.SetBackgroundColor(tcell.ColorWhite)
	self.main_frame.AddItem(child_frames.address_flex, 0, 0, 1, 1, 0, 0, true)
	self.main_frame.AddItem(child_frames.command_page, 1, 0, 3, 1, 0, 0, true)
	self.main_frame.AddItem(child_frames.convenient_page, 0, 1, 4, 1, 0, 0, true)

	return self.main_frame
}

func (self *MainFrame) WriteLog(text string, new_line bool) {
	self.child_frames.convinient_frame.log_text_frame.WriteLog(text, new_line)
}
