package ui

import (
	"os"

	"github.com/rivo/tview"
)

// https://github.com/rivo/tview/wiki
var (
	curent_dir, _ = os.Getwd()
	json_path     = curent_dir + "/data.json"
)

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
