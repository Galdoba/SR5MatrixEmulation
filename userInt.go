package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
	"time"
	"strings"
)

var (
	viewArr = []string{"v1", "v2", "v3", "v4", "v5"}
	active  = 0
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	out, err := g.View("v2")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 4 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 0, 0, int(0.2*float32(maxX)), maxY-7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Status"
		v.Editable = true
		v.Wrap = true

		if _, err = setCurrentViewOnTop(g, "v1"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("v2", int(0.6*float32(maxX)), 0, maxX-1, maxY-7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Event log"
		v.Wrap = true
		v.Autoscroll = true
	}
	if v, err := g.SetView("v3", 0, maxY-7, int(0.6*float32(maxX)), maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Grid Info"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "Press TAB to change current view")
	}
	if v, err := g.SetView("v4", int(0.6*float32(maxX)), maxY-7, maxX-1, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Console"
		v.Editable = true
	}
    if v, err := g.SetView("v5", int(0.2*float32(maxX)), 0, int(0.6*float32(maxX)), maxY-7); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Enviroment"
		v.Wrap = true
		v.Autoscroll = true
		fmt.Fprint(v, "Press TAB to change current view")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func station() {
	
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

    

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}
    if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, v5Output); err != nil {
		log.Panicln(err)
	}
    if err := g.SetKeybinding("", gocui.KeyHome, gocui.ModNone, start); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
        
		log.Panicln(err)
	}
   // createRooster()
}

func logOutput(g *gocui.Gui, v *gocui.View) error {
    out, err := g.View("v5")
	if err != nil {
		return err
	}
    fmt.Fprintln(out, "Enter was pressed...")
    out, err = g.View("v2")
    if err != nil {
		return err
	}
    out.Clear()
    
	fmt.Fprintln(out, "Current Time:", time.Now().Format(time.RFC1123))
    return nil
}

func v5Output (g *gocui.Gui, v *gocui.View) error {
    eventLog, err := g.View("v2")
	if err != nil {
		return err
	}
   
    console, err := g.View("v4")
    if err != nil {
		return err
	}
    text := console.Buffer()
    text = strings.TrimRight(text, "\n")
    text = "\033[32m" + text + "\033[0m"
    console.SetCursor(0,0)
    console.Clear()
    
	
    fmt.Fprintln(eventLog, text)
    return nil


    return nil
}

func start (g *gocui.Gui, v *gocui.View) error {
         //createRooster()
    return nil
}
