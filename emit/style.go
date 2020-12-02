package emit

import "github.com/gdamore/tcell/v2"

var (
	style tcell.Style
	masks []tcell.AttrMask
)

func Decompose() (fg tcell.Color, bg tcell.Color, attr tcell.AttrMask) {
	return style.Decompose()
}

func Foreground(c tcell.Color) {
	style = style.Foreground(c)
}

func Background(c tcell.Color) {
	style = style.Background(c)
}

func Attributes(attrs tcell.AttrMask) {
	masks = []tcell.AttrMask{attrs}
	style = style.Attributes(attrs)
}

func PushAttributes(mask tcell.AttrMask) {
	masks = append(masks, mask)
	_, _, attrs := style.Decompose()
	style = style.Attributes(attrs | mask)
}

func PopAttributes() {
	if len(masks) < 1 {
		return
	}
	masks = masks[:len(masks)-1]
	var m tcell.AttrMask
	for _, v := range masks {
		m |= v
	}
	style = style.Attributes(m)
}
