package ui

import (
	fyne2 "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	//"github.com/google/gopacket"
)

func loadPkgInfo() (*fyne2.Container, binding.String) {
	bs := binding.NewString()
	s := " "
	bs.Set(s)
	l := widget.NewLabelWithData(bs)
	scrollC := container.NewVScroll(l)
	speparator := widget.NewSeparator()
	PkgInfoContainer := container.NewBorder(nil, speparator, nil, nil, scrollC)
	PkgInfoContainer.Resize(fyne2.NewSize(1600, 240))
	return PkgInfoContainer, bs
}
func PkgBytes2String(PkgBytes []byte) string {
	res := ""
	for i, v := range PkgBytes {
		if i != 0 && i%8 == 0 {
			res += "   "
		}
		if i != 0 && i%16 == 0 {
			res += "\n"
		}
		res += byte2Hex(v) + " "
	}
	return res
}
func byte2Hex(b byte) string {
	care := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	bb := int(b)
	res := ""
	if b == 0 {
		return "00"
	}
	if b < 16 {
		return "0" + care[bb%16]
	}
	for bb > 0 {
		res = care[bb%16] + res
		bb /= 16
	}
	return res

}
