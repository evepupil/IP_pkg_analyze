package ui

import (
	"IP_pkg_analyze/ip"
	fyne2 "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var DeviceName string

func Run() {
	a := app.New()
	//a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("简单的IP抓包工具")
	appIconPath, err := fyne2.LoadResourceFromPath("C:\\Users\\Administrator\\Pictures\\QQ截图20190324000505.png")
	if err != nil {
		fyne2.LogError("icon加载失败", err)
	}
	w.SetIcon(appIconPath)
	loadMenus(w)
	//p:=PkgRow{Source: "src",Dest: "dst"}
	b := ip.Get_if_list()
	Layers := loadLayers()
	PkgInfo := loadPkgInfo()
	PkgList, list := loadPkgList()
	PkgListContainer := container.NewWithoutLayout(PkgList)
	//PkgListContainer.Resize(fyne2.NewSize(1600,280))
	//PkgInfoContainer:=container.NewBorder(PkgListContainer,nil,nil,nil,PkgInfo)
	//PkgInfoContainer.Resize(fyne2.NewSize(1600,280))
	Layers.Move(fyne2.NewPos(0, 280))
	PkgInfo.Move(fyne2.NewPos(0, 520))
	AllContainer := container.NewWithoutLayout(PkgListContainer, Layers, PkgInfo)
	w.SetContent(AllContainer)
	w.Resize(fyne2.NewSize(1600, 800))
	DeviceName = b[0].NPFName
	go ip.GetPkg(b[0].NPFName, list)
	w.ShowAndRun()
}
