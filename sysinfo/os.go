package sysinfo

type OSRelease struct {
	ID         string
	Name       string
	Version    string
	Platform   string
	Arch       string
	PrettyName string
}

func OSInfo() *OSRelease {
	s := &OSRelease{}
	s.GetOSInfo()
	return s
}
