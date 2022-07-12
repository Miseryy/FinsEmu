package ui

import (
	udp "FinsEmu/UDP"
	"bytes"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type EmuUI struct {
	app *tview.Application

	main_frame   *tview.Grid
	head         *tview.Grid
	left         *tview.Grid
	address_view *tview.Grid

	address_IF *tview.InputField
	port_IF    *tview.InputField

	udp_soc *udp.Udp_Sock
}

func New() *EmuUI {
	eui := &EmuUI{
		app:          tview.NewApplication(),
		main_frame:   tview.NewGrid(),
		head:         tview.NewGrid(),
		left:         tview.NewGrid(),
		address_view: tview.NewGrid(),
		udp_soc:      udp.New(),

		address_IF: tview.NewInputField(),
		port_IF:    tview.NewInputField(),
	}

	u := eui
	u.MakeHeadFrame()
	u.MakeLeftFrame()
	u.MakeMainFrame()

	return u
}

func (ui *EmuUI) SetAddress(addr string, port int) {
	ui.udp_soc.SetAddr(addr, port)
}

func (ui *EmuUI) MakeMainFrame() {
	ui.main_frame.SetRows(0, 0, 0, 0, 0, 0).SetColumns(0, 0, 0).SetBorders(true)
	ui.main_frame.SetBackgroundColor(tcell.ColorWhite)
	ui.main_frame.AddItem(ui.head, 0, 0, 1, 1, 0, 0, true)
	ui.main_frame.AddItem(ui.left, 1, 0, 2, 1, 0, 0, true)

}

func (ui *EmuUI) MakeHeadFrame() {
	change_addr_button := tview.NewButton("ChangeAddress")
	change_addr_button.SetBorder(true).SetBackgroundColor(tcell.ColorBlue)
	change_addr_button.SetSelectedFunc(func() { /*to change view*/ })

	ui.address_IF.
		SetLabel("Address:").
		SetFieldWidth(10).
		SetDoneFunc(func(key tcell.Key) {
			// fmt.Println(key)
			ui.app.SetFocus(ui.port_IF)

		})

	ui.port_IF.
		SetLabel("Port   :").
		SetFieldWidth(10).
		SetAcceptanceFunc(tview.InputFieldInteger).
		SetDoneFunc(func(key tcell.Key) {
			// fmt.Println(key)
			ui.app.SetFocus(ui.left)

		})

	ui.address_view.SetBackgroundColor(tcell.ColorBlack)

	ui.address_view.SetRows(0, 0).SetColumns(0, 0)
	ui.address_view.AddItem(ui.address_IF, 0, 0, 1, 2, 0, 0, true)
	ui.address_view.AddItem(ui.port_IF, 1, 0, 1, 2, 0, 0, true)

	ui.head.SetRows(0, 0).SetColumns(0)

	ui.head.
		AddItem(ui.address_view, 0, 0, 2, 2, 0, 0, true)

}

func (ui *EmuUI) MakeLeftFrame() {
	change_addr_button := tview.NewButton("ChangeAddress")
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

func Test() {
	text := tview.NewTextView()
	text.SetText("test")
	text.SetTitle("Text").SetBorder(true)

	text2 := tview.NewTextView()
	text2.SetText(string(bytes.Repeat([]byte("sfasf"), 300)))
	text2.SetTitle("Text2").SetBorder(true)
	text2.SetFocusFunc(func() {

	})

	text3 := tview.NewTextView()
	text3.SetText("test3")
	text3.SetTitle("Text3").SetBorder(true)

	main_flex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(text, 0, 1, false).
		AddItem(text2, 0, 1, true).
		AddItem(text3, 0, 1, false)
	main_flex.SetTitle("main_flex").SetBorder(true)

	app := tview.NewApplication()
	if err := app.SetRoot(main_flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
