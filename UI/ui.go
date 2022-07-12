package ui

import (
	udp "FinsEmu/UDP"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type EmuUI struct {
	app *tview.Application

	main_frame   *tview.Grid
	head         *tview.Grid
	left         *tview.Grid
	address_view *tview.Grid

	address_IF      *tview.InputField
	port_IF         *tview.InputField
	set_addr_button *tview.Button

	udp_soc *udp.Udp_Sock

	elements []tview.Primitive
}

func New() *EmuUI {
	eui := &EmuUI{
		app:          tview.NewApplication(),
		main_frame:   tview.NewGrid(),
		head:         tview.NewGrid(),
		left:         tview.NewGrid(),
		address_view: tview.NewGrid(),
		udp_soc:      udp.New(),

		address_IF:      tview.NewInputField(),
		port_IF:         tview.NewInputField(),
		set_addr_button: tview.NewButton("Set"),
	}

	u := eui
	u.elements = []tview.Primitive{
		u.address_IF,
		u.port_IF,
		u.set_addr_button,
	}

	u.MakeHeadFrame()
	u.MakeLeftFrame()
	u.MakeMainFrame()

	return u
}

func (ui *EmuUI) SetAddress(addr string, port int) {
	ui.udp_soc.SetAddr(addr, port)
}

func (ui *EmuUI) Focus(reverse bool) {
	max := len(ui.elements)
	for i, el := range ui.elements {
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

		ui.app.SetFocus(ui.elements[i])
		return
	}

	ui.app.SetFocus(ui.elements[0])
}

func (ui *EmuUI) FocusEvent(event *tcell.EventKey) {
	switch event.Key() {
	case tcell.KeyCtrlN:
		ui.Focus(false)
	case tcell.KeyCtrlP:
		ui.Focus(true)

	}

}

func (ui *EmuUI) MakeMainFrame() {
	ui.main_frame.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			ui.app.SetFocus(ui.head)

		}

		return event
	})

	ui.main_frame.SetRows(0, 0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0, 0, 0, 0).SetBorders(true)
	// ui.main_frame.SetBackgroundColor(tcell.ColorWhite)
	ui.main_frame.AddItem(ui.head, 0, 0, 3, 2, 0, 0, true)
	ui.main_frame.AddItem(ui.left, 3, 0, 2, 1, 0, 0, true)
}

func (ui *EmuUI) MakeHeadFrame() {
	ui.set_addr_button.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		ui.FocusEvent(event)

		return event
	})

	ui.set_addr_button.SetSelectedFunc(func() {
		addr := ui.address_IF.GetLabel()
		port := ui.port_IF.GetLabel()
		p, _ := strconv.Atoi(port)
		ui.SetAddress(addr, p)

	})

	field_width := 30

	ui.address_IF.
		SetLabel("Address:").
		SetFieldWidth(field_width).
		SetDoneFunc(func(key tcell.Key) {
			ui.Focus(false)

		}).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			ui.FocusEvent(event)

			return event
		})

	ui.port_IF.
		SetLabel("Port   :").
		SetFieldWidth(field_width).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			ui.Focus(false)

		}).
		SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
			ui.FocusEvent(event)

			return event
		})

	ui.address_view.SetBackgroundColor(tcell.ColorBlack)
	ui.head.SetBorder(true).SetTitle("<A>")

	ui.address_view.SetRows(0, 0).SetColumns(0, 0)
	ui.address_view.AddItem(ui.address_IF, 0, 0, 1, 2, 0, 0, true)
	ui.address_view.AddItem(ui.port_IF, 1, 0, 1, 2, 0, 0, true)

	ui.head.SetRows(0, 0, 0, 0).SetColumns(0, 0)

	ui.head.
		AddItem(ui.address_view, 0, 0, 3, 2, 0, 0, true).
		AddItem(ui.set_addr_button, 3, 0, 1, 1, 0, 0, true)

}

func (ui *EmuUI) MakeLeftFrame() {
	change_addr_button := tview.NewButton("")
	change_addr_button.SetBorder(true).SetBackgroundColor(tcell.ColorBlue)
	change_addr_button.SetSelectedFunc(func() { /*to change view*/ })

	// ui.left.
	// 	AddItem(ui.address_view, 0, 0, 3, 1, 10, 10, false).
	// 	AddItem(change_addr_button, 1, 0, 3, 1, 10, 10, false).SetBorder(true).SetTitle("left")

}

func (ui *EmuUI) RunApp() {
	if err := ui.app.SetRoot(ui.main_frame, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
