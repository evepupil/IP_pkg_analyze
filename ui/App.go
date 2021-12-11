package ui

import (
	"IP_pkg_analyze/ip"
	"fmt"
	fyne "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)
func Run()  {

	a := app.New()
	//a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("简单的IP抓包工具")
	appIconPath,err:=fyne.LoadResourceFromPath("C:\\Users\\Administrator\\Pictures\\QQ截图20190324000505.png")
	if err!=nil{
		panic(err)
	}
	w.SetIcon(appIconPath)
	loadMenus(w)
	//p:=PkgRow{Source: "src",Dest: "dst"}
	b:= ip.Get_if_list()
	PkgList,list:=loadPkgList(w)
	fmt.Println("www",list)
	w.SetContent(container.NewWithoutLayout(PkgList))
	w.Resize(fyne.NewSize(1600, 800))
	go ip.GetPkg(b[0].NPFName,list)
	w.ShowAndRun()
}

