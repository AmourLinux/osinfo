package goInfo

import (
	"fmt"
)

type OSInfo struct {
	GoOS     string
	GoOSBig  string
	Kernel   string
	Platform string
	OS       string
	Hostname string
	CPUs     int
}

func (osi *OSInfo) PrintFields() {
	fmt.Println("GoOS:", osi.GoOS)
	fmt.Println("GoOSBig:", osi.GoOSBig)
	fmt.Println("Kernel:", osi.Kernel)
	fmt.Println("Platform:", osi.Platform)
	fmt.Println("OS:", osi.OS)
	fmt.Println("Hostname:", osi.Hostname)
	fmt.Println("CPUs:", osi.CPUs)
}

func (osi *OSInfo) String() string {
	return fmt.Sprintf("GoOS:%v,GoOSBig:%v,Kernel:%v,Platform:%v,OS:%v,Hostname:%v,CPUs:%v", osi.GoOS, osi.GoOSBig, osi.Kernel, osi.Platform, osi.OS, osi.Hostname, osi.CPUs)
}
