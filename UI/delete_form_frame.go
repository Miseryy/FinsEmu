package ui

import (
	jsonutil "FinsEmu/JsonUtil"
	"fmt"
	"sort"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type DeleteFormFrame struct {
	js                   *jsonutil.MyJson
	change2LogFrame_call func()
}

func NewDeleteFormFrame(j *jsonutil.MyJson) *DeleteFormFrame {
	return &DeleteFormFrame{
		js: j,
	}
}

func (self *DeleteFormFrame) MakeFrame() tview.Primitive {
	table := tview.NewTable().SetBorders(true)
	table.SetBorder(true).SetTitle("Delete DM Data")
	table.SetBorderPadding(0, 0, 10, 10)

	self.makeCells(table)

	return table
}

func (self *DeleteFormFrame) makeCells(table *tview.Table) {
	self.js.LoadJson()
	items := self.js.GetMap()
	var keys []int
	for key := range items {
		i_key, _ := strconv.Atoi(key)
		keys = append(keys, i_key)
	}

	table.SetSelectable(true, false)

	table.SetCell(0, 0, tview.NewTableCell("DM").SetTextColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorRed).SetAlign(tview.AlignCenter).SetExpansion(1).SetMaxWidth(10))
	table.SetCell(0, 1, tview.NewTableCell("Data(Hex)").SetTextColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorRed).SetAlign(tview.AlignCenter).SetExpansion(1).SetMaxWidth(6))

	sort.Ints(keys)

	for i, k := range keys {
		kk := strconv.Itoa(k)
		val := items[kk]
		var value int
		switch v := val.(type) {
		case float64:
			value = int(v)
		}

		table.SetCell(i+1, 0, tview.NewTableCell(kk).SetAlign(tview.AlignCenter))
		hex_fmt := fmt.Sprintf("0x%04X", value)
		table.SetCell(i+1, 1, tview.NewTableCell(hex_fmt).SetAlign(tview.AlignCenter))
	}

	table.Select(1, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			self.change2LogFrame_call()
		}
	}).SetSelectedFunc(func(row int, column int) {
		table.GetCell(row, column).SetTextColor(tcell.ColorRed)
		key := table.GetCell(row, 0).Text
		self.js.DeleteItem(key).WriteJson()
		table.Clear()
		self.makeCells(table)
	}).SetSelectionChangedFunc(func(row, column int) {
		if row == 0 {
			table.Select(1, 0)
			return
		}
	})
}
