package sysinfo

import (
	"fmt"
)

type SysInfo struct {
	Load1     float64
	Load5     float64
	Load15    float64
	UpTime    int64
	Totalram  uint64
	Freeram   uint64
	Sharedram uint64
	Bufferram uint64
	Totalswap uint64
	Freeswap  uint64
	Procs     uint16
	Platform  string
	Arch      string
}

func GetSysInfo() *SysInfo {
	s := &SysInfo{}
	s.Update()
	return s
}

func (s *SysInfo) LoadAvgString() string {
	return fmt.Sprintf("load average: %2.2f, %2.2f, %2.2f", s.Load1, s.Load5, s.Load15)
}
