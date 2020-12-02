package screen

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
)

var Screen tcell.Screen

func init() {
	var err error
	Screen, err = tcell.NewScreen()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	if err := Screen.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func Run(input func(), draw func()) {
	for {
		input()
		Screen.Clear()
		draw()
		Screen.Sync()
	}
}
