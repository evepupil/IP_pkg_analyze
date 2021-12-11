package ui

import (
	fyne "fyne.io/fyne/v2"
)

func loadMenus(w fyne.Window)  {
	tools_key := []string{
		"文件(F)", "编辑(E)", "视图(V)", "跳转(G)",
		"捕获(C)", "分析(A)", "统计(S)", "电话(Y)",
		"无线(W)", "工具(T)", "帮助(H)",
	}
	var tools map[string][]string
	Menus :=[]*fyne.Menu{}
	for _,key:=range tools_key{
		tMenu:=[]*fyne.MenuItem{}
		for _,item:=range tools[key]{
			t:=fyne.NewMenuItem(item, func() {

			})
			tMenu=append(tMenu,t)
		}
		Menus=append(Menus,fyne.NewMenu(key,tMenu...))
	}
	w.SetMainMenu(fyne.NewMainMenu(Menus...))

}
