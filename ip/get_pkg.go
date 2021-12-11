package ip

import (
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
	"strconv"
	"time"
)

type PkgRow struct {
	No       int
	Time     time.Time
	Source   string
	Dest     string
	Protocol string
	Length   int
	Info     string
}

//设备名：pcap.FindAllDevs()返回的设备的Name
//snaplen：捕获一个数据包的多少个字节，一般来说对任何情况65535是一个好的实践，如果不关注全部内容，只关注数据包头，可以设置成1024
//promisc：设置网卡是否工作在混杂模式，即是否接收目的地址不为本机的包
//timeout：设置抓到包返回的超时。如果设置成30s，那么每30s才会刷新一次数据包；设置成负数，会立刻刷新数据包，即不做等待
//要记得释放掉handle

var (
	device       string = "eth0"
	snapshot_len int32  = 1024
)

func GetPkg(device string, list binding.StringList) {
	No := 1
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	fmt.Println("gogogo")
	// Use the handle as a packet source to process all packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		//packet:=<-packetSource.Packets()
		if packet != nil {
			p := anlysePacket(packet)
			p.No = No
			No++
			fmt.Println(gopacket.LayerString(packet.TransportLayer()))

			list.Append(p.formatePkgListInfo())
		}
	}
}

func setFlliter(device string, p string, port int) {
	// Open device
	handle, err = pcap.OpenLive(device, snapshot_len, promiscuous, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	// Set filter
	var filter string = p + " and port " + strconv.Itoa(port)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Only capturing " + p + " port " + strconv.Itoa(port) + " packets.")

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Do something with a packet here.
		fmt.Println(packet)
	}
}
func anlysePacket(p gopacket.Packet) PkgRow {
	var nilEndpoint gopacket.Endpoint = gopacket.Endpoint{}
	var nilFlow gopacket.Flow = gopacket.Flow{}
	pkgrow := PkgRow{Time: p.Metadata().Timestamp,
		Length: p.Metadata().Length,
	}
	if p.NetworkLayer()!=nil{  //网际层
		if p.NetworkLayer().NetworkFlow().Src() != nilEndpoint {
			pkgrow.Source = p.NetworkLayer().NetworkFlow().Src().String()
		}
		if p.NetworkLayer().NetworkFlow().Dst()!=nilEndpoint {
			pkgrow.Dest = p.NetworkLayer().NetworkFlow().Dst().String()
		}
	}
	if p.TransportLayer() != nil {   //传输层
		if p.TransportLayer().TransportFlow() != nilFlow {
			pkgrow.Protocol = p.TransportLayer().TransportFlow().EndpointType().String()
			pkgrow.Info = p.TransportLayer().TransportFlow().String()
		}
	}
	return pkgrow

}
func (p PkgRow) formatePkgListInfo() string {
	res := ""
	res += strconv.Itoa(p.No)
	t := time.Now().Format("\"2006-01-02T15:04:05\"")
	res += blankAdd(15-len(res)) + t[1:len(t)-1]
	res += blankAdd(45-len(res)) + p.Source
	res += blankAdd(70-len(res)) + p.Dest
	res += blankAdd(98-len(res)) + p.Protocol
	res += blankAdd(115-len(res)) + strconv.Itoa(p.Length)
	res += blankAdd(129-len(res)) + p.Info
	return res
}
func blankAdd(n int) string {
	res := ""
	for n > 0 {
		n--
		res += " "
	}
	return res

}
