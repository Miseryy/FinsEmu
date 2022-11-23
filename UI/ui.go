package ui

import (
	"fmt"
	"os"

	"github.com/rivo/tview"
)

// https://github.com/rivo/tview/wiki
var (
	curent_dir, _     = os.Getwd()
	data_json_path    = curent_dir + "/data.json"
	setting_json_path = curent_dir + "/setting.json"
	log_path          = curent_dir + "/data"
)

func RunApp() {
	if _, err := os.Stat(data_json_path); err != nil {
		f, err := os.Create(data_json_path)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		f.Write([]byte("{}"))
		f.Close()
	}

	app := tview.NewApplication()
	frames := Frames{
		App: app,
	}
	main_f := NewMainFrame(&frames)

	if err := app.SetRoot(main_f.MakeFrame(), true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
