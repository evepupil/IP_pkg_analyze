package ui

import (
	"bytes"
	"fmt"
	fyne2 "fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/google/gopacket"
	"strconv"
	//"time"
)

type LayersInfo struct {
	PkgInfos                                                                  []string
	LinkLayersInfos, NetWorkLayersInfos, TransportLayersInfos, AppLayersInfos []string
}

var (
	LayersWidget fyne2.Widget
	LayersData   map[string][]string
	LayersAll    []string
)

func loadLayers() *fyne2.Container {
	LayersData = map[string][]string{}
	LayersWidget = widget.NewTreeWithStrings(LayersData)
	scrollC := container.NewScroll(LayersWidget)
	speparator := widget.NewSeparator()
	PkgInfoContainer := container.NewBorder(nil, speparator, nil, nil, scrollC)
	PkgInfoContainer.Resize(fyne2.NewSize(1600, 240))
	return PkgInfoContainer
}
func NewLayersData(FrameNo int, packet gopacket.Packet) map[string][]string {
	PkgInfoBbranch, L0Infos := getPkgInfoData(FrameNo, packet)
	LayersAll = append(LayersAll, PkgInfoBbranch)
	LayersData[PkgInfoBbranch] = L0Infos
	if packet.LinkLayer() != nil {
		LinkLayerBranch, L1Infos := getLinkLayerData(packet)
		LayersData[LinkLayerBranch] = L1Infos
		LayersAll = append(LayersAll, LinkLayerBranch)
	}
	if packet.NetworkLayer() != nil {
		NetWorkLayerBranch, L2Infos := getNetWorkLayerData(packet)
		LayersData[NetWorkLayerBranch] = L2Infos
		LayersAll = append(LayersAll, NetWorkLayerBranch)
	}
	if packet.TransportLayer() != nil {
		TransportLayerBranch, L3Infos := getTransportLayerData(packet)
		LayersData[TransportLayerBranch] = L3Infos
		LayersAll = append(LayersAll, TransportLayerBranch)
	}
	if packet.ApplicationLayer() != nil {
		AppLayerBranch, L4Infos := getAppLayerData(packet)
		LayersData[AppLayerBranch] = L4Infos
		LayersAll = append(LayersAll, AppLayerBranch)
	}
	//fmt.Print(PkgInfoBuffer.String())
	return nil
}
func getPkgInfoData(FrameNo int, packet gopacket.Packet) (branch string, nodes []string) {
	var PkgInfoBuffer, InterfaceBuffer bytes.Buffer
	metadata := packet.Metadata()
	fmt.Fprintf(&PkgInfoBuffer, "Frame %d: %d bytes on wire (%d bits),%d bytes captured(%d bits) "+
		"on interface %s , id:%d", FrameNo, metadata.Length, metadata.Length*8, metadata.CaptureLength, metadata.CaptureLength*8,
		DeviceName, metadata.InterfaceIndex)
	branch = PkgInfoBuffer.String()
	fmt.Fprintf(&InterfaceBuffer, "Interface id: %d (%s)",
		metadata.InterfaceIndex, DeviceName) //设备信息
	time := metadata.Timestamp.Format("2006-01-02T15:04:05") //时间
	No := "Frame Number: " + strconv.Itoa(FrameNo)           //No
	FrameLength := "Frame Length: " + strconv.Itoa(metadata.Length) + "(" + strconv.Itoa(metadata.Length*8) + "bits)"
	CaptureLength := "Capture Length" + strconv.Itoa(metadata.CaptureLength) + "(" + strconv.Itoa(metadata.CaptureLength*8) + "bits)"
	nodes = append(nodes, InterfaceBuffer.String(), "Arrival Time : "+time, No, FrameLength, CaptureLength)
	return
}
func getLinkLayerData(packet gopacket.Packet) (branch string, nodes []string) {
	linkLayerMetaData := packet.LinkLayer()
	var linkLayerInfoBuffer bytes.Buffer
	fmt.Fprintf(&linkLayerInfoBuffer, "%s , Src: %s , Dst: %s", linkLayerMetaData.LayerType().String(),
		linkLayerMetaData.LinkFlow().Src().String(), linkLayerMetaData.LinkFlow().Dst().String())
	branch = linkLayerInfoBuffer.String()
	Dst := "Destination: " + linkLayerMetaData.LinkFlow().Dst().String()
	Src := "Source: " + linkLayerMetaData.LinkFlow().Src().String()
	Type := "Type: IPV6(0x86dd)" //IPV6
	if linkLayerMetaData.LayerContents()[12] == 8 {
		Type = "Type: IPV4(0x0800)" //IPV4
	}
	nodes = append(nodes, Dst, Src, Type)
	return
}
func getNetWorkLayerData(packet gopacket.Packet) (branch string, nodes []string) {
	newtworkLayerMetaData := packet.NetworkLayer()
	var networkLayerInfoBuffer bytes.Buffer
	fmt.Fprintf(&networkLayerInfoBuffer, "Internet Protocol Version %d, Src: %s, Dst: %s",
		newtworkLayerMetaData.LayerContents()[0]/16, newtworkLayerMetaData.NetworkFlow().Src(), newtworkLayerMetaData.NetworkFlow().Dst())

	return
}
func getTransportLayerData(packet gopacket.Packet) (branch string, nodes []string) {
	return
}
func getAppLayerData(packet gopacket.Packet) (branch string, nodes []string) {
	return
}
