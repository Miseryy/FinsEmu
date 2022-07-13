package ui

import (
	udp "FinsEmu/UDP"
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type AddressFrame struct {
	AddressP       AddressViewPri
	frames         *Frames
	udp_soc        *udp.Udp_Sock
	write_log_call func(string)
}

type AddressViewPri struct {
	address_IF      *tview.InputField
	port_IF         *tview.InputField
	set_addr_button *tview.Button
}

func NewAddressFrame(f *Frames) *AddressFrame {
	return &AddressFrame{
		udp_soc: udp.New(),
		AddressP: AddressViewPri{
			address_IF:      tview.NewInputField(),
			port_IF:         tview.NewInputField(),
			set_addr_button: tview.NewButton("Set"),
		},
		frames: f,
	}
}

func (self *AddressFrame) MakeFrame() tview.Primitive {
	elements := []tview.Primitive{
		self.AddressP.address_IF,
		self.AddressP.port_IF,
		self.AddressP.set_addr_button,
	}

	self.AddressP.set_addr_button.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		self.frames.FocusEvent(elements, event)

		return event
	})

	self.AddressP.set_addr_button.SetSelectedFunc(func() {
		addr := self.AddressP.address_IF.GetText()
		port := self.AddressP.port_IF.GetText()
		p, _ := strconv.Atoi(port)
		self.SetAddress(addr, p)
		s := fmt.Sprintf("Set Address And Port\nAddress::%s\nPort::%s\n", addr, port)
		// self.frames.WriteLog(s)
		self.write_log_call(s)

	})

	field_width := 30

	self.AddressP.address_IF.
		SetLabel("Address:").
		SetFieldWidth(field_width).
		SetDoneFunc(func(key tcell.Key) {
			self.frames.Focus(elements, false)

		}).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			self.frames.FocusEvent(elements, event)

			return event
		})

	self.AddressP.port_IF.
		SetLabel("Port   :").
		SetFieldWidth(field_width).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			self.frames.Focus(elements, false)

		}).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			self.frames.FocusEvent(elements, event)

			return event
		})

	address_main := tview.NewFlex()
	address_form := tview.NewForm()

	address_form.AddInputField("Address", "", 40, nil, nil)
	address_form.AddInputField("Port", "", 40, tview.InputFieldInteger, nil)
	address_form.AddButton("Set", func() {
		addr := address_form.GetFormItem(0).(*tview.InputField).GetText()
		port := address_form.GetFormItem(1).(*tview.InputField).GetText()
		p, _ := strconv.Atoi(port)
		self.SetAddress(addr, p)
		s := fmt.Sprintf("Set Address And Port\nAddress::%s\nPort::%s\n", addr, port)
		// self.frames.WriteLog(s)
		self.write_log_call(s)

	})

	address_main.SetBorder(true).SetTitle("Address & Port <A>")

	address_main.SetDirection(tview.FlexRow)

	address_main.
		AddItem(address_form, 0, 1, true)

	return address_main
}

func (self *AddressFrame) SetAddress(addr string, port int) {
	self.udp_soc.SetAddr(addr, port)
}
