package main

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

var c chan struct{}

func main() {
	c = make(chan struct{})
	go onChh()
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Pretty awesome超级棒")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.ClickedCh = c
	// Sets the icon of a menu item. Only available on Mac and Windows.
	mQuit.SetIcon(icon.Data)
}

func onExit() {
	// clean up here
}

func onChh() {
	for {
		select {
		case <-c:
			fmt.Println("ok")
			os.Exit(0)
		}
	}
}
