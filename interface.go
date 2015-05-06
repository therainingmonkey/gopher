package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

// TODO Colours, frames, formatting etc.

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// Draw the footer
	if f, err := g.SetView("footer", 0, maxY-3, maxX-1, maxY-1); err != nil {
		// Only happens the first time the view is drawn
		if err != gocui.ErrorUnkView {
			return err
		}
		fmt.Fprintln(f, "(G)o somewhere; go (B)ackwards")
	}

	// Draw the main section
	if m, err := g.SetView("main", 0, 0, maxX-1, maxY-4); err != nil {
		// Only happens the first time the view is drawn
		if err != gocui.ErrorUnkView {
			return err
		}
		parsedLines, err := loadHomepage()
		if err != nil {
			return err
		}
		for _, line := range parsedLines {
			fmt.Fprintln(m, line[0])
		}
		if err := g.SetCurrentView("main"); err != nil {
			return err
		}
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.Quit
}

func dialogURL(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if u, err := g.SetView("dialogURL", maxX/5, maxY/2-1, maxX/5*4, maxY/2+1); err != nil {
		if err != gocui.ErrorUnkView {
			return err
		}
		// Update the footer
		f, err := g.View("footer")
		if err != nil {
			return err
		}
		f.Clear()
		fmt.Fprintln(f, "^Q: Cancel; Enter: Go!")
		u.Editable = true
		g.SetCurrentView("dialogURL")
		g.ShowCursor = true
	}
	return nil
}

func exitDialog(g *gocui.Gui, v *gocui.View) error {
	view, err := g.View(v.Name())
	if err != nil {
		return err
	}
	if err = g.DeleteView(view.Name()); err != nil {
		log.Panicln(err)
	}
	g.ShowCursor = false
	f, err := g.View("footer")
	if err != nil {
		return err
	}
	f.Clear()
	fmt.Fprintln(f, "(G)o somewhere; go (B)ackwards")
	g.SetCurrentView("main")
	return nil
}

// TODO error handling, restructuring?
func visitURL(g *gocui.Gui, url string) error {
	bodyLines, err := loadPage(url)
	if err != nil {
		return err
	}
	parsedLines := parseBody(bodyLines)
	m, err := g.View("main")
	if err != nil {
		return err
	}
	m.Clear()
	for _, line := range parsedLines {
		fmt.Fprintln(m, line[0])
	}
	dialogURL, _ := g.View("dialogURL")
	exitDialog(g, dialogURL)
	return nil
}

func main() {
	var err error
	gui := gocui.NewGui()
	if err = gui.Init(); err != nil {
		log.Panicln(err)
	}
	defer gui.Close()
	gui.SetLayout(layout)
	if err = setKeybindings(gui); err != nil {
		log.Panicln(err)
	}
	if err = gui.MainLoop(); err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}
