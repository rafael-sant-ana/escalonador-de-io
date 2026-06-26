package consumers

import (
	"io-scheduling/disk"
)

type IoHandler struct {
	diskInfo        disk.DiskInfo
	requests        []int
	currentPosition int
}
