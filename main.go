package main

import(
	"fmt"
	"log"
	_"database/sql"
	"strings"
	_"math/rand"
	_"time"
	//"goui_adv/protagonist"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jroimartin/gocui"
)

var intro int = 0

func storyColor(s string, c string) string {
	switch c {
	case "him":
		return "\033[35;1m"+s+"\033[0m"
	case "hiw":
		return "\033[37;1m"+s+"\033[0m"
	case "tst1":
		return "\033[31;1m"+s+"\033[0m"
	default:
		return "\033[37;1m"+s+"\033[0m"
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	if v, err := g.SetView("v1", 0, 0, maxX-30, maxY-10); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Story"
		v.Wrap = true
		
		intro := storyColor("Full Moon Madness\n\n", "him")
		fmt.Fprintf(v, intro)

		hiMagenta := storyColor("Press Enter to begin.", "him")
		fmt.Fprintf(v, hiMagenta)
	}

	if v, err := g.SetView("v2", maxX-30, 0, maxX-1, maxY/2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Inventory"
	}

	if v, err := g.SetView("v3", maxX-30, maxY/2, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Stats"
	}

	if v, err := g.SetView("v4", 0, maxY-10, maxX-50, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Actions"
		v.Highlight = true
		//v.SelFgColor = gocui.ColorGreen
		v.Editable = true

		if _, err := g.SetCurrentView("v4"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("v5", maxX-50, maxY-10, maxX-30, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Location"
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func createPlayer(g *gocui.Gui, v *gocui.View) error {

	vbuf := v.ViewBuffer()

	v, err := g.View("v1")
	if err != nil {
		return err
	}

	if vbuf == "" {
		v.Clear()
			
		enterName := storyColor("Please enter your characters first and last name separated by a space and press enter.", "hiw")
		fmt.Fprintf(v, enterName)
	} else {
		name := strings.TrimRight(vbuf, " ")
	
		fullName := strings.Split(name, " ")
	
		fname := fullName[0]
		lname := fullName[1]
	
		v.Clear()
			
		helloName := storyColor("Hello "+fname+" "+lname+", let us develop your character. What is your occupation?", "hiw")
		fmt.Fprintf(v, helloName)

	}



	return nil

}

func action(g *gocui.Gui, v *gocui.View) error {
	vbuf := v.ViewBuffer()
	word := strings.TrimSpace(vbuf)

	if word == "next" {

		v.Clear()
		v.SetCursor(0, 0)

		v, err := g.View("v1")
		if err != nil {
			return err
		}

		v.Clear()
		
		hiMagenta := storyColor("And now we have moved beyond", "hiw")
		fmt.Fprintf(v, hiMagenta)

	} else {

		v.Clear()
		v.SetCursor(0, 0)

		v, err := g.View("v1")
		if err != nil {
			return err
		}

		v.Clear()

		hiMagenta := storyColor("Oops, I messed up", "tst1")
		fmt.Fprintf(v, hiMagenta)
		fmt.Fprintf(v, "\n"+vbuf)
	}

	return nil
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}

	defer g.Close()

	g.Highlight = true
	g.Cursor = true

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if intro == 0 {
		if err := g.SetKeybinding("v4", gocui.KeyEnter, gocui.ModNone, createPlayer); err != nil {
		log.Panicln(err)
		}
	} else {
		if err := g.SetKeybinding("v4", gocui.KeyEnter, gocui.ModNone, action); err != nil {
			log.Panicln(err)
		}

	}


	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}