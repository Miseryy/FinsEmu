package main

import ui "FinsEmu/UI"

func main() {
	// ui.Test()
	u := ui.New()
	u.SetAddress("127.0.0.1", 60000)
	u.RunApp()

}
