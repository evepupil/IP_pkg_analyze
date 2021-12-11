package ui

import (
	"fmt"
	fyne2 "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	//"time"
)
func loadPkgList(w fyne2.Window) (*fyne2.Container,binding.StringList) {
	pkgText:=[]string{
		"No.     ","Time                          ","Source              ","Dst                      ",
		"Protocol   ","Length     ","Info                                 ",
	}
	pkgTextContainer:=container.NewHBox()
	for _,v:=range pkgText{
		pkgTextContainer.Add(widget.NewButton(v, func() {

		}))
	}
	PkgStringList:=binding.NewStringList()

	PkgList:=widget.NewListWithData(PkgStringList,
		func() fyne2.CanvasObject {
			return widget.NewLabel("")
		},
		func(item binding.DataItem, object fyne2.CanvasObject) {
			i:=item.(binding.String)
			l:=object.(*widget.Label)
			s,_:=i.Get()
			l.SetText(s)
		})
	PkgList.OnSelected= func(id widget.ListItemID) {
		fmt.Println(id)

	}
	s:=widget.NewSeparator()
	PkgInfoContainer:=container.NewBorder(pkgTextContainer,nil,nil,s,PkgList)
	PkgInfoContainer.Resize(fyne2.NewSize(1600,280))
	return PkgInfoContainer,PkgStringList

}
