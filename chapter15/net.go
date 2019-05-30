package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("0.0.0.0", "8888"))
	fmt.Println(net.SplitHostPort("0.0.0.0:8888"))

	fmt.Println(net.LookupAddr("61.135.169.125"))

	fmt.Println(net.LookupHost("www.baidu.com"))

	fmt.Println(net.ParseCIDR("192.168.0.1/24"))

	fmt.Println(net.ParseIP("127.0.0.1"))
	fmt.Println(net.ParseIP("::1"))
	fmt.Println(net.LookupIP("www.baidu.com"))

	ip, ipnet, _ := net.ParseCIDR("192.168.0.1/24")
	fmt.Println(ipnet.Contains(ip))
	fmt.Println(ipnet.Network())

	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		fmt.Println(addr.String(), addr.Network())
	}

	intfs, _ := net.Interfaces()
	for _, intf := range intfs {
		fmt.Println(intf.Index, intf.Name, intf.HardwareAddr, intf.MTU, intf.Flags)
		fmt.Println(intf.Addrs())
		fmt.Println(intf.MulticastAddrs())
	}
}
