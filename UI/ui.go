package ui

import (
	udp "FinsEmu/UDP"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type AddressViewPrimitive struct {
	address_IF      *tview.InputField
	port_IF         *tview.InputField
	set_addr_button *tview.Button
}

type EmuUI struct {
	app *tview.Application

	main_frame   *tview.Grid
	head         *tview.Grid
	left         *tview.Pages
	address_view *tview.Grid
	log_view     *tview.Pages

	AddresViewS AddressViewPrimitive

	udp_soc *udp.Udp_Sock

	elements []tview.Primitive
}

func New() *EmuUI {
	eui := &EmuUI{
		app:          tview.NewApplication(),
		main_frame:   tview.NewGrid(),
		head:         tview.NewGrid(),
		left:         tview.NewPages(),
		address_view: tview.NewGrid(),
		log_view:     tview.NewPages(),

		AddresViewS: AddressViewPrimitive{
			address_IF:      tview.NewInputField(),
			port_IF:         tview.NewInputField(),
			set_addr_button: tview.NewButton("Set"),
		},

		udp_soc: udp.New(),
	}

	u := eui
	u.elements = []tview.Primitive{
		u.AddresViewS.address_IF,
		u.AddresViewS.port_IF,
		u.AddresViewS.set_addr_button,
	}

	u.MakeHeadFrame()
	u.MakeUtilFrame()
	u.MakeLogFrame()
	u.MakeMainFrame()

	return u
}

func (self *EmuUI) SetAddress(addr string, port int) {
	self.udp_soc.SetAddr(addr, port)
}

func (self *EmuUI) Focus(reverse bool) {
	max := len(self.elements)
	for i, el := range self.elements {
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

		self.app.SetFocus(self.elements[i])
		return
	}

	self.app.SetFocus(self.elements[0])
}

func (self *EmuUI) FocusEvent(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyCtrlN:
		self.Focus(false)
	case tcell.KeyCtrlP:
		self.Focus(true)

	}

}

func (self *EmuUI) MakeMainFrame() {
	self.main_frame.SetBorder(true).SetTitle("FinsUDPEmurator").SetTitleAlign(0).SetTitleColor(tcell.ColorYellowGreen)
	self.main_frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			self.app.SetFocus(self.head)

		case tcell.KeyCtrlU:
			self.app.SetFocus(self.left)

		case tcell.KeyCtrlL:
			button := tview.NewButton("button")
			button.SetSelectedFunc(func() {

			})

		}

		return event
	})

	self.main_frame.SetRows(8, 0, 8).SetColumns(45, 0)
	// self.main_frame.SetBackgroundColor(tcell.ColorWhite)
	self.main_frame.AddItem(self.head, 0, 0, 1, 1, 0, 0, true)
	self.main_frame.AddItem(self.left, 1, 0, 3, 1, 0, 0, true)
	self.main_frame.AddItem(self.log_view, 0, 1, 4, 1, 0, 0, true)
}

func (self *EmuUI) MakeHeadFrame() {
	self.AddresViewS.set_addr_button.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		self.FocusEvent(event)

		return event
	})

	self.AddresViewS.set_addr_button.SetSelectedFunc(func() {
		addr := self.AddresViewS.address_IF.GetLabel()
		port := self.AddresViewS.port_IF.GetLabel()
		p, _ := strconv.Atoi(port)
		self.SetAddress(addr, p)

	})

	field_width := 30

	self.AddresViewS.address_IF.
		SetLabel("Address:").
		SetFieldWidth(field_width).
		SetDoneFunc(func(key tcell.Key) {
			self.Focus(false)

		}).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			self.FocusEvent(event)

			return event
		})

	self.AddresViewS.port_IF.
		SetLabel("Port   :").
		SetFieldWidth(field_width).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			self.Focus(false)

		}).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			self.FocusEvent(event)

			return event
		})

	self.address_view.SetBackgroundColor(tcell.ColorBlack)
	self.head.SetBorder(true).SetTitle("<A>ddress & Port")

	self.address_view.SetRows(2, 2).SetColumns(0, 0)
	self.address_view.AddItem(self.AddresViewS.address_IF, 0, 0, 1, 2, 0, 0, true)
	self.address_view.AddItem(self.AddresViewS.port_IF, 1, 0, 1, 2, 0, 0, true)

	self.head.SetRows(4, 0, 1).SetColumns(10, 0, 0, 10)

	self.head.
		AddItem(self.address_view, 0, 0, 1, 4, 0, 0, true).
		AddItem(self.AddresViewS.set_addr_button, 1, 3, 1, 1, 0, 0, true)

}

func (self *EmuUI) MakeUtilMainPage() {

}

func (self *EmuUI) MakeUtilAddPage() {

}

func (self *EmuUI) MakeUtilFrame() {
	self.left.SetBorder(true).SetTitle("<U>til")
	reset_button := tview.NewButton("ResetLog")

	g := tview.NewGrid()
	g.SetRows(0, 0, 0, 0).SetColumns(0, 0, 0)
	g.AddItem(reset_button, 4, 2, 1, 1, 0, 0, true)

	g2 := tview.NewGrid()
	g2.SetRows(0, 0, 0, 0).SetColumns(0, 0, 0)
	g2.AddItem(reset_button, 3, 1, 1, 1, 0, 0, true)
	self.left.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyRune:
			switch event.Rune() {
			case 'A':
				self.left.SwitchToPage("Main")
			case 'C':
				self.left.SwitchToPage("Main2")
			}
		}

		return event
	})

	self.left.
		AddPage("Main", g, true, true).
		AddPage("Main2", g2, true, false)

}

func (self *EmuUI) MakeLogFrame() {
	textView := tview.NewTextView()
	textView.SetBorder(true).SetTitle("Log")
	textView.SetText("test")

	form := tview.NewForm()
	form.SetBorder(true).SetTitle("Add Data")
	form.AddInputField("DM", "", 20, nil, nil)
	form.AddInputField("Data", "", 20, nil, nil)
	form.AddButton("Save", nil)
	form.AddButton("Cancel", func() {
		form.Clear(true)
		self.log_view.SwitchToPage("log")

	})

	self.log_view.
		AddPage("log", textView, true, true).
		AddPage("Add", form, true, false)
}

func (self *EmuUI) RunApp() {
	if err := self.app.SetRoot(self.main_frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
