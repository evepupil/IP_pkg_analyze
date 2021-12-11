package main

import "IP_pkg_analyze/ui"

func main() {
	// Find all devices
	b:= get_if_list()
	ui.InitFont()
	//getPkg(b[0].NPFName)
	go GetPkg(b[0].NPFName)
	ui.Run()  //会阻塞


}
