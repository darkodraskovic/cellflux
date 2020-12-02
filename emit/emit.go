package emit

import (
	"github.com/gdamore/tcell/v2"
)

var screen tcell.Screen

func Init(s tcell.Screen) {
	screen = s
}

func Char(x, y int, r rune) {
	screen.SetContent(x, y, r, nil, style)
}

func String(x, y int, str string) {
	w, _ := screen.Size()
	for _, c := range str {
		screen.SetContent(x%w, y+(x/w), c, nil, style)
		x++
	}
}

func HLine(x1, x2, y int) {
	for i := x1; i < x2; i++ {
		screen.SetContent(i, y, tcell.RuneHLine, nil, style)
	}

}

func VLine(x, y1, y2 int) {
	for i := y1; i < y2; i++ {
		screen.SetContent(x, i, tcell.RuneVLine, nil, style)
	}
}

func Rect(x, y, w, h int) {
	r := x + w - 1
	l := y + h - 1
	screen.SetContent(x, y, tcell.RuneULCorner, nil, style)
	screen.SetContent(r, y, tcell.RuneURCorner, nil, style)
	screen.SetContent(r, l, tcell.RuneLRCorner, nil, style)
	screen.SetContent(x, l, tcell.RuneLLCorner, nil, style)
	HLine(x+1, r, y)
	HLine(x+1, r, l)
	VLine(x, y+1, l)
	VLine(r, y+1, l)
}

func HWall(x1, x2, y int) {
	for i := x1; i < x2; i++ {
		screen.SetContent(i, y, tcell.RuneBlock, nil, style)
	}

}

func VWall(x, y1, y2 int) {
	for i := y1; i < y2; i++ {
		screen.SetContent(x, i, tcell.RuneBlock, nil, style)
	}
}

func Room(x, y, w, h int) {
	right := x + w
	low := y + h
	HWall(x, right, y)
	HWall(x, right, low-1)
	VWall(x, y, low)
	VWall(right-1, y, low)
}

func Log(str string) {
	w, h := screen.Size()
	offset := h - len(str)/w - 2
	HLine(0, w, offset)
	String(0, offset, "LOG ")
	String(0, offset+1, str)
}
