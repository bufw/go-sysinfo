package sysinfo

import (
	"runtime"

	"gopkg.in/ini.v1"
)

func (o *OSRelease) GetOSInfo() *OSRelease {
	osini, err := ini.Load("/etc/os-release")
	if err != nil {

	}
	o.Name = osini.Section("").Key("NAME").String()
	o.Version = osini.Section("").Key("VERSION_ID").String()
	o.Arch = runtime.GOARCH
	o.Platform = runtime.GOOS
	o.PrettyName = osini.Section("").Key("PRETTY_NAME").String()
	return o
}
