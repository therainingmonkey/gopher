package main

import (
	"github.com/jroimartin/gocui"
)

func setKeybindings(g *gocui.Gui) (err error) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	if err := g.SetKeybinding("main", 'g', gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return dialogURL(g)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("dialogURL", gocui.KeyCtrlQ, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return exitDialog(g, v)
		}); err != nil {
		return err
	}
	if err := g.SetKeybinding("dialogURL", gocui.KeyEnter, gocui.ModNone,
		func(g *gocui.Gui, v *gocui.View) error {
			return visitURL(g, v.Buffer())
		}); err != nil {
		return err
	}
	return nil
}
