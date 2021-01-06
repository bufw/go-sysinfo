package sysinfo

import (
	"runtime"
	"syscall"

	"gopkg.in/ini.v1"
)

func (o *OSRelease) GetOSInfo() *OSRelease {
	var uname syscall.Utsname
	syscall.Uname(&uname)
	o.Sysname = Int8ToString(uname.Sysname[:])
	o.Nodename = Int8ToString(uname.Nodename[:])
	o.KernelVersion = Int8ToString(uname.Release[:])
	o.UnameVersion = Int8ToString(uname.Version[:])
	o.Machine = Int8ToString(uname.Machine[:])
	o.Domainname = Int8ToString(uname.Domainname[:])

	// osini, err := ini.Load("/etc/os-release")
	var osini *ini.File
	var err error
	linuxRelease := []string{"/etc/os-release", "/etc/lsb-release"}
	for _, v := range linuxRelease {
		osini, err = ini.Load(v)
		if err != nil {
			continue
		}
		break
	}
	if err != nil {
		o.Default()
		return o
	}
	o.ID = osini.Section("").Key("ID").String()
	o.Name = osini.Section("").Key("NAME").String()
	o.Version = osini.Section("").Key("VERSION_ID").String()
	o.Arch = runtime.GOARCH
	o.Platform = runtime.GOOS
	o.PrettyName = osini.Section("").Key("PRETTY_NAME").String()
	return o
}

func (o *OSRelease) Default() {
	o.Name = "Unknown"
	o.Version = ""
	o.Arch = runtime.GOARCH
	o.Platform = runtime.GOOS
	o.PrettyName = "Generic Linux"
}
