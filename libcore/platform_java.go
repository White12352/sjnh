package libcore

var (
	intfBox BoxPlatformInterface
	intfGUI GUIInterface
)

var useProcfs bool

type GUIInterface interface {
	UseOfficialAssets() bool
	SelectorCallback(tag string)
}

// iif
type BoxPlatformInterface interface {
	AutoDetectInterfaceControl(fd int32) error
	OpenTun(singTunOptionsJson, tunPlatformOptionsJson string) (int, error)
	UseProcFS() bool
	FindConnectionOwner(ipProtocol int32, sourceAddress string, sourcePort int32, destinationAddress string, destinationPort int32) (int32, error)
	PackageNameByUid(uid int32) (string, error)
	UIDByPackageName(packageName string) (int32, error)
	ReadWIFIState() *WIFIState

	StartDefaultInterfaceMonitor(listener InterfaceUpdateListener) error
	CloseDefaultInterfaceMonitor(listener InterfaceUpdateListener) error
	GetInterfaces() (NetworkInterfaceIterator, error)
	UsePlatformInterfaceGetter() bool
}

type InterfaceUpdateListener interface {
	UpdateDefaultInterface(interfaceName string, interfaceIndex int32)
}

type NetworkInterface struct {
	Index     int32
	MTU       int32
	Name      string
	Addresses StringIterator
}

type NetworkInterfaceIterator interface {
	Next() *NetworkInterface
	HasNext() bool
}
