package sysinfo

type OSRelease struct {
	ID            string
	Name          string
	Version       string
	Platform      string
	Arch          string
	PrettyName    string
	KernelVersion string
	// Hostname      string
	Sysname      string
	Nodename     string
	UnameVersion string
	Machine      string
	Domainname   string
}

func OSInfo() *OSRelease {
	s := &OSRelease{}
	s.GetOSInfo()
	return s
}

func (o *OSRelease) isSystemVersion() {

}
