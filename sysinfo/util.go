package sysinfo

func Int8ToString(a []int8) string {
	buf := make([]byte, 0, len(a))
	for _, val := range a {
		if val == 0x00 {
			break
		}
		buf = append(buf, byte(val))
	}
	return string(buf)
}
