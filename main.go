package main

import (
	"fmt"
	"log"
	"os"

	"github.com/darkodraskovic/cellflux/emit"
	"github.com/darkodraskovic/cellflux/mainflux"
	"github.com/darkodraskovic/cellflux/screen"
	"github.com/gdamore/tcell/v2"
)

var (
	title  string = "CELLFLUX"
	code   string
	user   mainflux.User = mainflux.User{Email: "john.doe@email.net", Password: "12345678"}
	token  mainflux.Token
	things mainflux.Things
)

func draw() {
	w, h := screen.Screen.Size()
	_ = h
	var col int
	emit.String(w/2-len(title)/2, col, title)
	col++
	emit.String(w/2-len(user.Email)/2, col, user.Email)
	col++
	emit.HLine(0, w, col)
	emit.String(0, col, "THINGS ")
	col++
	emit.PushAttributes(tcell.AttrBold)
	emit.String(0, col, fmt.Sprintf("Total: %d\n", things.Total))
	col++
	emit.PushAttributes(tcell.AttrItalic)
	emit.String(0, col, fmt.Sprintf("Offset: %d\n", things.Offset))
	col++
	emit.String(0, col, fmt.Sprintf("Limit: %d\n", things.Limit))
	col++
	emit.PopAttributes()
	for _, value := range things.Things {
		emit.String(0, col, fmt.Sprintf("%+v", value))
		col++
	}
	emit.PopAttributes()
	emit.Log(fmt.Sprintf("%+v", things))

	// emit.HWall(40, 50, 30)
	// emit.Rect(20, 20, 10, 10)
	// emit.Room(20, 20, 10, 10)
	// emit.Rect(0, 0, w, h)

}

func input() {
	scr := screen.Screen
	switch ev := scr.PollEvent().(type) {
	case *tcell.EventResize:
		scr.Sync()
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			scr.Fini()
			os.Exit(0)
		}
	}
}

func init() {

}

func main() {
	emit.Init(screen.Screen)
	emit.Background(tcell.ColorBlack)
	emit.Foreground(tcell.ColorGreen)

	var err error
	token, err = mainflux.GetToken(user.Email, user.Password)
	if err != nil {
		log.Fatalln(err)
	}
	// code, err = mainflux.CreateUser("john.doe1@email.net", "12345678")
	if err != nil {
		log.Fatalln(err)
	}

	things, err = mainflux.GetThings(token)
	if err != nil {

	}

	screen.Run(input, draw)
}
