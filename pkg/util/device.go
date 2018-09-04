package util

import (
	"fmt"
	"path/filepath"

	"github.com/shirou/gopsutil/disk"
)

func IsDeviceUnderRoot(p string) error {
	const root = "/"
	var (
		rootVolName string
		curVolName  string
	)
	ps, e := disk.Partitions(true)
	if e != nil {
		return e
	}
	for i := range ps {
		if ps[i].Mountpoint == root {
			rootVolName = ps[i].Device
		}
		if filepath.HasPrefix(p, ps[i].Mountpoint) && (len(curVolName) == 0 || ps[i].Mountpoint != root) {
			curVolName = ps[i].Device
		}
	}
	if curVolName == rootVolName {
		return fmt.Errorf("path is under root device, \"%s\"==\"%s\"", curVolName, rootVolName)
	}
	return nil
}
