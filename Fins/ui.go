package fins

import (
	"github.com/rivo/tview"
)

func Test() {
	text := tview.NewTextView()
	text.SetText("test")
	text.SetTitle("Text").SetBorder(true)

	text2 := tview.NewTextView()
	text2.SetText("test2")
	text2.SetTitle("Text2").SetBorder(true)

	text3 := tview.NewTextView()
	text3.SetText("test3")
	text3.SetTitle("Text3").SetBorder(true)

	main_flex := tview.NewFlex().SetDirection(tview.FlexColumn).
		AddItem(text, 0, 1, false).
		AddItem(text2, 0, 1, false).
		AddItem(text3, 0, 1, false)
	main_flex.SetTitle("main_flex").SetBorder(true)

	app := tview.NewApplication()
	if err := app.SetRoot(main_flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

}
