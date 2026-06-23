package consumers

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

type IoHandler struct {
	diskInfo DiskInfo
	requests []int
}
