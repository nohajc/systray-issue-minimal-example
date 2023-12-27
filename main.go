package main

import (
	_ "embed"

	"fyne.io/systray"
)

//go:embed st-icon.png
var appIcon []byte

func main() {
	systray.Run(onReady, func() {})
}

type ItemGroup struct {
	checkedIdx int
	items      []*systray.MenuItem
}

func (g *ItemGroup) AddItems(items ...*systray.MenuItem) {
	if len(items) == 0 {
		return
	}
	if len(g.items) == 0 {
		items[0].Check()
	}
	g.items = append(g.items, items...)
}

func (g *ItemGroup) Check(idx int) {
	if idx < 0 || idx >= len(g.items) {
		return
	}
	g.items[g.checkedIdx].Uncheck()
	g.items[idx].Check()
	g.checkedIdx = idx
}

func onReady() {
	systray.SetTemplateIcon(appIcon, appIcon)
	item1 := systray.AddMenuItemCheckbox("item 1", "", true)
	item2 := systray.AddMenuItemCheckbox("item 2", "", false)
	item3 := systray.AddMenuItemCheckbox("item 3", "", false)
	group1 := &ItemGroup{}
	group1.AddItems(item1, item2, item3)

	submenu := systray.AddMenuItem("submenu", "")
	subitem1 := submenu.AddSubMenuItemCheckbox("subitem 1", "", true)
	subitem2 := submenu.AddSubMenuItemCheckbox("subitem 2", "", false)
	subitem3 := submenu.AddSubMenuItemCheckbox("subitem 3", "", false)
	group2 := &ItemGroup{}
	group2.AddItems(subitem1, subitem2, subitem3)

	for {
		select {
		case <-item1.ClickedCh:
			group1.Check(0)
		case <-item2.ClickedCh:
			group1.Check(1)
		case <-item3.ClickedCh:
			group1.Check(2)
		case <-subitem1.ClickedCh:
			group2.Check(0)
		case <-subitem2.ClickedCh:
			group2.Check(1)
		case <-subitem3.ClickedCh:
			group2.Check(2)
		}
	}
}
