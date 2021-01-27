package main

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	gtk.Init(nil)

	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window")
	}
	win.SetTitle("Pen trainer")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	ebox, err := gtk.EventBoxNew()
	if err != nil {
		log.Fatal("Unable to create fixed")
	}
	fixed, err := gtk.FixedNew()
	if err != nil {
		log.Fatal("Unable to create fixed")
	}

	score := 0

	l, err := gtk.LabelNew(strconv.Itoa(score))
	if err != nil {
		log.Fatal("Unable to create btt")
	}

	x := 10
	y := 10

	btt, err := gtk.ButtonNewWithLabel("X")
	if err != nil {
		log.Fatal("Unable to create btt")
	}
	btt.Connect("pressed", func() {
		w, h := win.GetSize()
		x = int(rand.Int31n(int32(w - 50)))
		y = int(rand.Int31n(int32(h - 50)))
		fixed.Move(btt, x, y)
		score++
		l.SetText(strconv.Itoa(score))
	})
	btt.SetSizeRequest(50, 50)

	ebox.Connect("button-press-event", func() {
		score = 0
		l.SetText("0")
	})

	w, h := win.GetSize()
	ebox.SetSizeRequest(w, h)
	fixed.Put(ebox, 0, 0)
	fixed.Put(btt, x, y)
	fixed.Put(l, w/2-10, h-30)

	//Used to be on win but that made the move function stop working on the btt connect
	fixed.Connect("size_allocate", func() {
		w, h := win.GetSize()
		fixed.Move(btt, x, y)
		fixed.Move(l, w/2-10, h-30)
		ebox.SetSizeRequest(w, h)
	})

	win.Add(fixed)

	win.SetDefaultSize(1920, 1080)
	win.ShowAll()
	gtk.Main()
}
