package ui

import (
	jsonutil "FinsEmu/JsonUtil"
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type AddressFrame struct {
	AddressP       AddressViewPri
	frames         *Frames
	address_form   *tview.Form
	write_log_call func(string)
	address        string
	port           string
	js             *jsonutil.MyJson
}

type AddressViewPri struct {
	address_IF      *tview.InputField
	port_IF         *tview.InputField
	set_addr_button *tview.Button
}

func NewAddressFrame(j *jsonutil.MyJson, f *Frames) *AddressFrame {
	return &AddressFrame{
		AddressP: AddressViewPri{
			address_IF:      tview.NewInputField(),
			port_IF:         tview.NewInputField(),
			set_addr_button: tview.NewButton("Set"),
		},
		frames: f,
		js:     j,
	}
}

func (self *AddressFrame) checkAddressInputs(text string, ch rune) bool {
	if ch >= '0' && ch <= '9' ||
		ch == '.' {
		return true
	}

	return false
}

func (self *AddressFrame) MakeFrame() tview.Primitive {
	e := self.js.LoadJson()
	if e != nil {
		fmt.Println(e)
	}

	json_map := self.js.GetMap()

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
	self.address_form = tview.NewForm()

	self.address_form.AddInputField("Address", "", 40, self.checkAddressInputs, nil)
	self.address_form.AddInputField("Port", "", 40, tview.InputFieldInteger, nil)

	address_main.SetBorder(true).SetTitle("Address & Port <A>")

	address_main.SetDirection(tview.FlexRow)

	address, ad_ok := json_map["address"].(string)
	port, po_ok := json_map["port"].(float64)
	if ad_ok && po_ok {
		self.address_form.GetFormItem(0).(*tview.InputField).SetText(address)
		self.address_form.GetFormItem(1).(*tview.InputField).SetText(strconv.Itoa(int(port)))
	}

	address_main.
		AddItem(self.address_form, 0, 1, true)

	return address_main
}

func (self *AddressFrame) SetAddress(addr string, port int) {
	self.js.AddItemString("address", addr).
		AddItemInt("port", int64(port)).
		WriteJson()
	self.frames.Udp.SetAddr(addr, port)
}
