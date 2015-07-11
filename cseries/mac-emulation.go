// reimplementation of macintosh_utilities.c in cseries
package cseries

import (
	"syscall"
)

func uptime() (int64, error) {
	sysinfo := syscall.Sysinfo_t{}

	if err := syscall.Sysinfo(&sysinfo); err != nil {
		return 0.0, err
	} else {
		return sysinfo.Uptime, nil
	}
}

func MachineTickCount() (int64, error) {
	return uptime()
}
