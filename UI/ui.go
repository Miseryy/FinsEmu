package ui

import "github.com/rivo/tview"

// https://github.com/rivo/tview/wiki

func RunApp() {
	app := tview.NewApplication()
	frames := Frames{
		App: app,
	}
	main_f := NewMainFrame(&frames)

	if err := app.SetRoot(main_f.MakeFrame(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
