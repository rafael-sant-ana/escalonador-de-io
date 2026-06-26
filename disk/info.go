package disk

type DiskInfo struct {
	maxDiskBytes    int
	initialPosition int
}

func NewDiskInfo(maxDiskBytes int, initialPosition int) DiskInfo {
	return DiskInfo{
		maxDiskBytes,
		initialPosition,
	}
}
