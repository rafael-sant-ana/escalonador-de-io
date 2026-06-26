package disk

type DiskInfo struct {
	maxDiskBytes    int
	initialPosition int
}

func (d DiskInfo) GetMaxDiskBytes() int {
	return d.maxDiskBytes
}

func (d DiskInfo) GetInitialPosition() int {
	return d.initialPosition
}

func NewDiskInfo(maxDiskBytes int, initialPosition int) DiskInfo {
	return DiskInfo{
		maxDiskBytes,
		initialPosition,
	}
}
