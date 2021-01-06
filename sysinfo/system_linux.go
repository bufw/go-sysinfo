package sysinfo

import (
	"runtime"
	"strings"
	"syscall"
)

func (s *SysInfo) Update() {
	info := syscall.Sysinfo_t{}
	syscall.Sysinfo(&info)
	s.Load1 = float64(info.Loads[0]) / 65536.0
	s.Load5 = float64(info.Loads[1]) / 65536.0
	s.Load15 = float64(info.Loads[2]) / 65536.0
	s.UpTime = info.Uptime //按秒计算
	s.Totalram = info.Totalram
	s.Totalswap = info.Totalswap
	s.Bufferram = info.Bufferram
	s.Freeram = info.Freeram
	s.Freeswap = info.Freeswap
	s.Sharedram = info.Sharedram
	s.Platform = strings.Title(runtime.GOOS)
	s.Arch = runtime.GOARCH
}
